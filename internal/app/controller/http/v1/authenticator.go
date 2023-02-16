package http

import (
	"os"

	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

func NewHttpAuthenticator(i *do.Injector) (*authboss.Authboss, error) {
	logger := do.MustInvoke[*usecase.Logger](i).WithSubsystem("http_authenticator")

	ab := authboss.New()
	ab.Config.Storage.Server = do.MustInvoke[*usecase.AuthBossServer](i)
	ab.Config.Storage.SessionState = do.MustInvoke[*usecase.AuthbossSession](i)

	ab.Config.Paths.Mount = "/auth"
	// TODO set from env file....
	ab.Config.Paths.RootURL = "http://" + os.Getenv("APP_HOST")

	ab.Config.Core.ViewRenderer = defaults.JSONRenderer{}
	ab.Config.Core.MailRenderer = defaults.JSONRenderer{}
	defaults.SetCore(&ab.Config, true, false)
	ab.Config.Core.Logger = do.MustInvoke[*usecase.AuthBossLogger](i)

	if err := ab.Init(); err != nil {
		return nil, err
	}

	logger.Info("http authentication service started")

	return ab, nil
}
