// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: pkg/common/pb/controller.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Controller_TaskStream_FullMethodName = "/pb.Controller/TaskStream"
)

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerClient interface {
	TaskStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[TaskStreamRequest, TaskStreamResponse], error)
}

type controllerClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerClient(cc grpc.ClientConnInterface) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) TaskStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[TaskStreamRequest, TaskStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Controller_ServiceDesc.Streams[0], Controller_TaskStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TaskStreamRequest, TaskStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Controller_TaskStreamClient = grpc.BidiStreamingClient[TaskStreamRequest, TaskStreamResponse]

// ControllerServer is the server API for Controller service.
// All implementations must embed UnimplementedControllerServer
// for forward compatibility.
type ControllerServer interface {
	TaskStream(grpc.BidiStreamingServer[TaskStreamRequest, TaskStreamResponse]) error
	mustEmbedUnimplementedControllerServer()
}

// UnimplementedControllerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedControllerServer struct{}

func (UnimplementedControllerServer) TaskStream(grpc.BidiStreamingServer[TaskStreamRequest, TaskStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method TaskStream not implemented")
}
func (UnimplementedControllerServer) mustEmbedUnimplementedControllerServer() {}
func (UnimplementedControllerServer) testEmbeddedByValue()                    {}

// UnsafeControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerServer will
// result in compilation errors.
type UnsafeControllerServer interface {
	mustEmbedUnimplementedControllerServer()
}

func RegisterControllerServer(s grpc.ServiceRegistrar, srv ControllerServer) {
	// If the following call pancis, it indicates UnimplementedControllerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Controller_ServiceDesc, srv)
}

func _Controller_TaskStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerServer).TaskStream(&grpc.GenericServerStream[TaskStreamRequest, TaskStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Controller_TaskStreamServer = grpc.BidiStreamingServer[TaskStreamRequest, TaskStreamResponse]

// Controller_ServiceDesc is the grpc.ServiceDesc for Controller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Controller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TaskStream",
			Handler:       _Controller_TaskStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/common/pb/controller.proto",
}
