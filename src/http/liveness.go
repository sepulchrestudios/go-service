package http

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/sepulchrestudios/go-service/src/proto"
	"google.golang.org/grpc"
)

// LivenessResponseMessageSuccess represents the "success" message for the liveness endpoints.
const LivenessResponseMessageSuccess string = "ok"

// LivenessServer represents the implementation of the "liveness" endpoints for gRPC and HTTP.
type LivenessServer struct {
	pb.UnimplementedLivenessServiceServer
}

// NewLivenessServer creates and returns a new LivenessServer struct instance.
func NewLivenessServer() *LivenessServer {
	return &LivenessServer{}
}

// RegisterLivenessServer takes a gRPC server registrar and a LivenessServer pointer, then registers them together.
func RegisterLivenessServer(grpcServer grpc.ServiceRegistrar, livenessServer pb.LivenessServiceServer) {
	pb.RegisterLivenessServiceServer(grpcServer, livenessServer)
}

// RegisterLivenessServerHandlers takes a context, mux, and gRPC client connection and registers the gateway handlers.
// Returns any error that may have occurred during the process.
func RegisterLivenessServerHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterLivenessServiceHandler(ctx, mux, conn)
}

func (l *LivenessServer) Live(ctx context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error) {
	return &pb.LivenessResponse{Message: LivenessResponseMessageSuccess}, nil
}

func (l *LivenessServer) Ready(ctx context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error) {
	return l.Live(ctx, req)
}
