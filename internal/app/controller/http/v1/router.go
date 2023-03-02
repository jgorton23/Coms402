package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/controller/http/middleware"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/httpserver"

	_ "github.com/volatiletech/authboss/v3/auth"
	_ "github.com/volatiletech/authboss/v3/logout"
	_ "github.com/volatiletech/authboss/v3/register"
)

func NewHttpV1Router(i *do.Injector) (*HttpV1Router, error) {
	c := &HttpV1Router{
		config: do.MustInvoke[*domain.Config](i),
		logger: do.MustInvoke[*usecase.Logger](i).
			WithSubsystem("http server"),
		httpAuthorization: do.MustInvoke[*usecase.HTTPAuthorization](i),
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
	httpAuthorization *usecase.HTTPAuthorization
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
	// HTTP Server
	mux := chi.NewRouter()
	mux.Use(chimiddleware.RequestID)
	mux.Use(chimiddleware.RealIP)
	mux.Use(middleware.NewStructuredLogger(c.logger))
	mux.Use(chimiddleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	// mux.Handle("/static/*", http.StripPrefix("/static/",
	// 	// http.FileServer(http.Dir("static")),
	// ))

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
			HandlerFromMuxWithBaseURL(c.httpV1, r, "/v1")

			r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello world"))
			})
		})
	})

	c.httpServer = httpserver.New(mux, httpserver.Port(c.config.HTTP.Port))
}
