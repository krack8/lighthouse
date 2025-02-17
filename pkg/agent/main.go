package main

import (
	"context"
	"crypto/x509"
	"encoding/json"
	agentClient "github.com/krack8/lighthouse/pkg/agent/client"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/config"
	_log "github.com/krack8/lighthouse/pkg/log"
	"github.com/krack8/lighthouse/pkg/tasks"
	"log"
	"sync"
	"time"
)

var taskMutex sync.Mutex

func main() {
	_log.InitializeLogger()
	config.InitEnvironmentVariables()
	config.InitiateKubeClientSet()
	// For demonstration, we'll just run a single worker that belongs to "GroupA".
	groupName := config.WorkerGroup
	controllerURL := config.ServerUrl
	secretName := config.AgentSecretName
	resourceNamespace := config.ResourceNamespace
	//authToken := "my-secret"
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
		defer cancel() // Cancel the context when the program exits

		conn, stream, err := agentClient.ConnectAndIdentifyWorker(ctx, controllerURL, secretName, resourceNamespace, groupName, caCertPool)
		if err != nil {
			_log.Logger.Fatalw("Failed to connect and identify worker", "error", err)
			continue
		}
		streamErrorChan := make(chan error)
		// handle incoming messages in a separate goroutine.
		go func() {
			for {
				in, err := stream.Recv()
				if err != nil {
					_log.Logger.Infow("Stream Recv error (worker)", "err", err)
					streamErrorChan <- err
					return
				}

				switch payload := in.Payload.(type) {

				case *pb.TaskStreamResponse_NewTask:
					task := payload.NewTask
					_log.Logger.Infow("Worker received a new task: ID=%s, payload=%s",
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

				case *pb.TaskStreamResponse_Ack:
					_log.Logger.Infow("Worker received an ACK from server: "+payload.Ack.Message, "info", "ACK")

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

		// Keep the worker alive.
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
