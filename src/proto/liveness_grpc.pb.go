// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: liveness.proto

package proto

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
	LivenessService_Live_FullMethodName  = "/proto.LivenessService/Live"
	LivenessService_Ready_FullMethodName = "/proto.LivenessService/Ready"
)

// LivenessServiceClient is the client API for LivenessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LivenessServiceClient interface {
	Live(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	Ready(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
}

type livenessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLivenessServiceClient(cc grpc.ClientConnInterface) LivenessServiceClient {
	return &livenessServiceClient{cc}
}

func (c *livenessServiceClient) Live(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, LivenessService_Live_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livenessServiceClient) Ready(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, LivenessService_Ready_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LivenessServiceServer is the server API for LivenessService service.
// All implementations must embed UnimplementedLivenessServiceServer
// for forward compatibility.
type LivenessServiceServer interface {
	Live(context.Context, *LivenessRequest) (*LivenessResponse, error)
	Ready(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	mustEmbedUnimplementedLivenessServiceServer()
}

// UnimplementedLivenessServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLivenessServiceServer struct{}

func (UnimplementedLivenessServiceServer) Live(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Live not implemented")
}
func (UnimplementedLivenessServiceServer) Ready(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}
func (UnimplementedLivenessServiceServer) mustEmbedUnimplementedLivenessServiceServer() {}
func (UnimplementedLivenessServiceServer) testEmbeddedByValue()                         {}

// UnsafeLivenessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LivenessServiceServer will
// result in compilation errors.
type UnsafeLivenessServiceServer interface {
	mustEmbedUnimplementedLivenessServiceServer()
}

func RegisterLivenessServiceServer(s grpc.ServiceRegistrar, srv LivenessServiceServer) {
	// If the following call pancis, it indicates UnimplementedLivenessServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LivenessService_ServiceDesc, srv)
}

func _LivenessService_Live_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivenessServiceServer).Live(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivenessService_Live_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivenessServiceServer).Live(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivenessService_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivenessServiceServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivenessService_Ready_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivenessServiceServer).Ready(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LivenessService_ServiceDesc is the grpc.ServiceDesc for LivenessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LivenessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LivenessService",
	HandlerType: (*LivenessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Live",
			Handler:    _LivenessService_Live_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _LivenessService_Ready_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liveness.proto",
}
