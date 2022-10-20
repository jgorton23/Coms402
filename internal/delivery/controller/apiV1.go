package controller

import (
	"encoding/json"
	"net/http"

	"github.com/samber/do"

	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/middleware"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent"
)

// Pattern used to verify UseCase conforms to required interfaces
var (
	assertHttpV1 = &httpV1Implem{}
	// Because of code generation the interface does not exist in
	// the usecase package
	_ ServerInterface = assertHttpV1
)

func NewHttpV1(i *do.Injector) (ServerInterface, error) {
	httpV1 := &httpV1Implem{
		log:  do.MustInvoke[usecase.Logger](i).WithSubsystem("controller http v1"),
		repo: do.MustInvoke[repo.DataBaseServiceUser](i),
	}
	return httpV1, nil
}

type httpV1Implem struct {
	repo repo.DataBaseServiceUser
	log  usecase.Logger
}

func (v1 httpV1Implem) ShowUserById(w http.ResponseWriter, r *http.Request, userId UserId) {
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

func (v1 httpV1Implem) ListUsers(w http.ResponseWriter, r *http.Request, params ListUsersParams) {
	users, err := v1.repo.Get(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	logger := middleware.GetLogEntry(r)

	logger.Info("test")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(users)
}
