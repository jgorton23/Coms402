package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/register"

	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/middleware"
	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/pkg/httpserver"
)

func New(i *do.Injector) (*Controller, error) {
	c := &Controller{
		config:            do.MustInvoke[*entity.Config](i),
		logger:            do.MustInvoke[*usecase.Logger](i).WithSubsystem("http server"),
		httpAuthorization: do.MustInvoke[*usecase.HttpAuthorization](i),
		httpV1:            do.MustInvoke[ServerInterface](i),
		authBoss:          do.MustInvoke[*authboss.Authboss](i),
	}
	return c, nil
}

type Controller struct {
	config            *entity.Config
	logger            *usecase.Logger
	httpV1            ServerInterface
	httpServer        *httpserver.Server
	httpAuthorization *usecase.HttpAuthorization
	authBoss          *authboss.Authboss
}

func (c *Controller) Shutdown() error {
	err := c.httpServer.Shutdown()
	if err != nil {
		return fmt.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}
	log.Print("controller service shutdown")
	return nil
}

func (c *Controller) Run() {
	// HTTP Server
	mux := chi.NewRouter()
	mux.Use(chimiddleware.RequestID)
	mux.Use(chimiddleware.RealIP)
	mux.Use(middleware.NewStructuredLogger(c.logger))
	mux.Use(chimiddleware.Recoverer)

	mux.Use(c.authBoss.LoadClientStateMiddleware)
	mux.Use(middleware.Authorizer(c.httpAuthorization, c.authBoss))

	mux.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

	// Mount the router to a path (this should be the same as the Mount path above)
	// mux in this example is a chi router, but it could be anything that can route to
	// the Core.Router.
	mux.Group(func(mux chi.Router) {
		mux.Use(authboss.ModuleListMiddleware(c.authBoss))
		mux.Mount("/auth", http.StripPrefix("/auth", c.authBoss.Config.Core.Router))
	})

	// Protected Routes
	mux.Group(func(r chi.Router) {
		r.Use(authboss.Middleware2(c.authBoss, authboss.RequireNone, authboss.RespondUnauthorized))

		// Openapi generated interfaces
		HandlerFromMuxWithBaseURL(c.httpV1, r, "/v1")
	})

	c.httpServer = httpserver.New(mux, httpserver.Port(c.config.HTTP.Port))
}
