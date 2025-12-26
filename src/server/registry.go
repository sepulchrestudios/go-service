package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/sepulchrestudios/go-service/src/proto"
	"google.golang.org/grpc"
)

// RegisterLivenessServer takes a gRPC server registrar and a LivenessServer pointer, then registers them together.
func RegisterLivenessServer(grpcServer grpc.ServiceRegistrar, livenessServer pb.LivenessServiceServer) {
	pb.RegisterLivenessServiceServer(grpcServer, livenessServer)
}

// RegisterLivenessServerHandlers takes a context, mux, and gRPC client connection and registers the gateway handlers.
// Returns any error that may have occurred during the process.
func RegisterLivenessServerHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterLivenessServiceHandler(ctx, mux, conn)
}
