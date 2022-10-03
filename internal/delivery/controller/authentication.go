package controller

import (
	"encoding/base64"
	"time"

	"github.com/gorilla/sessions"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	abrenderer "github.com/volatiletech/authboss-renderer"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"

	"iseage/bank/config"
	"iseage/bank/internal/usecase"
)

const (
	sessionCookieName = "bank_app"
)

func newAuthentication(cfg *config.Config, abuc usecase.AuthBossUseCase, logger usecase.LoggerAdapter) *authboss.Authboss {
	ab := authboss.New()

	cookieStoreKey, _ := base64.StdEncoding.DecodeString(cfg.HTTP.CookieStoreKey)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(cfg.HTTP.SessionStoreKey)

	cookieStore := abclientstate.NewCookieStorer(cookieStoreKey, nil)
	cookieStore.HTTPOnly = false
	cookieStore.Secure = false

	sessionStore := abclientstate.NewSessionStorer(sessionCookieName, sessionStoreKey, nil)
	cstore := sessionStore.Store.(*sessions.CookieStore)
	cstore.Options.HttpOnly = false
	cstore.Options.Secure = false
	cstore.MaxAge(int((30 * 24 * time.Hour) / time.Second))

	ab.Config.Storage.Server = abuc
	ab.Config.Storage.SessionState = sessionStore
	ab.Config.Storage.CookieState = cookieStore

	ab.Config.Paths.Mount = "/auth"
	ab.Config.Paths.RootURL = "http://localhost:" + cfg.HTTP.Port

	ab.Config.Core.ViewRenderer = defaults.JSONRenderer{}
	ab.Config.Core.MailRenderer = abrenderer.NewEmail("/auth", "ab_views")

	//TODO switch to using a custom logger so the logs can be more consistent
	defaults.SetCore(&ab.Config, true, false)

	if err := ab.Init(); err != nil {
		logger.Error(err.Error())
	}
	return ab
}