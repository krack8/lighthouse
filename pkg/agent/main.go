package main

import (
	"context"
	"crypto/x509"
	"encoding/json"
	agentClient "github.com/krack8/lighthouse/pkg/agent/client"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	_log "github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/controller/auth/utils"
	"io"
	corev1 "k8s.io/api/core/v1"
	"log"
	"strings"
	"sync"
	"time"
)

var taskMutex sync.Mutex
var logsMutex sync.Mutex

func main() {
	_log.InitializeLogger()
	config.InitEnvironmentVariables()
	k8s.InitiateKubeClientSet()

	// For demonstration, we'll just run a single worker that belongs to "GroupA".
	groupName := config.AgentGroup
	controllerGrpcServerHost := config.ControllerGrpcServerHost
	secretName := config.AgentSecretName
	resourceNamespace := config.ResourceNamespace

	groupName, err := utils.GetAgentGroup(secretName, resourceNamespace)
	if err != nil {
		log.Fatalf("Failed to get "+secretName+" secret: %v", err)
	}
	// Validate and provide a default if needed
	if groupName == "" {
		_log.Logger.Errorw("Missing env variable AGENT_GROUP", "err", "AGENT_GROUP env variable is not found in kubernetes secret")
	}

	_log.Logger.Infow("Starting agent", "groupName", groupName)

	tasks.InitTaskRegistry()
	var caCertPool *x509.CertPool
	if config.TlsServerCustomCa != "" {
		caCert := []byte(config.TlsServerCustomCa)

		// Create a new certificate pool
		caCertPool = x509.NewCertPool()

		// Append the certificate to the pool
		if !caCertPool.AppendCertsFromPEM(caCert) {
			log.Fatalf("Failed to append CA certificate to pool")
		}
	}
	streamRecoveryMaxAttempt := 10
	streamRecoveryInterval := 5 * time.Second
	for streamRecoveryAttempt := 0; streamRecoveryAttempt < streamRecoveryMaxAttempt; streamRecoveryAttempt++ {
		ctx, cancel := context.WithCancel(context.Background())
		//defer cancel() // Cancel the context when the program exits

		conn, stream, err := agentClient.ConnectAndIdentifyWorker(ctx, controllerGrpcServerHost, secretName, resourceNamespace, groupName, caCertPool)
		if err != nil {
			_log.Logger.Fatalw("Failed to connect and identify agent", "error", err)
			continue
		}
		streamErrorChan := make(chan error)
		// handle incoming messages in a separate goroutine.
		go func() {
			for {
				in, err := stream.Recv()
				if err != nil {
					_log.Logger.Infow("Stream Receive error (agent)", "err", err)
					streamErrorChan <- err
					return
				}

				switch payload := in.Payload.(type) {

				case *pb.TaskStreamResponse_NewTask:
					task := payload.NewTask
					_log.Logger.Infow("Agent received a new task: ID=%s, payload=%s",
						task.Id, task.Payload)
					go func(taskID, taskPayload string, task *pb.Task) {
						taskMutex.Lock()
						defer taskMutex.Unlock()
						TaskResult := &pb.TaskResult{}
						res, err := tasks.TaskSelector(task)
						if err != nil {
							TaskResult.Success = false
							TaskResult.Output = err.Error()
						} else {
							output, err := json.Marshal(res)
							if err != nil {
								TaskResult.Success = false
								TaskResult.Output = err.Error()
							}
							TaskResult.Success = true
							TaskResult.Output = string(output)
						}
						TaskResult.TaskId = taskID
						resultMsg := &pb.TaskStreamRequest{
							Payload: &pb.TaskStreamRequest_TaskResult{
								TaskResult: TaskResult,
							},
						}

						// Send the result back to the controller.
						if err = stream.Send(resultMsg); err != nil {
							_log.Logger.Errorw("Failed to send task result", "err", err)
						}
					}(task.Id, task.Payload, task)

				case *pb.TaskStreamResponse_NewPodLogsStream:
					podLogsTask := payload.NewPodLogsStream
					_log.Logger.Infow("Agent received a new pod logs task: ID=%s, payload=%s",
						podLogsTask.Id, podLogsTask.Payload)
					go func(taskID, taskPayload string, task *pb.PodLogsStream) {
						logsMutex.Lock()
						defer logsMutex.Unlock()
						var input k8s.GetPodLogsInputParams
						err = json.Unmarshal([]byte(podLogsTask.Input), &input)
						if err != nil {
							_log.Logger.Errorw("Failed to unmarshal pod logs input", "err", err)
						}
						logResult := &pb.LogsResult{}
						podLogOptions := corev1.PodLogOptions{Follow: true}
						if input.Container != "" {
							podLogOptions.Container = input.Container
						}
						req := k8s.GetKubeClientSet().CoreV1().Pods(input.NamespaceName).GetLogs(input.Pod, &podLogOptions)
						podLogs, err := req.Stream(context.Background())
						if err != nil {
							_log.Logger.Errorw("failed to get pod logs of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
						}
						defer podLogs.Close()
						logResult.TaskId = taskID

						buf := make([]byte, 2048)
						for {
							numBytes, err := podLogs.Read(buf)
							_log.Logger.Infow("Pod logs read", "numBytes", string(buf[:numBytes]))
							if err == io.EOF {
								_log.Logger.Errorw("error reading pod logs stream", "err", err)
								break
							}
							if err != nil {
								_log.Logger.Errorw("failed to read logs of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
								break
							}

							//// Create and send LogMessage through gRPC stream
							//logMessage := &pb.LogsResult{
							//	//... populate log message fields
							//
							//	Output: string(buf[:numBytes]),
							//}
							logResult.Output = string(buf[:numBytes])
							if err := stream.Send(&pb.TaskStreamRequest{
								Payload: &pb.TaskStreamRequest_LogsResult{
									LogsResult: logResult,
								},
							}); err != nil {
								_log.Logger.Errorw("failed to send log message of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
								break
							}
						}
						logResult.Output = ""
						resultMsg := &pb.TaskStreamRequest{
							Payload: &pb.TaskStreamRequest_LogsResult{
								LogsResult: logResult,
							},
						}
						if err = stream.Send(resultMsg); err != nil {
							_log.Logger.Errorw("Failed to send task result", "err", err)
						}
					}(podLogsTask.Id, podLogsTask.Payload, podLogsTask)

				case *pb.TaskStreamResponse_Ack:
					_log.Logger.Infow("Agent received an ACK from server: "+payload.Ack.Message, "info", "ACK")

					if strings.ReplaceAll(payload.Ack.Message, " ", "") == "InvalidAgentToken" {
						_log.Logger.Errorw("Unauthorized: token is invalid", "err", "unauthorized")
						// Close the gRPC connection
						if err = conn.Close(); err != nil {
							_log.Logger.Warnw("Failed to close gRPC connection", "error", err)
						}

						// Cancel the context to stop any ongoing operations
						cancel()

						// Close the error channel to signal disconnect
						close(streamErrorChan)
						log.Fatalf("Unauthorized: token is invalid")
						return
					}

					if payload.Ack.Message == "group_name_required" {
						_log.Logger.Errorw("Connection rejected: group name is required", "err", "rejected")
						log.Fatalf("Connection rejected: group name is required")
					}

					if payload.Ack.Message == "disconnect_requested" {
						_log.Logger.Infow("Disconnect requested, starting shutdown")
						// Close the gRPC connection
						if err := conn.Close(); err != nil {
							_log.Logger.Warnw("Failed to close gRPC connection", "error", err)
						}

						// Cancel the context to stop any ongoing operations
						cancel()

						// Close the error channel to signal disconnect
						close(streamErrorChan)
						log.Fatalf("Detached from server. Disconnected !!")
						return
					}

				default:
					_log.Logger.Infow("Unknown payload from server.", "payload", "default")
				}
			}
		}()

		// Keep the agent alive.
		select {
		case <-ctx.Done(): // Context cancelled (e.g., shutdown)
			_log.Logger.Infow("Context cancelled")
			break // Exit outer loop
		case err := <-streamErrorChan: // Stream error received
			_log.Logger.Warnw("Stream error detected", "error", err)
			if err := conn.Close(); err != nil { // Close connection immediately
				_log.Logger.Warnw("Failed to close controller connection", "error", err)
			}
			time.Sleep(streamRecoveryInterval)
			continue // Retry connecting
		}
	}
}
