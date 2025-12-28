package server

import "errors"

// ErrNilLivenessService is a sentinel error representing a nil liveness service.
var ErrNilLivenessService = errors.New("liveness service is nil")
