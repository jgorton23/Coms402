// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/casbin/casbin/v2/model"
	"github.com/volatiletech/authboss/v3"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

// usecase interfaces
type (
	// Logger is a generic logging interface
	Logger interface {
		Trace(msg string, fields ...map[string]interface{})
		Debug(msg string, fields ...map[string]interface{})
		Info(msg string, fields ...map[string]interface{})
		Warn(msg string, fields ...map[string]interface{})
		Error(msg string, fields ...map[string]interface{})

		// WithFields annotates a logger with some context and it as a new instance.
		WithFields(fields map[string]interface{}) Logger
		WithSubsystem(subsystem string) Logger
	}

	// Conforms to the authboss interfaces
	AuthBoss interface {
		Load(context.Context, string) (authboss.User, error)
		Save(context.Context, authboss.User) error
		New(context.Context) authboss.User
		Create(context.Context, authboss.User) error
	}

	Casbin interface {
		LoadPolicy(model model.Model) error
		SavePolicy(model model.Model) error
		AddPolicy(string, string, []string) error
		RemovePolicy(string, string, []string) error
		RemoveFilteredPolicy(string, string, int, ...string) error
	}
)
