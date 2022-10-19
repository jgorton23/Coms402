package repo

import (
	"context"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent"
)

// Service Interfaces
type (
	Service interface {
		HealthCheck() error
		Shutdown() error
	}

	ConfigService interface {
		Load(string) error
		Get() *entity.Config
	}

	DataBaseServiceUser interface {
		Get(context.Context) ([]entity.User, error)
		GetById(context.Context, int) (entity.User, error)
		GetByEmail(context.Context, string) (entity.User, error)
		Exists(context.Context, string) (bool, error)
		Create(context.Context, entity.User) (entity.User, error)
		Update(context.Context, entity.User) error
	}

	DatabaseService interface {
		Client() *ent.Client
		Service
	}
)
