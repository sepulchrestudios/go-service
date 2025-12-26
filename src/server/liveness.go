package server

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/sepulchrestudios/go-service/src/proto"
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

	// readinessChannel is the channel used to signal readiness (i.e. a response from ReadinessFunction).
	readinessChannel chan struct{}
}

// NewLivenessServer creates and returns a new LivenessServer struct instance.
func NewLivenessServer() *LivenessServer {
	readinessChannel := make(chan struct{})
	return &LivenessServer{
		LivenessFunction: func() (*pb.LivenessResponse, error) {
			return &pb.LivenessResponse{Message: LivenessResponseMessageSuccess}, nil
		},
		ReadinessFunction: func() (*pb.ReadinessResponse, error) {
			<-readinessChannel
			return &pb.ReadinessResponse{Message: ReadinessResponseMessageSuccess}, nil
		},
		readinessChannel: readinessChannel,
	}
}

// Live performs the liveness check by invoking LivenessFunction. If no custom function is provided, it returns
// a success response by default.
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

// MarkReady signals that the service is ready to receive traffic.
func (l *LivenessServer) MarkReady() {
	close(l.readinessChannel)
}

// Ready performs the readiness check by invoking ReadinessFunction. If no custom function is provided, it returns
// a success response by default.
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
