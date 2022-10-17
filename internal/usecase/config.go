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
	logger Logger
}

func NewConfigUseCase(path string, la Logger) func(i *do.Injector) (*ConfigUseCase, error) {
	return func(i *do.Injector) (*ConfigUseCase, error) {
		la.Info("starting subsystem")
		config := do.MustInvoke[*repo.CleanEnvService](i)
		cs := ConfigUseCase{config: config, logger: la}
		err := cs.config.Load(path)
		if err != nil {
			return nil, err
		}
		cs.logger.Info("started subsystem")
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
	cs.logger.Info("shutdown config")
	return nil
}
