// Package logger provides structured logging functionality with different log levels.
// This demonstrates creating a custom logger and using pointers for configuration.
package logger

import (
	"cli-calculator/internal/constants"
	"fmt"
	"io"
	"os"
	"time"
)

// Logger represents a structured logger with configuration.
// It uses a pointer to LogConfig to demonstrate pointer usage in Go.
type Logger struct {
	config *LogConfig // Pointer to configuration
	output io.Writer  // Where to write logs (stdout, file, etc.)
}

// LogConfig holds logger configuration.
// Using pointers for optional fields is a common Go pattern.
type LogConfig struct {
	Level      constants.LogLevel // Minimum level to log
	TimeFormat string             // Time format for timestamps
	Prefix     string             // Optional prefix for log messages
	Enabled    bool               // Whether logging is enabled
}

// Global logger instance (package-level variable)
var defaultLogger *Logger

// init initializes the default logger when the package is imported.
func init() {
	defaultLogger = NewLogger(nil)
}

// NewLogger creates a new logger with the given configuration.
// If config is nil, it uses default settings.
// This demonstrates pointer parameters and nil checks.
func NewLogger(config *LogConfig) *Logger {
	// If config is nil, create default config
	if config == nil {
		config = &LogConfig{
			Level:      constants.LogLevelInfo,
			TimeFormat: "2006-01-02 15:04:05",
			Prefix:     constants.AppName,
			Enabled:    true,
		}
	}

	return &Logger{
		config: config,
		output: os.Stdout,
	}
}

// SetLevel changes the minimum log level.
func (l *Logger) SetLevel(level constants.LogLevel) {
	l.config.Level = level
}

// SetOutput changes the output writer.
func (l *Logger) SetOutput(w io.Writer) {
	l.output = w
}

// Enable enables or disables logging.
func (l *Logger) Enable(enabled bool) {
	l.config.Enabled = enabled
}

// log is the internal logging method.
func (l *Logger) log(level constants.LogLevel, format string, args ...interface{}) {
	// Check if logging is enabled and level is sufficient
	if !l.config.Enabled || level < l.config.Level {
		return
	}

	// Format timestamp
	timestamp := time.Now().Format(l.config.TimeFormat)

	// Format the message
	message := fmt.Sprintf(format, args...)

	// Build the log line
	logLine := fmt.Sprintf("[%s] [%s] [%s] %s\n",
		timestamp,
		l.config.Prefix,
		level.String(),
		message,
	)

	// Write to output
	fmt.Fprint(l.output, logLine)
}

// Debug logs a debug-level message.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(constants.LogLevelDebug, format, args...)
}

// Info logs an info-level message.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(constants.LogLevelInfo, format, args...)
}

// Warn logs a warning-level message.
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(constants.LogLevelWarn, format, args...)
}

// Error logs an error-level message.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(constants.LogLevelError, format, args...)
}

// Package-level functions that use the default logger

// Debug logs a debug message using the default logger.
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info logs an info message using the default logger.
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warn logs a warning message using the default logger.
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Error logs an error message using the default logger.
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// SetLevel sets the log level for the default logger.
func SetLevel(level constants.LogLevel) {
	defaultLogger.SetLevel(level)
}

// GetDefaultLogger returns the default logger instance.
// This allows users to configure the default logger if needed.
func GetDefaultLogger() *Logger {
	return defaultLogger
}
