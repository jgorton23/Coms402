package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	abrenderer "github.com/volatiletech/authboss-renderer"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/volatiletech/authboss/v3/lock"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/recover"
	_ "github.com/volatiletech/authboss/v3/register"
	"github.com/volatiletech/authboss/v3/remember"

	"iseage/bank/config"
	v1 "iseage/bank/internal/delivery/controller/http/api/v1"
	"iseage/bank/internal/delivery/middleware"
	"iseage/bank/internal/usecase"
	"iseage/bank/pkg/httpserver"
)

const (
	sessionCookieName = "bank_app"
)

func New(cfg *config.Config, logger usecase.LoggerAdapter, abuc usecase.AuthBossUseCase) {
	// HTTP Server
	mux := chi.NewRouter()
	mux.Use(chimiddleware.RequestID)
	mux.Use(chimiddleware.RealIP)
	mux.Use(chimiddleware.Recoverer)
	mux.Use(middleware.NewStructuredLogger(logger))

	// loading usecase's onto context
	mux.Use(middleware.LoggerCtx(logger))

	// // Mounting the backend
	// handler.Route("/api", api.NewRouter(cfg, l, *userUseCase))

	// // Mounting the frontend
	// handler.Route("/", www.NewRouter)

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

	defaults.SetCore(&ab.Config, true, false)

	if err := ab.Init(); err != nil {
		logger.Error(err.Error())
	}

	mux.Use(ab.LoadClientStateMiddleware, remember.Middleware(ab))

	mux.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

	mux.Group(func(r chi.Router) {
		r.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized))
		r.Use(lock.Middleware(ab))
		r.Route("/v1", v1.NewRouter)
	})

	// Mount the router to a path (this should be the same as the Mount path above)
	// mux in this example is a chi router, but it could be anything that can route to
	// the Core.Router.
	mux.Group(func(mux chi.Router) {
		mux.Use(authboss.ModuleListMiddleware(ab))
		mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		// route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(mux, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	httpServer := httpserver.New(mux, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
