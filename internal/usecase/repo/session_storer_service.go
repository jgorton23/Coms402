package repo

import (
	"encoding/base64"
	"errors"
	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"
	"net/http"
	"strings"
	"time"
)

func NewSessionStorerService(i *do.Injector) (*SessionStorerService, error) {
	conf := do.MustInvoke[*entity.Config](i)

	sessionStoreKey, _ := base64.StdEncoding.DecodeString(conf.HTTP.SessionStoreKey)

	cookieStore := sessions.NewCookieStore(sessionStoreKey)

	// 12 hours, set this to something because if we don't then sessions
	// may never expire as long as the browser remains opened.
	cookieStore.MaxAge(int((time.Hour * 12) / time.Second))
	cookieStore.Options.HttpOnly = false
	cookieStore.Options.Secure = false

	return &SessionStorerService{
		Name:  conf.App.Name,
		Store: cookieStore,
	}, nil

}

type SessionStorerService struct {
	Name string

	sessions.Store
}

// ReadState loads the session from the request context
func (s SessionStorerService) ReadState(r *http.Request) (authboss.ClientState, error) {
	// Note that implementers of Get in gorilla all return a new session
	session, err := s.Store.Get(r, s.Name)
	if err != nil {
		e, ok := err.(securecookie.Error)
		if ok && !e.IsDecode() {
			// We ignore decoding errors, but nothing else
			return nil, err
		}

		// Get in gorilla does not return new sessions if a bad one exists
		// New() also happens to parse the cookie in r, and returns the same
		// decode error but still returns a new session
		session, err = s.Store.New(r, s.Name)
		if session == nil {
			return nil, errors.New("could not create new session")
		}
	}

	cs := &sessionState{
		session: session,
	}

	return cs, nil
}

// WriteState to the responsewriter
func (s SessionStorerService) WriteState(w http.ResponseWriter, state authboss.ClientState, ev []authboss.ClientStateEvent) error {
	// This should never be nil (despite what authboss.ClientStateReadWriter
	// interface says) because all Get methods return a new session in gorilla.
	// In cases where Get returns an error, we ensure we create a new session
	ses := state.(*sessionState)

	for _, ev := range ev {
		switch ev.Kind {
		case authboss.ClientStateEventPut:
			ses.session.Values[ev.Key] = ev.Value
		case authboss.ClientStateEventDel:
			delete(ses.session.Values, ev.Key)

		case authboss.ClientStateEventDelAll:
			if len(ev.Key) == 0 {
				// Delete the entire session
				ses.session.Options.MaxAge = -1
			} else {
				whitelist := strings.Split(ev.Key, ",")
				for key := range ses.session.Values {
					if k, ok := key.(string); ok {

						dontDelete := false
						for _, w := range whitelist {
							if w == k {
								dontDelete = true
								break
							}
						}

						if !dontDelete {
							delete(ses.session.Values, key)
						}
					}
				}
			}
		}
	}

	return s.Store.Save(nil, w, ses.session)
}

// sessionState is an authboss.ClientState implementation that
// holds the request's session values for the duration of the request.
type sessionState struct {
	session *sessions.Session
}

// Get a key from the session
func (s sessionState) Get(key string) (string, bool) {
	str, ok := s.session.Values[key]
	if !ok {
		return "", false
	}
	value := str.(string)

	return value, ok
}
