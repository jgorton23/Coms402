package usecase

import (
	"net/http"
	"strings"

	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

// Pattern to verify AuthbossSession conforms to the required interfaces
var (
	_assertAuthbossSession                                = &AuthbossSession{}
	_                      authboss.ClientStateReadWriter = _assertAuthbossSession
)

func NewAuthbossSession(i *do.Injector) (*AuthbossSession, error) {
	conf := do.MustInvoke[*domain.Config](i)

	return &AuthbossSession{
		name:         conf.App.Name,
		sessionsRepo: do.MustInvoke[SessionStateRepo](i),
	}, nil
}

type AuthbossSession struct {
	name         string
	sessionsRepo SessionStateRepo
}

// Basically Load in session world

func (s AuthbossSession) ReadState(r *http.Request) (authboss.ClientState, error) {
	values, err := s.sessionsRepo.Get(r)

	if err != nil {
		return nil, err
	}

	cs := &sessionState{
		values: values,
		r:      r,
	}

	return cs, nil
}

// Basically Save in session world

func (s AuthbossSession) WriteState(w http.ResponseWriter, state authboss.ClientState, ev []authboss.ClientStateEvent) error {
	// This should never be nil (despite what authboss.ClientStateReadWriter
	// interface says) because all Get methods return a new session in gorilla.
	// In cases where Get returns an error, we ensure we create a new session
	ses := state.(*sessionState)

	for _, ev := range ev {
		switch ev.Kind {
		case authboss.ClientStateEventPut:
			ses.values[ev.Key] = ev.Value
		case authboss.ClientStateEventDel:
			delete(ses.values, ev.Key)

		case authboss.ClientStateEventDelAll:
			if len(ev.Key) == 0 {
				for _, key := range ses.values {
					delete(ses.values, key)
				}
			} else {
				whitelist := strings.Split(ev.Key, ",")
				for key := range ses.values {
					if k, ok := key.(string); ok {
						dontDelete := false
						for _, w := range whitelist {
							if w == k {
								dontDelete = true

								break
							}
						}

						if !dontDelete {
							delete(ses.values, key)
						}
					}
				}
			}
		}
	}

	return s.sessionsRepo.Save(ses.r, w, ses.values)
}

// sessionState is an authboss.ClientState implementation that
// holds the request's session values for the duration of the request.
type sessionState struct {
	values map[interface{}]interface{}
	r      *http.Request
}

// Get a key from the session
func (s sessionState) Get(key string) (string, bool) {
	str, ok := s.values[key]
	if !ok {
		return "", false
	}

	value := str.(string)

	return value, ok
}
