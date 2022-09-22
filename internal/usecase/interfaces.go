// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"iseage/bank/internal/entity"
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

	User interface {
		Create(context.Context, entity.User) (entity.User, error)
	}

	UserRepo interface {
		Exists(context.Context, entity.User) (bool, error)
		Create(context.Context, entity.User) (entity.User, error)
	}

	// Translation -.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) ([]entity.Translation, error)
	}

	// TranslationRepo -.
	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	// TranslationWebAPI -.
	TranslationWebAPI interface {
		Translate(entity.Translation) (entity.Translation, error)
	}
)
