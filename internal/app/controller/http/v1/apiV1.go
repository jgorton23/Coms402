package http

import (
	"net/http"

	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern used to verify UseCase conforms to required interfaces
var (
	_assertHttpV1 = &httpV1Implem{}
	// Because of code generation the interface does not exist in
	// the usecase package
	_ ServerInterface = _assertHttpV1
)

func NewHttpV1(i *do.Injector) (ServerInterface, error) {
	httpV1 := &httpV1Implem{
		logger: do.MustInvoke[*usecase.Logger](i).WithSubsystem("controller http v1"),
		// userRepo: do.MustInvoke[domain.UserRepo](i),
	}

	return httpV1, nil
}

type httpV1Implem struct {
	// userRepo domain.UserRepo
	logger *usecase.Logger
}

func (v1 httpV1Implem) ShowUserById(w http.ResponseWriter, r *http.Request, userId UserId) {
	panic("Not Implemented")
}

func (v1 httpV1Implem) ListUsers(w http.ResponseWriter, r *http.Request, params ListUsersParams) {
	panic("Not Implemented")
}
