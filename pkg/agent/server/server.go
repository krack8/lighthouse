package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/config"
	_log "github.com/krack8/lighthouse/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func ConnectAndIdentifyWorker(ctx context.Context, controllerURL, secretName, resourceNamespace, groupName string, caCertPool *x509.CertPool) (*grpc.ClientConn, grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse], error) {
	maxAttempts := 30 // Maximum retry attempts
	retryInterval := 2 * time.Second
	var conn *grpc.ClientConn
	var err error
	var tlsConfig *tls.Config
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if config.IsInternalServer() {
			conn, err = grpc.NewClient(controllerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
		} else {
			tlsConfig = &tls.Config{}
			tlsConfig.InsecureSkipVerify = config.IsTlsInsecureSkipVerify()
			if caCertPool != nil {
				tlsConfig.RootCAs = caCertPool
			}
			conn, err = grpc.NewClient(controllerURL, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
		}
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
