package server

import (
	"context"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/config"
	_log "github.com/krack8/lighthouse/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func ConnectAndIdentifyWorker(ctx context.Context, controllerURL, secretName, resourceNamespace, groupName string) (*grpc.ClientConn, grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse], error) {
	maxAttempts := 30 // Maximum retry attempts
	retryInterval := 2 * time.Second
	var conn *grpc.ClientConn
	var err error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		conn, err = grpc.NewClient(controllerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			_log.Logger.Warnw(fmt.Sprintf("Failed to dial controller. Retrying %d", attempt+1), "error", err)
			time.Sleep(retryInterval)
			continue
		}
	}

	if err != nil {
		log.Fatalf("Failed to dial controller after max attempts %d", maxAttempts)
	}

	for attempt := 0; attempt < maxAttempts; attempt++ {
		stream, err := pb.NewControllerClient(conn).TaskStream(ctx)
		if err != nil {
			_log.Logger.Warnw(fmt.Sprintf("Failed to create TaskStream. Retrying %d", attempt+1), "err", err)
			time.Sleep(retryInterval)
			continue
		}

		secretToken := ""
		if config.IsAuth() {
			secretToken, err = utils.GetSecret(secretName, resourceNamespace)
			if err != nil {
				_log.Logger.Warnw(fmt.Sprintf("Failed to get secret. Retrying %d", attempt+1), "err", err)
				time.Sleep(5 * time.Second)
				continue
			}
		} else {
			_log.Logger.Warnw("Auth disabled. Skipping get secret...", "warn", "auth-disabled")
		}

		err = sendWorkerIdentification(ctx, stream, groupName, secretToken)
		if err != nil {
			_log.Logger.Warnw(fmt.Sprintf("Failed to send worker info, reconnecting to stream. Retrying %d", attempt), "err", err)
			_ = stream.CloseSend() // Close the stream before reconnecting
			continue               // Reconnect to the stream
		}

		return conn, stream, nil // Success! Return the stream
	}
	return nil, nil, fmt.Errorf("failed to connect to server after max attempts %d", maxAttempts)
}

func sendWorkerIdentification(ctx context.Context, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse], groupName, secretToken string) error {
	for {
		select {
		case <-ctx.Done(): // Check for context cancellation
			return ctx.Err()
		default:
			err := stream.Send(&pb.TaskStreamRequest{
				Payload: &pb.TaskStreamRequest_WorkerInfo{
					WorkerInfo: &pb.WorkerIdentification{
						GroupName: groupName,
						AuthToken: secretToken,
					},
				},
			})
			if err != nil {
				_log.Logger.Warnw("Failed to send worker info", "err", err)
				time.Sleep(2 * time.Second)
				continue // Retry sending
			}
			return nil // Successfully sent
		}
	}
}

//
//func main() {
//	_log.InitializeLogger()
//	config.InitEnvironmentVariables()
//	config.InitiateKubeClientSet()
//
//	groupName := "GroupA"
//	tasks.InitTaskRegistry()
//
//	controllerURL := os.Getenv("CONTROLLER_URL")
//	secretName := os.Getenv("AGENT_SECRET_NAME")
//	resourceNamespace := os.Getenv("RESOURCE_NAMESPACE")
//
//	// Retry loop for stream recovery
//	for {
//		ctx, cancel := context.WithCancel(context.Background()) // Create a new context for each attempt
//		conn, stream, err := server.ConnectAndIdentifyWorker(ctx, controllerURL, secretName, resourceNamespace, groupName)
//		if err != nil {
//			_log.Logger.Fatalw("Failed to connect and identify worker", "error", err)
//			time.Sleep(5 * time.Second) // Wait before retrying connection
//			continue                    // Retry connecting
//		}
//		defer func() { // Close connection and cancel context when loop exits or error occurs
//			if err := conn.Close(); err != nil {
//				_log.Logger.Warnw("Failed to close controller connection", "error", err)
//			}
//			cancel()
//		}()
//
//		// Handle incoming messages in a separate goroutine.
//		go func() {
//			for {
//				in, err := stream.Recv()
//				if err != nil {
//					_log.Logger.Infow("Stream Recv error (worker)", "err", err)
//					// Stream error occurred. Break out of the loop to trigger reconnect.
//					break
//				}
//				// Process the received message 'in'
//				_log.Logger.Infow("Received message", "message", in) // Example: log the message
//				// ... your message processing logic ...
//			}
//		}()
//
//		// Keep the main goroutine alive.  You might have other tasks here.
//		// The loop will exit if the stream.Recv() goroutine exits.
//		select {} // Block indefinitely
//	}
//}
//
//func main() {
//	_log.InitializeLogger()
//	config.InitEnvironmentVariables()
//	config.InitiateKubeClientSet()
//
//	groupName := "GroupA"
//	controllerURL := os.Getenv("CONTROLLER_URL")
//	secretName := os.Getenv("AGENT_SECRET_NAME")
//	resourceNamespace := os.Getenv("RESOURCE_NAMESPACE")
//
//	tasks.InitTaskRegistry()
//	streamRecoveryMaxAttempt := 10
//	streamRecoveryInterval := 5 * time.Second
//
//	for streamRecoveryAttempt := 0; streamRecoveryAttempt < streamRecoveryMaxAttempt; streamRecoveryAttempt++ {
//		ctx, cancel := context.WithCancel(context.Background()) // Create context WITHIN the loop
//		defer cancel()                                          // Cancel the context when the loop iteration finishes
//
//		conn, stream, err := ConnectAndIdentifyWorker(ctx, controllerURL, secretName, resourceNamespace, groupName)
//		if err != nil {
//			_log.Logger.Errorw("Failed to connect and identify worker", "error", err)
//			time.Sleep(streamRecoveryInterval)
//			continue // Retry connecting
//		}
//
//		// Channel to signal stream errors
//		streamErrorChan := make(chan error)
//
//		go func() {
//			for {
//				in, err := stream.Recv()
//				if err != nil {
//					_log.Logger.Infow("Stream Recv error (worker)", "err", err)
//					streamErrorChan <- err // Signal the error
//					return                 // Goroutine exits
//				}
//				// Process the received message 'in'
//				_log.Logger.Infow("Received message", "message", in)
//				// ... your message processing logic ...
//			}
//		}()
//
//		select {
//		case <-ctx.Done(): // Context cancelled (e.g., shutdown)
//			_log.Logger.Infow("Context cancelled")
//			break // Exit outer loop
//		case err := <-streamErrorChan: // Stream error received
//			_log.Logger.Warnw("Stream error detected", "error", err)
//			if err := conn.Close(); err != nil { // Close connection immediately
//				_log.Logger.Warnw("Failed to close controller connection", "error", err)
//			}
//			time.Sleep(streamRecoveryInterval)
//			continue // Retry connecting
//		}
//	}
//	_log.Logger.Errorw("Max stream recovery attempts reached")
//}
