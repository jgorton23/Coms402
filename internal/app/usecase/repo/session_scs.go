package repo

import (
	"context"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify loggerLogrusImplem conforms to the required interfaces
var (
	_assertSessionSCSImplem                          = &sessionSCSImplem{}
	_                       usecase.SessionStateRepo = _assertSessionSCSImplem
)

func SessionSCSBuilder() *sessionSCSBuilder {
	return &sessionSCSBuilder{
		opts: struct {
			Lifetime     time.Duration
			IdleTimeout  time.Duration
			CookieName   string
			CookieDomain string
			Store        scs.Store
		}{
			Lifetime:     3 * time.Hour,
			IdleTimeout:  20 * time.Minute,
			CookieName:   "session_id",
			CookieDomain: "localhost",
		},
	}
}

type sessionSCSBuilder struct {
	opts struct {
		Lifetime     time.Duration
		IdleTimeout  time.Duration
		CookieName   string
		CookieDomain string
		Store        scs.Store
	}
}

func (builder *sessionSCSBuilder) WithLifetime(lifetime time.Duration) *sessionSCSBuilder {
	builder.opts.Lifetime = lifetime
	return builder
}

func (builder *sessionSCSBuilder) WithIdleTimeout(timeout time.Duration) *sessionSCSBuilder {
	builder.opts.IdleTimeout = timeout
	return builder
}

func (builder *sessionSCSBuilder) WithCookieName(name string) *sessionSCSBuilder {
	builder.opts.CookieName = name
	return builder
}
func (builder *sessionSCSBuilder) WithCookieDomain(domain string) *sessionSCSBuilder {
	builder.opts.CookieDomain = domain
	return builder
}

func (builder *sessionSCSBuilder) WithStore(store scs.Store) *sessionSCSBuilder {
	builder.opts.Store = store
	return builder
}

func (builder *sessionSCSBuilder) NewSCSSessionRepo() (usecase.SessionStateRepo, error) {

	implem := &sessionSCSImplem{}

	implem.sessionManager = scs.New()
	implem.sessionManager.Store = builder.opts.Store

	implem.sessionManager.Lifetime = builder.opts.Lifetime
	implem.sessionManager.IdleTimeout = builder.opts.IdleTimeout
	implem.sessionManager.Cookie.Name = builder.opts.CookieName
	implem.sessionManager.Cookie.Domain = builder.opts.CookieDomain
	implem.sessionManager.Cookie.HttpOnly = true
	implem.sessionManager.Cookie.Path = "/"
	implem.sessionManager.Cookie.Persist = false
	implem.sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	implem.sessionManager.Cookie.Secure = true

	return implem, nil
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
