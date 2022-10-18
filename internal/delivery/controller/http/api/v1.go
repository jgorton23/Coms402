package api

import (
	"encoding/json"
	"net/http"

	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent"
	"github.com/samber/do"
)

// Pattern used to verify UseCase conforms to required interfaces
var (
	assertHttpV1 = &HttpV1{}
	// Because of code generation the interface does not exist in
	// the usecase package
	_ ServerInterface = assertHttpV1
)

func NewHttpV1(i *do.Injector) (*HttpV1, error) {
	httpV1 := &HttpV1{}
	httpV1.log = do.MustInvoke[*usecase.LoggerUseCase](i).WithSubsystem("controller http v1")
	httpV1.repo = do.MustInvoke[*repo.DataBaseServiceUser](i)

	return httpV1, nil
}

type HttpV1 struct {
	repo usecase.DataBaseServiceUser
	log  usecase.Logger
}

func (v1 HttpV1) ShowUserById(w http.ResponseWriter, r *http.Request, userId UserId) {
	u, err := v1.repo.GetByEmail(r.Context(), userId)

	if err != nil {
		if ent.IsNotFound(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// generic error
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(u)
}

func (v1 HttpV1) ListUsers(w http.ResponseWriter, r *http.Request, params ListUsersParams) {
	users, err := v1.repo.Get(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(users)
}
