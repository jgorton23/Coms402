package middleware

import (
	"fmt"
	"net/http"

	"github.com/volatiletech/authboss/v3"

	"github.com/MatthewBehnke/apis/internal/app/usecase"
)

// Authorizer Authz is a middleware that controls the access to the HTTP service, it is based
// on Casbin, which supports access control models like ACL, RBAC, ABAC.
// The plugin determines whether to allow a request based on (role, path, method).
// role: the authenticated role.
// path: the URL for the requested resource.
// method: one of HTTP methods like GET, POST, PUT, DELETE.
//
// This middleware should be inserted fairly early in the middleware stack to
// protect subsequent layers. All the denied requests will not go further.
//
// It's notable that this middleware should be behind the authentication (e.g.,
// HTTP basic authentication, OAuth), so this plugin can get the logged-in role
// to perform the authorization.
func Authorizer(httpAuthorization *usecase.HTTPAuthorization, a *authboss.Authboss) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := GetLogEntry(r).WithSubsystem("http authorization")

			user := ""

			// Try to load user
			u, err := a.CurrentUser(r)

			if err != nil {
				if err.Error() != "user not found" {
					GetLogEntry(r).Warn(fmt.Sprintf("user not found: %v", err.Error()))
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				}

				user = "anonymous"
			} else {
				user = u.GetPID()
			}

			method := r.Method
			path := r.URL.Path

			ok, err := httpAuthorization.EnforceUser(user, path, method)

			if err != nil {
				log.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			if ok {
				log.Info(fmt.Sprintf("user: %v, allowed on: %v %v", user, method, path))
				next.ServeHTTP(w, r)
			} else {
				log.Info(fmt.Sprintf("user: %v, denied on: %v %v", user, method, path))
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			}
		}

		return http.HandlerFunc(fn)
	}
}
