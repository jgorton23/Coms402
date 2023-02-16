package repo

// import (
// 	"encoding/base64"
// 	"errors"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/securecookie"
// 	"github.com/gorilla/sessions"
// 	"github.com/samber/do"

// 	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
// 	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
// )

// // Pattern to verify sessionGorillaImplem conforms to the required interfaces
// var (
// 	_assertGorillaSessionRepoImplem                          = &sessionGorillaImplem{}
// 	_                               usecase.SessionStateRepo = _assertGorillaSessionRepoImplem
// )

// func NewGorillaSessionRepo(i *do.Injector) (usecase.SessionStateRepo, error) {
// 	conf := do.MustInvoke[*domain.Config](i)

// 	sessionStoreKey, _ := base64.StdEncoding.DecodeString(conf.HTTP.SessionStoreKey)

// 	cookieStore := sessions.NewCookieStore(sessionStoreKey)

// 	// 12 hours, set this to something because if we don't then sessions
// 	// may never expire as long as the browser remains opened.
// 	cookieStore.MaxAge(int((time.Hour * 12) / time.Second))
// 	cookieStore.Options.HttpOnly = true
// 	cookieStore.Options.Secure = true

// 	return &sessionGorillaImplem{
// 		store: cookieStore,
// 		name:  conf.App.Name,
// 	}, nil
// }

// type sessionGorillaImplem struct {
// 	name  string
// 	store sessions.Store
// }

// func (s sessionGorillaImplem) Get(r *http.Request) (map[interface{}]interface{}, error) {
// 	// Note that implementers of Get in gorilla all return a new session
// 	session, err := s.store.Get(r, s.name)
// 	if err != nil {
// 		e, ok := err.(securecookie.Error)
// 		if ok && !e.IsDecode() {
// 			// We ignore decoding errors, but nothing else
// 			return nil, err
// 		}

// 		// Get in gorilla does not return new sessions if a bad one exists
// 		// New() also happens to parse the cookie in r, and returns the same
// 		// decode error but still returns a new session
// 		session, _ = s.store.New(r, s.name)
// 		if session == nil {
// 			return nil, errors.New("could not create new session")
// 		}
// 	}

// 	return session.Values, nil
// }

// func (s sessionGorillaImplem) Save(r *http.Request, w http.ResponseWriter, values map[interface{}]interface{}) error {
// 	ses, err := s.store.Get(r, s.name)

// 	if err != nil {
// 		return err
// 	}

// 	ses.Values = values

// 	return s.store.Save(nil, w, ses)
// }
