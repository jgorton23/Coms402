package usecase

import (
	"github.com/samber/do"
	"logur.dev/logur"
)

func NewLogger(i *do.Injector) (*Logger, error) {
	return &Logger{
		do.MustInvoke[LoggerRepo](i),
	}, nil
}

// Logger wraps a LoggerRepo and exposes it.
type Logger struct {
	logger LoggerRepo
}

// Trace logs a trace event.
func (l Logger) Trace(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}

// Debug logs a debug event.
func (l Logger) Debug(msg string, fields ...map[string]interface{}) {
	l.logger.Debug(msg, fields...)
}

// Info logs an info event.
func (l Logger) Info(msg string, fields ...map[string]interface{}) {
	l.logger.Info(msg, fields...)
}

// Warn logs a warning event.
func (l Logger) Warn(msg string, fields ...map[string]interface{}) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error event.
func (l Logger) Error(msg string, fields ...map[string]interface{}) {
	l.logger.Error(msg, fields...)
}

// WithFields annotates a loggerRepo with some context and returns it as a new instance.
func (l Logger) WithFields(fields map[string]interface{}) *Logger {
	return &Logger{logger: logur.WithFields(l.logger, fields)}
}

// WithSubsystem annotates a loggerRepo with a subsystem and returns it as a new instance.
func (l Logger) WithSubsystem(subsystem string) *Logger {
	return &Logger{logger: logur.WithFields(l.logger, map[string]interface{}{"subsystem": subsystem})}
}
