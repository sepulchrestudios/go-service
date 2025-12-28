package log

import "errors"

// ErrCannotCreateLogger is a sentinel error representing a failure to create a logger.
var ErrCannotCreateLogger = errors.New("cannot create logger")

// ErrZapLoggerCannotBeNil is a sentinel error representing an attempt to use a nil zap-based logger pointer.
var ErrZapLoggerCannotBeNil = errors.New("zap logger instance cannot be nil")
