package usecase

import (
	"github.com/samber/do"
	"logur.dev/logur"
)

func NewLogger(logger logur.Logger) func(i *do.Injector) (Logger, error) {
	return func(i *do.Injector) (Logger, error) {
		return &loggerImplem{
			logger: logger,
		}, nil
	}
}

// loggerImplem wraps a logur logger and exposes it under a custom interface.
type loggerImplem struct {
	logger logur.Logger
}

// Trace logs a trace event.
func (l loggerImplem) Trace(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

// Debug logs a debug event.
func (l loggerImplem) Debug(msg string, fields ...map[string]interface{}) {
	l.logger.Debug(msg, fields...)
}

// Info logs an info event.
func (l loggerImplem) Info(msg string, fields ...map[string]interface{}) {
	l.logger.Info(msg, fields...)
}

// Warn logs a warning event.
func (l loggerImplem) Warn(msg string, fields ...map[string]interface{}) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error event.
func (l loggerImplem) Error(msg string, fields ...map[string]interface{}) {
	l.logger.Error(msg, fields...)
}

// WithFields annotates a logger with some context and it as a new instance.
func (l loggerImplem) WithFields(fields map[string]interface{}) Logger {
	return &loggerImplem{logger: logur.WithFields(l.logger, fields)}
}

// WithFields annotates a logger with some context and it as a new instance.
func (l loggerImplem) WithSubsystem(subsystem string) Logger {
	return &loggerImplem{logger: logur.WithFields(l.logger, map[string]interface{}{"subsystem": subsystem})}
}
