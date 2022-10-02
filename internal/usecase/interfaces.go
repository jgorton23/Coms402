// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"iseage/bank/internal/entity"

	"github.com/volatiletech/authboss/v3"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Logger -.
	Logger interface {
		Trace(msg string, fields ...map[string]interface{})
		Debug(msg string, fields ...map[string]interface{})
		Info(msg string, fields ...map[string]interface{})
		Warn(msg string, fields ...map[string]interface{})
		Error(msg string, fields ...map[string]interface{})

		// WithFields annotates a logger with some context and it as a new instance.
		WithFields(fields map[string]interface{}) Logger
	}

	// Conforms to the authboss interfaces
	AuthBoss interface {
		Load(context.Context, string) (authboss.User, error)
		Save(context.Context, authboss.User) error
		New(context.Context) authboss.User
		Create(context.Context, authboss.User) error
	}

	UserRepo interface {
		GetById(context.Context, int) (entity.User, error)
		GetByEmail(context.Context, string) (entity.User, error)
		Exists(context.Context, string) (bool, error)
		Create(context.Context, entity.User) (entity.User, error)
		Update(context.Context, entity.User) error
	}
)
