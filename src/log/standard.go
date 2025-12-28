package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// StandardLogger represents a struct that provides logging capabilities.
type StandardLogger struct {
	debugMode bool
	logger    *zap.Logger
}

// NewLogger takes a boolean describing whether "debug / development mode" should be turned on, plus a set of zap
// options, and returns a new Logger instance as well as any error that may have occurred when attempting to create
// said logger.
func NewStandardLogger(shouldUseDebugMode bool, options ...zap.Option) (*StandardLogger, error) {
	var logger *zap.Logger
	var err error
	if shouldUseDebugMode {
		logger, err = zap.NewDevelopment(options...)
	} else {
		logger, err = zap.NewProduction(options...)
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateLogger, err)
	}
	return &StandardLogger{
		debugMode: shouldUseDebugMode,
		logger:    logger,
	}, nil
}

// NewStandardLoggerFromZapLogger takes an existing zap logger pointer and returns a new Logger instance plus any error
// that may have occurred.
func NewStandardLoggerFromZapLogger(logger *zap.Logger) (*StandardLogger, error) {
	if logger == nil {
		return nil, ErrZapLoggerCannotBeNil
	}
	return &StandardLogger{
		debugMode: logger.Core().Enabled(zapcore.DebugLevel),
		logger:    logger,
	}, nil
}

// Debug logs a message at debug-level.
func (l *StandardLogger) Debug(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Debug(msg, fields...)
}

// DPanic logs a development panic message. If the logger is in development mode, a panic is also raised.
func (l *StandardLogger) DPanic(msg string, fields ...zapcore.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.DPanic(msg, fields...)
}

// Error logs a message at error-level.
func (l *StandardLogger) Error(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Error(msg, fields...)
}

// Fatal logs a message at fatal-level and then exits the application with os.Exit(1) to denote a failure.
func (l *StandardLogger) Fatal(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Fatal(msg, fields...)
}

// GetZapCore returns the zapcore.Core implementation used within the logger.
func (l *StandardLogger) GetZapCore() zapcore.Core {
	if l == nil || l.logger == nil {
		return nil
	}
	return l.logger.Core()
}

// GetZapLogger returns the zap.Logger implementation used within the logger.
func (l *StandardLogger) GetZapLogger() *zap.Logger {
	if l == nil {
		return nil
	}
	return l.logger
}

// Info logs a message at info-level.
func (l *StandardLogger) Info(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Info(msg, fields...)
}

// IsUsingDebugMode returns a boolean describing whether "debug / development mode" is turned on for this logger.
func (l *StandardLogger) IsUsingDebugMode() bool {
	if l == nil {
		return false
	}
	return l.debugMode
}

// Log logs a message at the specified level with the specified fields.
func (l *StandardLogger) Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Log(lvl, msg, fields...)
}

// Panic logs a message at panic-level, then a panic is raised.
func (l *StandardLogger) Panic(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Panic(msg, fields...)
}

// Sync flushes any buffered log entries. This should be called before application exit.
func (l *StandardLogger) Sync() error {
	if l == nil || l.logger == nil {
		return nil
	}
	return l.logger.Sync()
}

// Warn logs a message at warn-level.
func (l *StandardLogger) Warn(msg string, fields ...zap.Field) {
	if l == nil || l.logger == nil {
		return
	}
	l.logger.Warn(msg, fields...)
}

// WithOptions clones the existing logger, applies the provided options, and returns a new resultant logger.
func (l *StandardLogger) WithOptions(opts ...zap.Option) *StandardLogger {
	if l == nil || l.logger == nil {
		return nil
	}
	return &StandardLogger{
		debugMode: l.debugMode,
		logger:    l.logger.WithOptions(opts...),
	}
}
