// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/casbin/casbin/v2/model"

	"github.com/MatthewBehnke/apis/internal/app/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// ConfigRepo -
	ConfigRepo interface {
		Load(string) error
		Get() *domain.Config
	}

	// AuthorizationPolicyRepo -
	AuthorizationPolicyRepo interface {
		// LoadPolicy loads all policy rules from the storage.
		LoadPolicy(model model.Model) error
		// SavePolicy saves all policy rules to the storage.
		SavePolicy(model model.Model) error

		// AddPolicy adds a policy rule to the storage.
		// This is part of the Auto-Save feature.
		AddPolicy(sec string, ptype string, rule []string) error
		// RemovePolicy removes a policy rule from the storage.
		// This is part of the Auto-Save feature.
		RemovePolicy(sec string, ptype string, rule []string) error
		// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
		// This is part of the Auto-Save feature.
		RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
	}

	// UserRepo -
	UserRepo interface {
		Get(context.Context) ([]domain.User, error)
		GetById(context.Context, int) (domain.User, error)
		GetByEmail(context.Context, string) (domain.User, error)
		Exists(context.Context, string) (bool, error)
		Create(context.Context, domain.User) (domain.User, error)
		Update(context.Context, domain.User) error
	}

	// LoggerRepo -
	LoggerRepo interface {
		Trace(msg string, fields ...map[string]interface{})
		Debug(msg string, fields ...map[string]interface{})
		Info(msg string, fields ...map[string]interface{})
		Warn(msg string, fields ...map[string]interface{})
		Error(msg string, fields ...map[string]interface{})
	}

	// AuthorizationEnforcerRepo -
	AuthorizationEnforcerRepo interface {
		EnforceRolePathMethod(role, path, method string) (bool, error)
		ReloadPolicy() error
	}
)
