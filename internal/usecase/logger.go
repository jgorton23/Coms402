package usecase

import (
	"logur.dev/logur"
)

type LoggerFields map[string]interface{}

// LoggerAdapter wraps a logur logger and exposes it under a custom interface.
type LoggerAdapter struct {
	logger logur.Logger
}

// NewLoggerAdapter returns a new Logger instance.
func NewLoggerAdapter(logger logur.Logger) *LoggerAdapter {
	return &LoggerAdapter{
		logger: logger,
	}
}

// Trace logs a trace event.
func (l *LoggerAdapter) Trace(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

// Debug logs a debug event.
func (l *LoggerAdapter) Debug(msg string, fields ...map[string]interface{}) {
	l.logger.Debug(msg, fields...)
}

// Info logs an info event.
func (l *LoggerAdapter) Info(msg string, fields ...map[string]interface{}) {
	l.logger.Info(msg, fields...)
}

// Warn logs a warning event.
func (l *LoggerAdapter) Warn(msg string, fields ...map[string]interface{}) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error event.
func (l *LoggerAdapter) Error(msg string, fields ...map[string]interface{}) {
	l.logger.Error(msg, fields...)
}

// WithFields annotates a logger with some context and it as a new instance.
func (l *LoggerAdapter) WithFields(fields map[string]interface{}) Logger {
	return &LoggerAdapter{logger: logur.WithFields(l.logger, fields)}
}
