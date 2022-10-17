package api

import (
	"encoding/json"
	"net/http"

	"iseage/bank/internal/usecase"
	"iseage/bank/pkg/database/ent"
)

// Pattern used to verify UseCase conforms to required interfaces
var (
	assertHttpV1 = &HttpV1{}
	// Because of code generation the interface does not exist in
	// the usecase package
	_ ServerInterface = assertHttpV1
)

type HttpV1 struct {
	repo usecase.UserRepo
	log  usecase.LoggerAdapter
}

func NewHttpV1(r usecase.UserRepo, l usecase.LoggerAdapter) *HttpV1 {
	return &HttpV1{
		repo: r,
		log:  l,
	}
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
