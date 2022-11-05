package v1

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

	"github.com/MatthewBehnke/exampleGoApi/internal/app/domain"
	"github.com/MatthewBehnke/exampleGoApi/internal/app/http/middleware"
	"github.com/MatthewBehnke/exampleGoApi/internal/app/usecase"
	"github.com/MatthewBehnke/exampleGoApi/pkg/httpserver"
)

func NewHttpV1Router(i *do.Injector) (*HttpV1Router, error) {
	c := &HttpV1Router{
		config:            do.MustInvoke[*domain.Config](i),
		logger:            do.MustInvoke[*usecase.Logger](i).WithSubsystem("http server"),
		httpAuthorization: do.MustInvoke[*usecase.HttpAuthorization](i),
		httpV1:            do.MustInvoke[ServerInterface](i),
		authBoss:          do.MustInvoke[*authboss.Authboss](i),
	}
	return c, nil
}

type HttpV1Router struct {
	config            *domain.Config
	logger            *usecase.Logger
	httpV1            ServerInterface
	httpServer        *httpserver.Server
	httpAuthorization *usecase.HttpAuthorization
	authBoss          *authboss.Authboss
}

func (c *HttpV1Router) Shutdown() error {
	err := c.httpServer.Shutdown()
	if err != nil {
		return fmt.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}
	log.Print("controller service shutdown")
	return nil
}

func (c *HttpV1Router) Run() {
	hub := newHub()
	go hub.run()

	// HTTP Server
	mux := chi.NewRouter()
	mux.Use(chimiddleware.RequestID)
	mux.Use(chimiddleware.RealIP)
	mux.Use(middleware.NewStructuredLogger(c.logger))
	mux.Use(chimiddleware.Recoverer)

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/*", http.StripPrefix("/static/",
		http.FileServer(http.Dir("/")),
	))

	mux.Group(func(r chi.Router) {
		r.Use(c.authBoss.LoadClientStateMiddleware)
		r.Use(middleware.Authorizer(c.httpAuthorization, c.authBoss))

		r.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

		// Mount the router to a path (this should be the same as the Mount path above)
		// mux in this example is a chi router, but it could be anything that can route to
		// the Core.HttpV1Router.
		r.Group(func(mux chi.Router) {
			mux.Use(authboss.ModuleListMiddleware(c.authBoss))
			mux.Mount("/auth", http.StripPrefix("/auth", c.authBoss.Config.Core.Router))
		})

		// Protected Routes
		r.Group(func(r chi.Router) {
			r.Use(authboss.Middleware2(c.authBoss, authboss.RequireNone, authboss.RespondUnauthorized))

			// Openapi generated interfaces
			// HandlerFromMuxWithBaseURL(c.httpV1, r, "/v1")

			// r.Get("/v1/ok", func(w http.ResponseWriter, r *http.Request) {
			// 	w.WriteHeader(http.StatusOK)
			// })

			r.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
				serveWs(hub, w, r)
			})

		})

	})

	c.httpServer = httpserver.New(mux, httpserver.Port(c.config.HTTP.Port))
}
