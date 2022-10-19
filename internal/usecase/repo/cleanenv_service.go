package repo

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
)

// func NewCleanEnvService(i *do.Injector) (*CleanEnvService, error) {
// 	return &CleanEnvService{}, nil
// }

func NewCleanEnvService() ConfigService {
	return &cleanEnvServiceImplem{}
}

type cleanEnvServiceImplem struct {
	cfg *entity.Config
}

func (c *cleanEnvServiceImplem) Load(path string) error {
	c.cfg = &entity.Config{}
	err := cleanenv.ReadConfig(path, c.cfg)
	if err != nil {
		return fmt.Errorf("CleanEnv error: %w", err)
	}
	err = cleanenv.ReadEnv(c.cfg)
	if err != nil {
		return err
	}
	return nil
}

func (c *cleanEnvServiceImplem) Get() *entity.Config {
	return c.cfg
}
