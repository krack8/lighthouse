package main

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	agentClient "github.com/krack8/lighthouse/pkg/agent/client"
	"github.com/krack8/lighthouse/pkg/agent/tasks"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	_log "github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/controller/auth/utils"
	"google.golang.org/grpc"
	"io"
	corev1 "k8s.io/api/core/v1"
	"log"
	"strings"
	"sync"
	"time"
)

var taskMutex sync.Mutex

// var logsMutex sync.Mutex
var logTaskMap = make(map[string]*logTask)
var logTaskMapMutex = sync.Mutex{}

const (
	PodLogBoolTrue     = "y"
	PodLogBoolFalse    = "n"
	TailLinesThreshold = int64(2500)
)

type logTask struct {
	cancel    context.CancelFunc
	heartbeat *time.Timer
}

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
					if podLogsTask.Payload == "cancel" {
						cancelTask(podLogsTask.Id, stream)
					} else if podLogsTask.Payload == "heartbeat" {
						keepAliveLogs(podLogsTask.Id, stream)
					} else {
						logsCtx, logsCancel := context.WithCancel(stream.Context())
						newLogTask := logTask{logsCancel, nil}
						logTaskMapMutex.Lock()
						logTaskMap[podLogsTask.Id] = &newLogTask
						logTaskMapMutex.Unlock()
						keepAliveLogs(podLogsTask.Id, stream)
						go func(taskID, taskPayload string, task *pb.PodLogsStream, ctx context.Context) {
							var input k8s.GetPodLogsInputParams
							err := json.Unmarshal([]byte(task.Input), &input)
							if err != nil {
								_log.Logger.Errorw("Failed to unmarshal pod logs input", "err", err)
							}
							logResult := &pb.LogsResult{TaskId: taskID, Cancel: false}
							podLogOptions := corev1.PodLogOptions{Follow: true}
							if input.Container != "" {
								podLogOptions.Container = input.Container
							}
							zero := int64(0)
							if input.TailLines != nil && *input.TailLines != zero {
								if *input.TailLines > TailLinesThreshold {
									threshold := TailLinesThreshold
									podLogOptions.TailLines = &threshold
								} else {
									podLogOptions.TailLines = input.TailLines
								}
							}
							if input.SinceSeconds != nil && *input.SinceSeconds > zero {
								podLogOptions.SinceSeconds = input.SinceSeconds
							}
							// checking timestamps
							switch input.Timestamps {
							case PodLogBoolTrue:
								podLogOptions.Timestamps = true
							case PodLogBoolFalse:
								podLogOptions.Timestamps = false
							}

							req := k8s.GetKubeClientSet().CoreV1().Pods(input.NamespaceName).GetLogs(input.Pod, &podLogOptions)
							podLogs, err := req.Stream(ctx)
							if err != nil {
								_log.Logger.Errorw("failed to get pod logs of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
								cancelTask(taskID, stream)
							}
							defer func() {
								if podLogs != nil {
									_ = podLogs.Close()
								}
							}()

							buf := make([]byte, 2048)
							for {
								select {
								case <-ctx.Done():
									_log.Logger.Infow("log streaming task cancelled", "logs-stream", "cancelled")
									if podLogs != nil {
										_ = podLogs.Close()
									}
									return // Exit the goroutine
								default:
									numBytes, err := podLogs.Read(buf)
									if err != nil {
										logTaskMapMutex.Lock()
										_, ok := logTaskMap[taskID]
										logTaskMapMutex.Unlock()
										if !ok {
											return
										}
										if err == io.EOF {
											_log.Logger.Infow("End of pod logs stream", "taskID", taskID)
										} else {
											_log.Logger.Errorw("failed to read logs of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
										}
										cancelTask(taskID, stream)
										return // Exit the goroutine
									}

									logResult.Output = buf[:numBytes]
									if err := stream.Send(&pb.TaskStreamRequest{
										Payload: &pb.TaskStreamRequest_LogsResult{
											LogsResult: logResult,
										},
									}); err != nil {
										_log.Logger.Errorw("failed to send log message of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
										cancelTask(taskID, stream)
									}
									time.Sleep(1 * time.Second)
								}
							}

						}(podLogsTask.Id, podLogsTask.Payload, podLogsTask, logsCtx)
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
				case *pb.TaskStreamResponse_ExecReq:
					task := payload.ExecReq
					//_log.Logger.Infof("Agent received a new pod exec task: ID=%s, Payload=%s",
					//	task.TaskId, task.Payload)
					err := tasks.PodExecTask(task.TaskId, task.Payload, task.Input, task.Command, stream)
					if err != nil {
						_log.Logger.Errorw("Failed to send pod exec result", "err", err)
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

func keepAliveLogs(taskID string, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) {
	log.Printf("keep alive func task with ID: %s", taskID)
	logsTimeout := 10 * time.Second

	logTaskMapMutex.Lock()
	defer logTaskMapMutex.Unlock()

	task, ok := logTaskMap[taskID]
	if !ok {
		log.Printf("No active task found with ID: %s", taskID)
		return
	}

	resetHeartbeat(task, taskID, logsTimeout, stream)
}

func resetHeartbeat(task *logTask, taskID string, logsTimeout time.Duration, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) {
	if task.heartbeat != nil {
		task.heartbeat.Stop()
	}

	task.heartbeat = time.AfterFunc(logsTimeout, func() {
		_log.Logger.Infow(fmt.Sprintf("logs timeout for task: %s", taskID), "logs-timeout", taskID)
		cancelTask(taskID, stream)
		//log.Printf("logs timeout for task: %s", taskID)
		//task.cancel()
		//logTaskMapMutex.Lock()
		//delete(logTaskMap, taskID)
		//logTaskMapMutex.Unlock()
		//log.Printf("Cancelled task with ID: %s", taskID)
	})
}

func cancelTask(taskID string, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) {
	_log.Logger.Infow(fmt.Sprintf("Cancelling func task with ID: %s", taskID), "log-cancel", taskID)
	logTaskMapMutex.Lock()
	defer logTaskMapMutex.Unlock()
	if task, ok := logTaskMap[taskID]; ok {
		task.cancel()
		if task.heartbeat != nil {
			task.heartbeat.Stop()
		}
		delete(logTaskMap, taskID)
		_log.Logger.Infow("Cancelled task", "log-cancel", taskID)
		if err := stream.Send(&pb.TaskStreamRequest{
			Payload: &pb.TaskStreamRequest_LogsResult{
				LogsResult: &pb.LogsResult{TaskId: taskID, Output: nil, Cancel: true},
			},
		}); err != nil {
			_log.Logger.Errorw("failed to send log closed message", "log-cancel", taskID)
		}
	} else {
		_log.Logger.Warnw("No active task found", "log-cancel", taskID)
		if err := stream.Send(&pb.TaskStreamRequest{
			Payload: &pb.TaskStreamRequest_LogsResult{
				LogsResult: &pb.LogsResult{TaskId: taskID, Output: nil, Cancel: true},
			},
		}); err != nil {
			_log.Logger.Errorw("failed to send log closed message [No active task found]", "log-cancel", taskID)
		}
	}
}
