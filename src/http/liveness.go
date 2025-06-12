package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/sepulchrestudios/go-service/src/proto"
	"google.golang.org/grpc"
)

// ErrFailedLivenessCheck is a sentinel error representing a failed liveness check.
var ErrFailedLivenessCheck = errors.New("liveness check failed")

// ErrFailedReadinessCheck is a sentinel error representing a failed readiness check.
var ErrFailedReadinessCheck = errors.New("readiness check failed")

// LivenessResponseMessageSuccess represents the "success" message for the liveness endpoint.
const LivenessResponseMessageSuccess string = "ok"

// ReadinessResponseMessageSuccess represents the "success" message for the readiness endpoint.
const ReadinessResponseMessageSuccess string = LivenessResponseMessageSuccess

// LivenessServer represents the implementation of the "liveness" endpoints for gRPC and HTTP.
type LivenessServer struct {
	pb.UnimplementedLivenessServiceServer

	// LivenessFunction is the function that will execute when the liveness check runs. By default, this just returns
	// a success response with a nil error but it can be reassigned to a custom implementation per-service.
	LivenessFunction func() (*pb.LivenessResponse, error)

	// ReadinessFunction is the function that will execute when the readiness check runs. By default, this just returns
	// a success response with a nil error but it can be reassigned to a custom implementation per-service.
	ReadinessFunction func() (*pb.ReadinessResponse, error)
}

// NewLivenessServer creates and returns a new LivenessServer struct instance.
func NewLivenessServer() *LivenessServer {
	return &LivenessServer{
		LivenessFunction: func() (*pb.LivenessResponse, error) {
			return &pb.LivenessResponse{Message: LivenessResponseMessageSuccess}, nil
		},
		ReadinessFunction: func() (*pb.ReadinessResponse, error) {
			return &pb.ReadinessResponse{Message: ReadinessResponseMessageSuccess}, nil
		},
	}
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
	if l.LivenessFunction == nil {
		return &pb.LivenessResponse{Message: LivenessResponseMessageSuccess}, nil
	}
	resp, err := l.LivenessFunction()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedLivenessCheck, err)
	}
	return resp, nil
}

func (l *LivenessServer) Ready(ctx context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error) {
	if l.ReadinessFunction == nil {
		return &pb.ReadinessResponse{Message: ReadinessResponseMessageSuccess}, nil
	}
	resp, err := l.ReadinessFunction()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedReadinessCheck, err)
	}
	return resp, nil
}
