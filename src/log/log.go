package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Contract is an interface that represents a zap-based logger.
type Contract interface {
	// Debug logs a message at debug-level.
	Debug(msg string, fields ...zap.Field)

	// DPanic logs a development panic message. If the logger is in development mode, a panic is also raised.
	DPanic(msg string, fields ...zap.Field)

	// Error logs a message at error-level.
	Error(msg string, fields ...zap.Field)

	// Fatal logs a message at fatal-level and then exits the application with os.Exit(1) to denote a failure.
	Fatal(msg string, fields ...zap.Field)

	// Info logs a message at info-level.
	Info(msg string, fields ...zap.Field)

	// IsUsingDebugMode returns a boolean describing whether "debug / development mode" is turned on for this logger.
	IsUsingDebugMode() bool

	// Log logs a message at the specified level with the specified fields.
	Log(lvl zapcore.Level, msg string, fields ...zap.Field)

	// Panic logs a message at panic-level, then a panic is raised.
	Panic(msg string, fields ...zap.Field)

	// Sync flushes any buffered log entries. This should be called before application exit.
	Sync() error

	// Warn logs a message at warn-level.
	Warn(msg string, fields ...zap.Field)
}
