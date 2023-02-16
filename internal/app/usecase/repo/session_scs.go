package repo

import (
	"context"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify loggerLogrusImplem conforms to the required interfaces
var (
	_assertSCSSessionRepoImplem                          = &sessionSCSImplem{}
	_                           usecase.SessionStateRepo = _assertSCSSessionRepoImplem
)

func NewSCSSessionRepo(i *do.Injector) (usecase.SessionStateRepo, error) {
	// config := do.MustInvoke[*domain.Config](i)
	sessionManager := scs.New()
	sessionManager.Store = do.MustInvoke[scs.Store](i)
	sessionManager.Lifetime = 3 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	sessionManager.Cookie.Name = "session_id"
	sessionManager.Cookie.Domain = "localhost"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Path = "/"
	sessionManager.Cookie.Persist = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = true

	return &sessionSCSImplem{
		sessionManager: sessionManager,
	}, nil
}

type sessionSCSImplem struct {
	sessionManager *scs.SessionManager
}

func (s sessionSCSImplem) Get(r *http.Request) (out map[interface{}]interface{}, err error) {
	ctx, err := s.loadCookie(r)
	if err != nil {
		return nil, err
	}

	out = make(map[interface{}]interface{})

	for _, v := range s.sessionManager.Keys(ctx) {
		out[v] = s.sessionManager.Get(ctx, v)
	}

	return out, nil
}

func (s sessionSCSImplem) Save(r *http.Request, w http.ResponseWriter, values map[interface{}]interface{}) error {
	ctx, err := s.loadCookie(r)
	if err != nil {
		return err
	}

	for k, v := range values {
		s.sessionManager.Put(ctx, k.(string), v)
	}

	switch s.sessionManager.Status(ctx) {
	case scs.Modified:
		token, expiry, err := s.sessionManager.Commit(ctx)
		if err != nil {
			return err
		}

		s.sessionManager.WriteSessionCookie(ctx, w, token, expiry)
	case scs.Destroyed:
		s.sessionManager.WriteSessionCookie(ctx, w, "", time.Time{})
	case scs.Unmodified:
	default:
		break
	}

	return nil
}

func (s sessionSCSImplem) loadCookie(r *http.Request) (context.Context, error) {
	var token string

	cookie, err := r.Cookie(s.sessionManager.Cookie.Name)

	if err == nil {
		token = cookie.Value
	}

	ctx, err := s.sessionManager.Load(r.Context(), token)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}
