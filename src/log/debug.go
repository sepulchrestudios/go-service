package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DebugLogger represents a struct that provides debug logging capabilities.
type DebugLogger struct {
	logger *zap.Logger
}

// Debug logs a message at debug-level.
func (l *DebugLogger) Debug(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Debug(msg, fields...)
}

// DPanic logs a development panic message. If the logger is in development mode, a panic is also raised.
func (l *DebugLogger) DPanic(msg string, fields ...zapcore.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.DPanic(msg, fields...)
}

// Sync flushes any buffered log entries. This should be called before application exit.
func (l *DebugLogger) Sync() error {
	if l == nil || l.logger == nil {
		return nil
	}
	return l.logger.Sync()
}

// WithOptions clones the existing logger, applies the provided options, and returns a new resultant logger.
func (l *DebugLogger) WithOptions(opts ...zap.Option) *DebugLogger {
	if l == nil || l.logger == nil {
		return nil
	}
	logger := l.logger.WithOptions(opts...)
	return &DebugLogger{
		logger: logger,
	}
}

// NewDebugLogger takes a set of zap options and returns a new DebugLogger instance as well as any error that may
// have occurred when attempting to create said logger.
func NewDebugLogger(options ...zap.Option) (*DebugLogger, error) {
	logger, err := zap.NewDevelopment(options...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateLogger, err)
	}
	return &DebugLogger{
		logger: logger,
	}, nil
}

// NewDebugLoggerFromZapLogger takes an existing zap logger pointer and returns a new DebugLogger instance plus any
// error that may have occurred.
func NewDebugLoggerFromZapLogger(logger *zap.Logger) (*DebugLogger, error) {
	if logger == nil {
		return nil, ErrZapLoggerCannotBeNil
	}
	return &DebugLogger{
		logger: logger,
	}, nil
}
