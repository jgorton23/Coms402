package usecase

import (
	"github.com/samber/do"
	"logur.dev/logur"
)

func NewAuthbossLogger(i *do.Injector) (AuthBossLogger, error) {
	return &authbossLoggerImplem{
		logger: do.MustInvoke[Logger](i).WithSubsystem("http authenticator"),
	}, nil
}

// authbossLoggerImplem wraps a logur logger and exposes it under a custom interface.
type authbossLoggerImplem struct {
	logger logur.Logger
}

// Info logs an info event.
func (l authbossLoggerImplem) Info(msg string) {
	l.logger.Info(msg)
}

// Error logs an error event.
func (l authbossLoggerImplem) Error(msg string) {
	l.logger.Error(msg)
}
