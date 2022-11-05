package repo

import (
	"github.com/MatthewBehnke/exampleGoApi/internal/app/usecase"
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/MatthewBehnke/exampleGoApi/internal/app/domain"
)

// Pattern to verify configCleanenvImplem conforms to the required interfaces
var (
	assertConfigRepoImplem                    = &configCleanenvImplem{}
	_                      usecase.ConfigRepo = assertConfigRepoImplem
)

func NewConfigRepo() usecase.ConfigRepo {
	return &configCleanenvImplem{}
}

type configCleanenvImplem struct {
	cfg *domain.Config
}

func (c *configCleanenvImplem) Load(path string) error {
	c.cfg = &domain.Config{}
	err := cleanenv.ReadConfig(path, c.cfg)
	if err != nil {
		return err
	}
	err = cleanenv.ReadEnv(c.cfg)
	if err != nil {
		return err
	}
	return nil
}

func (c *configCleanenvImplem) Get() *domain.Config {
	return c.cfg
}
