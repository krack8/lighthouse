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
	"log"
	"strings"
	"sync"
	"time"
)

var taskMutex sync.Mutex

func main() {
	_log.InitializeLogger()
	config.InitEnvironmentVariables()
	k8s.InitiateKubeClientSet()

	// For demonstration, we'll just run a single worker that belongs to "GroupA".
	groupName := config.AgentGroup
	controllerGrpcServerHost := config.ControllerGrpcServerHost
	secretName := config.AgentSecretName
	resourceNamespace := config.ResourceNamespace

	if config.IsAuth() {
		agentGroupName, err := utils.GetAgentGroup(secretName, resourceNamespace)
		if err != nil {
			log.Fatalf("Failed to get "+secretName+" secret: %v", err)
		}
		groupName = agentGroupName
		// Group name is required
		if groupName == "" {
			_log.Logger.Errorw("Missing env variable AGENT_GROUP", "err", "AGENT_GROUP env variable is not found in kubernetes secret")
		}
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
					_ = tasks.LogStreamTask(podLogsTask, stream)
				case *pb.TaskStreamResponse_ExecReq:
					task := payload.ExecReq
					err := tasks.PodExecTask(task.TaskId, task.Payload, task.Input, task.Command, stream)
					if err != nil {
						_log.Logger.Errorw("Failed to send pod exec result", "err", err)
						return
					}
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
