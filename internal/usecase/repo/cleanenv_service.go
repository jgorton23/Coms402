package repo

import (
	"fmt"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/samber/do"
)

func NewCleanEnvService(i *do.Injector) (*CleanEnvService, error) {
	return &CleanEnvService{}, nil
}

type CleanEnvService struct {
	cfg *entity.Config
}

func (c *CleanEnvService) Load(path string) error {
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

func (c *CleanEnvService) Get() *entity.Config {
	return c.cfg
}
