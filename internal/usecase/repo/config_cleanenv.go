package repo

import (
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
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
	cfg *entity.Config
}

func (c *configCleanenvImplem) Load(path string) error {
	c.cfg = &entity.Config{}
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

func (c *configCleanenvImplem) Get() *entity.Config {
	return c.cfg
}
