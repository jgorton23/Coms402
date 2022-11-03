package usecase

import (
	"encoding/base64"
	"time"

	"github.com/gorilla/sessions"
	"github.com/samber/do"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"

	"github.com/MatthewBehnke/exampleGoApi/internal/domain"
)

func NewHttpAuthenticator(i *do.Injector) (*authboss.Authboss, error) {
	logger := do.MustInvoke[*Logger](i).WithSubsystem("http_authenticator")
	conf := do.MustInvoke[*domain.Config](i)

	sessionStoreKey, _ := base64.StdEncoding.DecodeString(conf.HTTP.SessionStoreKey)

	sessionStore := abclientstate.NewSessionStorer(conf.App.Name, sessionStoreKey, nil)
	cstore := sessionStore.Store.(*sessions.CookieStore)
	cstore.Options.HttpOnly = false
	cstore.Options.Secure = false
	cstore.MaxAge(int((30 * 24 * time.Hour) / time.Second))

	ab := authboss.New()
	ab.Config.Storage.Server = do.MustInvoke[*AuthBossServer](i)
	ab.Config.Storage.SessionState = sessionStore

	ab.Config.Paths.Mount = "/auth"
	//TODO set this from the config
	ab.Config.Paths.RootURL = "http://localhost:" + conf.HTTP.Port

	ab.Config.Core.ViewRenderer = defaults.JSONRenderer{}
	ab.Config.Core.MailRenderer = defaults.JSONRenderer{}
	defaults.SetCore(&ab.Config, true, false)
	ab.Config.Core.Logger = do.MustInvoke[*AuthBossLogger](i)

	if err := ab.Init(); err != nil {
		return nil, err
	}

	logger.Info("http authentication service started")

	return ab, nil
}
