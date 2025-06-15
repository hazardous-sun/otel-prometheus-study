package logger

import (
	"context"
	"log/slog"
	"os"
)

var (
	// AppLogger is the main logger instance
	AppLogger *slog.Logger
)

func init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// Use TextHandler (human-readable); use JSONHandler for structured logs
	AppLogger = slog.New(slog.NewTextHandler(os.Stdout, opts))

	// Optional: override default logger
	slog.SetDefault(AppLogger)
}

// LogDebug
// Logs a debug message with optional key-value pairs
func LogDebug(msg string, attrs ...any) {
	AppLogger.Debug(msg, attrs...)
}

// LogInfo
// Logs an info message
func LogInfo(msg string, attrs ...any) {
	AppLogger.Info(msg, attrs...)
}

// LogWarning
// Logs a warning message
func LogWarning(msg string, attrs ...any) {
	AppLogger.Warn(msg, attrs...)
}

// LogError
// Logs an error message
func LogError(err error, attrs ...any) {
	AppLogger.Error("error occurred", append([]any{"error", err.Error()}, attrs...)...)
}

// LogSuccess
// Is just a wrapper for info with a "success=true" flag
func LogSuccess(msg string, attrs ...any) {
	AppLogger.Info(msg, append(attrs, "success", true)...)
}

// WithContext
// Returns a logger bound to a context
func WithContext(ctx context.Context) *slog.Logger {
	return slog.Default()
}
