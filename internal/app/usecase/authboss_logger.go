package usecase

import (
	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"
)

var (
	assertAuthBossLogger                 = &AuthBossLogger{}
	_                    authboss.Logger = assertAuthBossLogger
)

func NewAuthBossLogger(i *do.Injector) (*AuthBossLogger, error) {
	return &AuthBossLogger{
		logger: do.MustInvoke[*Logger](i).WithSubsystem("http authenticator"),
	}, nil
}

// AuthBossLogger wraps a Logger and exposes it under a custom interface.
type AuthBossLogger struct {
	logger *Logger
}

// Info logs an info event.
func (l AuthBossLogger) Info(msg string) {
	l.logger.Info(msg)
}

// Error logs an error event.
func (l AuthBossLogger) Error(msg string) {
	l.logger.Error(msg)
}
