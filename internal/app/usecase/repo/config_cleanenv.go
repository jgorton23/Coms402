package repo

import (
	"github.com/ilyakaznacheev/cleanenv"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify configCleanenvImplem conforms to the required interfaces
var (
	_assertConfigRepoImplem                    = &configCleanenvImplem{}
	_                       usecase.ConfigRepo = _assertConfigRepoImplem
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
