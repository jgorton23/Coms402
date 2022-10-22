package controller

import (
	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
)

func NewAuthenticator(i *do.Injector) (*authboss.Authboss, error) {
	log := do.MustInvoke[usecase.Logger](i).WithSubsystem("http authenticator")
	conf := do.MustInvoke[*entity.Config](i)

	ab := authboss.New()
	ab.Config.Storage.Server = do.MustInvoke[usecase.AuthBossServer](i)
	ab.Config.Storage.SessionState = do.MustInvoke[*repo.SessionStorerService](i)
	ab.Config.Storage.CookieState = do.MustInvoke[*repo.CookieStorerService](i)

	ab.Config.Paths.Mount = "/auth"
	//TODO set this from the config
	ab.Config.Paths.RootURL = "http://localhost:" + conf.HTTP.Port

	ab.Config.Core.ViewRenderer = defaults.JSONRenderer{}
	ab.Config.Core.MailRenderer = defaults.JSONRenderer{}
	defaults.SetCore(&ab.Config, true, false)
	ab.Config.Core.Logger = do.MustInvoke[usecase.AuthBossLogger](i)

	if err := ab.Init(); err != nil {
		return nil, err
	}

	log.Info("http authentication started")

	return ab, nil
}
