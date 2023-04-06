package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/do"

	v1 "git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/controller/http/v1"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase/repo"
)

// Run sets all dependencies of the app and starts the http server
func Run(conf *domain.Config, ctx context.Context) {
	injector := do.New()

	// TODO remove
	do.ProvideValue(injector, conf)

	// Setup Logger
	do.Provide(injector, func(i *do.Injector) (usecase.LoggerRepo, error) {
		return repo.LoggerLogrusBuilder().
			WithColor(!conf.LOG.NoColor).
			WithFormat(conf.LOG.Format).
			WithLevel(conf.LOG.Level).
			New()
	})

	do.Provide(injector, func(i *do.Injector) (*repo.DatabaseEnt, error) {
		return repo.DatabaseEntBuilder().
			WithMaxOpenConns(conf.PG.PoolMax).
			WithPostgres(conf.PG.URL)
	})

	do.Provide(injector, func(i *do.Injector) (*repo.DatabaseEntImplemAuthorizationPolicy, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		return db.RepoAuthorizationPolicy(), nil
	})

	do.Provide(injector, func(i *do.Injector) (usecase.UserRepo, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		return db.RepoUser(), nil
	})

	do.Provide(injector, func(i *do.Injector) (usecase.UserToCompanyRepo, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		return db.RepoRoles(), nil
	})

	do.Provide(injector, func(i *do.Injector) (usecase.ItemBatchRepo, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		return db.RepoItemBatch(), nil
	})

	do.Provide(injector, func(i *do.Injector) (usecase.CompanyRepo, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		return db.RepoCompany(), nil
	})

	do.Provide(injector, func(i *do.Injector) (usecase.CertificationRepo, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		return db.RepoCertification(), nil
	})

	do.Provide(injector, func(i *do.Injector) (usecase.SessionStateRepo, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		return repo.SessionSCSBuilder().
			WithStore(db.RepoSessions()).
			NewSCSSessionRepo()
	})

	do.Provide(injector, func(i *do.Injector) (usecase.IAuthorizationEnforcerRepo, error) {
		db := do.MustInvoke[*repo.DatabaseEnt](i)
		auth, err := repo.AuthorizationEnforcerBuilder().
			WithAdapter(db.RepoAuthorizationPolicy()).
			New()

		if err != nil {
			return nil, err
		}
		auth.AddPolicyRolePathMethod("*", "/metrics", "*")
		auth.AddPolicyRolePathMethod("*", "/healthz", "*")
		auth.AddPolicyRolePathMethod("*", "/docs", "*")
		auth.AddPolicyRolePathMethod("*", "/docs/*", "*")
		auth.AddPolicyRolePathMethod("*", "/api/auth/*", "*")
		auth.AddPolicyRolePathMethod("user", "/api/v1/*", "*")

		return auth, nil
	})

	do.Provide(injector, usecase.NewLogger)
	do.Provide(injector, usecase.NewAuthBossLogger)
	do.Provide(injector, usecase.NewAuthBossServer)
	do.Provide(injector, usecase.NewAuthbossSession)
	do.Provide(injector, usecase.NewHTTPAuthorization)
	do.Provide(injector, usecase.NewCompany)
	do.Provide(injector, usecase.NewUserToCompany)
	do.Provide(injector, usecase.NewUser)
	do.Provide(injector, usecase.NewItemBatch)
	do.Provide(injector, usecase.NewCertification)

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
