// Package app configures and runs application.
package app

import (
	"log"

	"github.com/samber/do"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"

	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/controller"
	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/controller/http/api"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
)

// Run creates objects via constructors.
func Run() {
	injector := do.New()

	config := repo.NewCleanEnvService()
	// do.Provide(injector, repo.NewCleanEnvService)
	// config := do.MustInvoke[*repo.CleanEnvService](injector)
	config.Load("./config.yml")
	do.ProvideValue(injector, config.Get())

	do.Provide(injector, repo.NewLogrusService())
	logrus := do.MustInvoke[logrusadapter.Logger](injector)

	do.Provide(injector, usecase.NewLoggerUseCase(&logrus))
	l := do.MustInvoke[*usecase.LoggerUseCase](injector)

	log.SetFlags(0) // Remove all flags
	// Override the global standard library logger to make sure everything uses our logger
	log.SetOutput(logur.NewLevelWriter(l, logur.Info))

	do.Provide(injector, repo.NewDatabaseService)
	do.Provide(injector, repo.NewDataBaseServiceUser)

	dbUser := do.MustInvoke[*repo.DataBaseServiceUser](injector)

	authBossUseCase := usecase.NewAuthBossUseCase(dbUser, *l)

	httpV1UseCase := api.NewHttpV1(dbUser, *l)

	controller.New(config.Get(), *l, *authBossUseCase, *httpV1UseCase)

	injector.Shutdown()
}
