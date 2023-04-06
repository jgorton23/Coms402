package http

import (
	"encoding/json"
	"net/http"

	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
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
		logger:               do.MustInvoke[*usecase.Logger](i).WithSubsystem("controller http v1"),
		companyUseCase:       do.MustInvoke[*usecase.Company](i),
		userUseCase:          do.MustInvoke[*usecase.User](i),
		userToCompanyUseCase: do.MustInvoke[*usecase.UserToCompany](i),
		itemBatchUseCase:     do.MustInvoke[*usecase.ItemBatch](i),
		certificationUseCase: do.MustInvoke[*usecase.Certification](i),
		itemToItemUseCase:    do.MustInvoke[*usecase.ItemToItem](i),
	}

	return httpV1, nil
}

type httpV1Implem struct {
	companyUseCase       *usecase.Company
	userUseCase          *usecase.User
	itemBatchUseCase     *usecase.ItemBatch
	userToCompanyUseCase *usecase.UserToCompany
	certificationUseCase *usecase.Certification
	itemToItemUseCase    *usecase.ItemToItem

	logger                *usecase.Logger
	authbossAuthenticator *authboss.Authboss
}

// Helper Functions //
func (v1 httpV1Implem) loadDomainUser(r *http.Request) (domain.User, error) {
	user, err := v1.authbossAuthenticator.CurrentUser(r)

	if err != nil {
		return domain.User{}, err
	}

	// Get the user email from authboss then query the database for the complete user so we can
	// get the user id to create a company with
	domainUser, err := v1.userUseCase.FindByEmail(r.Context(), user.GetPID())

	if err != nil {
		return domain.User{}, err
	}

	return domainUser, nil
}

func respondWithError(w http.ResponseWriter, r *http.Request, message string, code int) {
	respondWithJson(w, r, code, Error{
		Code:    int32(code),
		Message: message,
	})
}

func respondWithJson(w http.ResponseWriter, r *http.Request, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
