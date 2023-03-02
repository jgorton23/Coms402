package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
		logger:         do.MustInvoke[*usecase.Logger](i).WithSubsystem("controller http v1"),
		companyUseCase: do.MustInvoke[*usecase.Company](i),
		userUseCase:    do.MustInvoke[*usecase.User](i),
	}

	return httpV1, nil
}

type httpV1Implem struct {
	companyUseCase        *usecase.Company
	userUseCase           *usecase.User
	authbossAuthenticator *authboss.Authboss
	logger                *usecase.Logger
}

func (v1 httpV1Implem) AddCompany(w http.ResponseWriter, r *http.Request) {

	var requestBody AddCompanyJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondWithError(w, r, "invalid company json object found", http.StatusBadRequest)
		return
	}

	domainUser, err := v1.loadDomainUser(r)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to load user: %v", err), http.StatusBadRequest)
		return
	}

	company, err := v1.companyUseCase.Create(r.Context(), domain.Company{
		Name: requestBody.Name,
	}, domainUser.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to create company: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusCreated, company)
}

func (v1 httpV1Implem) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var requestBody UpdateCompanyJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondWithError(w, r, "invalid company json object found", http.StatusBadRequest)
		return
	}

	domainUser, err := v1.loadDomainUser(r)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to load user: %v", err), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(requestBody.Uuid)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	err = v1.companyUseCase.Update(r.Context(), domain.Company{
		UUID: id,
		Name: requestBody.Name,
	}, domainUser.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	company, err := v1.companyUseCase.GetByUUID(r.Context(), id)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusOK, company)
}

func (v1 httpV1Implem) GetCompanyByUUID(w http.ResponseWriter, r *http.Request, companyUUID string) {

	id, err := uuid.Parse(companyUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	company, err := v1.companyUseCase.GetByUUID(r.Context(), id)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusOK, company)
}

// Helper Functions
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
