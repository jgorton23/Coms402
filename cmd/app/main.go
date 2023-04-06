package main

import (
	"context"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

func main() {

	cfg := &domain.Config{}
	err := cleanenv.ReadConfig("config.yml", cfg)

	if err != nil {
		os.Exit(-1)
	}

	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		os.Exit(-1)
	}

	app.Run(cfg, context.Background())
}
