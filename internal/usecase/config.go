package usecase

import (
	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
	"github.com/samber/do"
)

var (
	configUserUseCase        = &ConfigUseCase{}
	_                 Config = configUserUseCase
)

type ConfigUseCase struct {
	config *repo.CleanEnvService
	Path   string
}

func NewConfigUseCase(path string) func(i *do.Injector) (*ConfigUseCase, error) {
	return func(i *do.Injector) (*ConfigUseCase, error) {
		config := do.MustInvoke[*repo.CleanEnvService](i)
		cs := ConfigUseCase{config: config}
		err := cs.config.Load(path)
		if err != nil {
			return nil, err
		}
		return &cs, nil
	}
}

func (cs ConfigUseCase) Get() *entity.Config {
	return cs.config.Get()
}

func (cs ConfigUseCase) HealthCheck() error {
	return nil
}

func (cs ConfigUseCase) Shutdown() error {
	return nil
}
