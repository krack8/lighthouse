package server

import (
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
	v1 "github.com/krack8/lighthouse/pkg/controller/grpc/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

func StartGrpcServer() {
	// Keep alive policy
	kaParams := keepalive.ServerParameters{
		MaxConnectionIdle:     5 * time.Minute,  // Allow idle connections for 5 minutes
		MaxConnectionAge:      0,                // No forced connection expiration
		MaxConnectionAgeGrace: 1 * time.Minute,  // Grace period before closing connection
		Time:                  30 * time.Second, // Send keepalive ping every 30s (adjust if needed)
		Timeout:               10 * time.Second, // Wait 10s for keepalive response
	}

	s := grpc.NewServer(grpc.KeepaliveParams(kaParams))

	pb.RegisterControllerServer(s, &v1.ControllerServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		log.Logger.Fatalf("Failed to listen: %v", err)
	}

	log.Logger.Infof("GRPC Server started on :%s", "50051")
	if err := s.Serve(lis); err != nil {
		log.Logger.Fatalf("Failed to serve: %v", err)
	}
}
