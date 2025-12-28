package service

import "errors"

// ErrFailedLivenessCheck is a sentinel error representing a failed liveness check.
var ErrFailedLivenessCheck = errors.New("liveness check failed")

// ErrFailedReadinessCheck is a sentinel error representing a failed readiness check.
var ErrFailedReadinessCheck = errors.New("readiness check failed")
