package controller

import (
	"encoding/base64"
	"github.com/gorilla/sessions"
	"github.com/samber/do"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"time"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
)

func NewAuthenticator(i *do.Injector) (*authboss.Authboss, error) {
	logger := do.MustInvoke[*usecase.Logger](i).WithSubsystem("http_authenticator")
	conf := do.MustInvoke[*entity.Config](i)

	sessionStoreKey, _ := base64.StdEncoding.DecodeString(conf.HTTP.SessionStoreKey)

	sessionStore := abclientstate.NewSessionStorer(conf.App.Name, sessionStoreKey, nil)
	cstore := sessionStore.Store.(*sessions.CookieStore)
	cstore.Options.HttpOnly = false
	cstore.Options.Secure = false
	cstore.MaxAge(int((30 * 24 * time.Hour) / time.Second))

	ab := authboss.New()
	ab.Config.Storage.Server = do.MustInvoke[*usecase.AuthBossServer](i)
	ab.Config.Storage.SessionState = sessionStore

	ab.Config.Paths.Mount = "/auth"
	//TODO set this from the config
	ab.Config.Paths.RootURL = "http://localhost:" + conf.HTTP.Port

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
