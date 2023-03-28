package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"

	"github.com/samber/do"

	v1 "git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/controller/http/v1"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase/repo"
)

func Run(conf *domain.Config, ctx context.Context) {
	injector := do.New()

	do.ProvideValue(injector, conf)
	do.Provide(injector, repo.NewLoggerRepo)
	do.Provide(injector, repo.NewDatabaseConnection)
	do.Provide(injector, repo.NewUserRepo)
	do.Provide(injector, repo.NewAuthorizationPolicyRepo)

	do.Provide(injector, repo.NewAuthorizationEnforcer)
	auth := do.MustInvoke[usecase.IAuthorizationEnforcerRepo](injector)
	auth.AddPolicyRolePathMethod("*", "/metrics", "*")
	auth.AddPolicyRolePathMethod("*", "/healthz", "*")
	auth.AddPolicyRolePathMethod("*", "/docs", "*")
	auth.AddPolicyRolePathMethod("*", "/docs/*", "*")
	auth.AddPolicyRolePathMethod("*", "/api/auth/*", "*")
	auth.AddPolicyRolePathMethod("user", "/api/v1/*", "*")

	do.Provide(injector, repo.NewSCSRepo)
	do.Provide(injector, repo.NewSCSSessionRepo)
	do.Provide(injector, repo.NewUserToCompanyRepo)
	do.Provide(injector, repo.NewCompanyRepo)
	do.Provide(injector, repo.NewItemBatchRepo)

	do.Provide(injector, usecase.NewLogger)
	do.Provide(injector, usecase.NewAuthBossLogger)
	do.Provide(injector, usecase.NewAuthBossServer)
	do.Provide(injector, usecase.NewAuthbossSession)
	do.Provide(injector, usecase.NewHTTPAuthorization)
	do.Provide(injector, usecase.NewCompany)
	do.Provide(injector, usecase.NewUserToCompany)
	do.Provide(injector, usecase.NewUser)
	do.Provide(injector, usecase.NewItemBatch)

	// HTTP stuff
	do.Provide(injector, v1.NewHttpAuthenticator)
	do.Provide(injector, v1.NewHttpV1)
	do.Provide(injector, v1.NewHttpV1Router)

	c := do.MustInvoke[*v1.HttpV1Router](injector)
	c.Run(injector)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Print("app - Run - signal: " + s.String())
	case <-ctx.Done():
		log.Print("Stopping")
	}

	// Shutdown
	err := injector.Shutdown()
	if err != nil {
		log.Print(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}

}
