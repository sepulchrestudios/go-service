package service

// LivenessServiceInterface represents the interface that all liveness service implementations must satisfy.
type LivenessServiceInterface interface {
	// DoLivenessCheck performs the liveness check.
	DoLivenessCheck() ([]byte, error)

	// DoMarkReady signals that the service is ready to receive traffic.
	DoMarkReady() error

	// DoReadinessCheck performs the readiness check.
	DoReadinessCheck() ([]byte, error)
}
