// Package api implements routing paths. Each services in own file.
package api

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/defaults"

	"iseage/bank/config"
	v1 "iseage/bank/internal/delivery/controller/http/api/v1"
	"iseage/bank/internal/usecase"
)

const (
	sessionCookieName = "bank_app"
)

func NewRouter(cfg *config.Config, r chi.Router, logger usecase.LoggerAdapter, user usecase.UserUseCase) func(r chi.Router) {
	return func(r chi.Router) {

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

		ab.Config.Paths.RootURL = "http://localhost:" + cfg.HTTP.Port

		ab.Config.Storage.Server = &user
		ab.Config.Storage.SessionState = sessionStore
		ab.Config.Storage.CookieState = cookieStore

		ab.Config.Core.ViewRenderer = defaults.JSONRenderer{}

		// ab.Config.Modules.RegisterPreserveFields = []string{"email", "name"}

		defaults.SetCore(&ab.Config, true, false)

		emailRule := defaults.Rules{
			FieldName: "email", Required: true,
			MatchError: "Must be a valid e-mail address",
			MustMatch:  regexp.MustCompile(`.*@.*\.[a-z]+`),
		}

		passwordRule := defaults.Rules{
			FieldName: "password", Required: true,
			MinLength: 4,
		}
		nameRule := defaults.Rules{
			FieldName: "name", Required: true,
			MinLength: 2,
		}

		ab.Config.Core.BodyReader = defaults.HTTPBodyReader{
			ReadJSON: true,
			Rulesets: map[string][]defaults.Rules{
				"register":    {emailRule, passwordRule, nameRule},
				"recover_end": {passwordRule},
			},
			Confirms: map[string][]string{
				"register":    {"password", authboss.ConfirmPrefix + "password"},
				"recover_end": {"password", authboss.ConfirmPrefix + "password"},
			},
			Whitelist: map[string][]string{
				"register": {"email", "name", "password"},
			},
		}

		if err := ab.Init(); err != nil {
			log.Fatal(err.Error())
		}

		// reg := &register.Register{}
		// if err := reg.Init(ab); err != nil {
		// 	log.Fatal(err.Error())
		// }

		r.Use(ab.LoadClientStateMiddleware)
		// r.Use(remember.Middleware(ab))

		r.Group(func(mux chi.Router) {
			mux.Use(authboss.ModuleListMiddleware(ab))
			mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
		})

		optionsHandler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-CSRF-TOKEN", nosurf.Token(r))
			w.WriteHeader(http.StatusOK)
		}

		// We have to add each of the authboss get/post routes specifically because
		// chi sees the 'Mount' above as overriding the '/*' pattern.
		routes := []string{"login", "logout", "recover", "recover/end", "register"}
		r.MethodFunc("OPTIONS", "/*", optionsHandler)
		for _, route := range routes {
			r.MethodFunc("OPTIONS", "/auth/"+route, optionsHandler)
		}

		r.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

		r.Group(func(r chi.Router) {
			r.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized))
			// r.Use(lock.Middleware(ab))
			// r.Use(confirm.Middleware(ab))
			r.Route("/v1", v1.NewRouter)
		})

		walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			route = strings.Replace(route, "/*/", "/", -1)
			fmt.Printf("%s %s\n", method, route)
			return nil
		}

		if err := chi.Walk(r, walkFunc); err != nil {
			fmt.Printf("Logging err: %s\n", err.Error())
		}

	}
}
