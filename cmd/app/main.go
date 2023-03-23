package main

import (
	"context"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase/repo"
)

func main() {

	conf := repo.NewConfigRepo()
	conf.Load("config.yml")

	app.Run(conf.Get(), context.Background())
}
