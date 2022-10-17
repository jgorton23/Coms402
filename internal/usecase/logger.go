package usecase

import (
	"github.com/samber/do"
	"logur.dev/logur"
)

type LoggerFields map[string]interface{}

// LoggerUseCase wraps a logur logger and exposes it under a custom interface.
type LoggerUseCase struct {
	logger logur.Logger
}

func NewLoggerUseCase(logger logur.Logger) func(i *do.Injector) (*LoggerUseCase, error) {
	return func(i *do.Injector) (*LoggerUseCase, error) {
		return &LoggerUseCase{
			logger: logger,
		}, nil
	}
}

// Trace logs a trace event.
func (l LoggerUseCase) Trace(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

// Debug logs a debug event.
func (l LoggerUseCase) Debug(msg string, fields ...map[string]interface{}) {
	l.logger.Debug(msg, fields...)
}

// Info logs an info event.
func (l LoggerUseCase) Info(msg string, fields ...map[string]interface{}) {
	l.logger.Info(msg, fields...)
}

// Warn logs a warning event.
func (l LoggerUseCase) Warn(msg string, fields ...map[string]interface{}) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error event.
func (l LoggerUseCase) Error(msg string, fields ...map[string]interface{}) {
	l.logger.Error(msg, fields...)
}

// WithFields annotates a logger with some context and it as a new instance.
func (l LoggerUseCase) WithFields(fields map[string]interface{}) Logger {
	return &LoggerUseCase{logger: logur.WithFields(l.logger, fields)}
}

// WithFields annotates a logger with some context and it as a new instance.
func (l LoggerUseCase) WithSubsystem(subsystem string) Logger {
	return &LoggerUseCase{logger: logur.WithFields(l.logger, map[string]interface{}{"subsystem": subsystem})}
}
