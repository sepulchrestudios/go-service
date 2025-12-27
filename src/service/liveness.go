package service

import (
	"errors"
	"fmt"
)

// ErrFailedLivenessCheck is a sentinel error representing a failed liveness check.
var ErrFailedLivenessCheck = errors.New("liveness check failed")

// ErrFailedReadinessCheck is a sentinel error representing a failed readiness check.
var ErrFailedReadinessCheck = errors.New("readiness check failed")

// LivenessResponseMessageSuccess represents the "success" message for the liveness endpoint.
const LivenessResponseMessageSuccess string = "ok"

// ReadinessResponseMessageSuccess represents the "success" message for the readiness endpoint.
const ReadinessResponseMessageSuccess string = LivenessResponseMessageSuccess

// LivenessService represents the implementation of the "liveness" service functionality.
type LivenessService struct {
	// LivenessFunction is the function that will execute when the liveness check runs. By default, this just returns
	// a success response with a nil error but it can be reassigned to a custom implementation per-service.
	LivenessFunction func() ([]byte, error)

	// ReadinessFunction is the function that will execute when the readiness check runs. By default, this just returns
	// a success response with a nil error but it can be reassigned to a custom implementation per-service.
	ReadinessFunction func() ([]byte, error)

	// readinessChannel is the channel used to signal readiness (i.e. a response from ReadinessFunction).
	readinessChannel chan struct{}
}

// NewLivenessService creates and returns a new LivenessService struct instance.
func NewLivenessService() *LivenessService {
	readinessChannel := make(chan struct{})
	return &LivenessService{
		LivenessFunction: func() ([]byte, error) {
			return []byte(LivenessResponseMessageSuccess), nil
		},
		ReadinessFunction: func() ([]byte, error) {
			<-readinessChannel
			return []byte(ReadinessResponseMessageSuccess), nil
		},
		readinessChannel: readinessChannel,
	}
}

// DoLivenessCheck performs the liveness check by invoking LivenessFunction. If no custom function is provided, it
// returns a success response by default.
func (l *LivenessService) DoLivenessCheck() ([]byte, error) {
	if l.LivenessFunction == nil {
		return []byte(LivenessResponseMessageSuccess), nil
	}
	resp, err := l.LivenessFunction()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedLivenessCheck, err)
	}
	return resp, nil
}

// DoMarkReady signals that the service is ready to receive traffic.
func (l *LivenessService) DoMarkReady() error {
	close(l.readinessChannel)
	return nil
}

// DoReadinessCheck performs the readiness check by invoking ReadinessFunction. If no custom function is provided, it
// returns a success response by default.
func (l *LivenessService) DoReadinessCheck() ([]byte, error) {
	if l.ReadinessFunction == nil {
		return []byte(ReadinessResponseMessageSuccess), nil
	}
	resp, err := l.ReadinessFunction()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedReadinessCheck, err)
	}
	return resp, nil
}
