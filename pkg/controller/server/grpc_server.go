package server

import (
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/pb"
	v1 "github.com/krack8/lighthouse/pkg/controller/grpc/api/v1"
	"github.com/krack8/lighthouse/pkg/log"
	"google.golang.org/grpc"
	"net"
)

func StartGrpcServer() {
	s := grpc.NewServer()

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
