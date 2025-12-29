package log

import "errors"

// ErrCannotCreateDebugLogger is a sentinel error representing a failure to create a debug logger.
var ErrCannotCreateDebugLogger = errors.New("cannot create debug logger")

// ErrCannotCreateLogger is a sentinel error representing a failure to create a logger.
var ErrCannotCreateLogger = errors.New("cannot create logger")

// ErrZapLoggerCannotBeNil is a sentinel error representing an attempt to use a nil zap-based logger pointer.
var ErrZapLoggerCannotBeNil = errors.New("zap logger instance cannot be nil")
