package server

import (
	"context"

	pb "github.com/sepulchrestudios/go-service/src/proto"
)

// Server represents the interface for a liveness and readiness server.
type Server interface {
	// Live performs the liveness check.
	Live(ctx context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error)

	// MarkReady signals that the service is ready to receive traffic.
	MarkReady()

	// Ready performs the readiness check.
	Ready(ctx context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error)
}
