package v1

import (
	"net/http"

	"github.com/samber/do"

	"github.com/MatthewBehnke/apis/internal/app/usecase"
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
		logger: do.MustInvoke[*usecase.Logger](i).WithSubsystem("controller http v1"),
		//userRepo: do.MustInvoke[domain.UserRepo](i),
	}
	return httpV1, nil
}

type httpV1Implem struct {
	//userRepo domain.UserRepo
	logger *usecase.Logger
}

func (v1 httpV1Implem) ShowUserById(w http.ResponseWriter, r *http.Request, userId UserId) {
	panic("Not Implemented")

	//u, err := v1.userRepo.GetByEmail(r.Context(), userId)
	//
	//if err != nil {
	//	if ent.IsNotFound(err) {
	//		w.WriteHeader(http.StatusNotFound)
	//		return
	//	}
	//	// generic error
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusFound)
	//json.NewEncoder(w).Encode(u)
}

func (v1 httpV1Implem) ListUsers(w http.ResponseWriter, r *http.Request, params ListUsersParams) {
	panic("Not Implemented")
	//users, err := v1.userRepo.Get(r.Context())
	//
	//if err != nil {
	//	w.WriteHeader(http.StatusNotFound)
	//	return
	//}
	//
	//logger := middleware.GetLogEntry(r)
	//
	//logger.Info("test")
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusFound)
	//json.NewEncoder(w).Encode(users)
}
