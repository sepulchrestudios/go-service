package server

import (
	"context"

	pb "github.com/sepulchrestudios/go-service/src/proto"
	"github.com/sepulchrestudios/go-service/src/service"
)

// LivenessServer represents the implementation of the "liveness" endpoints for gRPC and HTTP.
type LivenessServer struct {
	pb.UnimplementedLivenessServiceServer

	// livenessService is the underlying liveness service implementation.
	livenessService service.LivenessServiceInterface
}

// NewLivenessServer creates and returns a new LivenessServer struct instance.
func NewLivenessServer(livenessService service.LivenessServiceInterface) *LivenessServer {
	return &LivenessServer{
		livenessService: livenessService,
	}
}

// Live performs the liveness check. Returns the response message and any error that may have occurred.
func (l *LivenessServer) Live(ctx context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error) {
	if l.livenessService == nil {
		return nil, ErrNilLivenessService
	}
	resp, err := l.livenessService.DoLivenessCheck()
	return &pb.LivenessResponse{Message: string(resp)}, err
}

// MarkReady signals that the service is ready to receive traffic.
func (l *LivenessServer) MarkReady() error {
	return l.livenessService.DoMarkReady()
}

// Ready performs the readiness check. Returns the response message and any error that may have occurred.
func (l *LivenessServer) Ready(ctx context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error) {
	if l.livenessService == nil {
		return nil, ErrNilLivenessService
	}
	resp, err := l.livenessService.DoReadinessCheck()
	return &pb.ReadinessResponse{Message: string(resp)}, err
}
