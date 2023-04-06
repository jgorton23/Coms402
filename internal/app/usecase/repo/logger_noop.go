package repo

import (
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify loggerNoopImplem conforms to the required interfaces
var (
	_assertLoggerNoopRepoImplem                    = &loggerNoopImplem{}
	_                           usecase.LoggerRepo = _assertLoggerNoopRepoImplem
)

// LoggerNoopBuilder a logger that does nothing
func LoggerNoopBuilder() *loggerNoopBuilder {
	return &loggerNoopBuilder{}
}

type loggerNoopBuilder struct {
	// logger *loggerNoopImplem
}

// New creates a new usecase.LoggerRepo that does nothing
func (builder *loggerNoopBuilder) New() (usecase.LoggerRepo, error) {
	return &loggerNoopImplem{}, nil
}

type loggerNoopImplem struct{}

func (l *loggerNoopImplem) Trace(msg string, fields ...map[string]interface{}) {}

func (l *loggerNoopImplem) Debug(msg string, fields ...map[string]interface{}) {}

func (l *loggerNoopImplem) Info(msg string, fields ...map[string]interface{}) {}

func (l *loggerNoopImplem) Warn(msg string, fields ...map[string]interface{}) {}

func (l *loggerNoopImplem) Error(msg string, fields ...map[string]interface{}) {}
