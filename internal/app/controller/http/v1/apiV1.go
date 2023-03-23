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
		logger:               do.MustInvoke[*usecase.Logger](i).WithSubsystem("controller http v1"),
		companyUseCase:       do.MustInvoke[*usecase.Company](i),
		userUseCase:          do.MustInvoke[*usecase.User](i),
		userToCompanyUseCase: do.MustInvoke[*usecase.UserToCompany](i),
	}

	return httpV1, nil
}

type httpV1Implem struct {
	companyUseCase        *usecase.Company
	userUseCase           *usecase.User
	authbossAuthenticator *authboss.Authboss
	logger                *usecase.Logger
	userToCompanyUseCase  *usecase.UserToCompany
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

func (v1 httpV1Implem) AddUserRole(w http.ResponseWriter, r *http.Request) {

	var requestBody AddUserRoleJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondWithError(w, r, "invalid company json object found", http.StatusBadRequest)
		return
	}

	var assignerUUID uuid.UUID
	var userUUID uuid.UUID

	if requestBody.UserUUID == nil {

		domainUser, err := v1.loadDomainUser(r)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("unable to load user: %v", err), http.StatusBadRequest)
			return
		}

		userUUID = domainUser.UUID
		assignerUUID = domainUser.UUID
	} else {
		userUUID, err = uuid.Parse(*requestBody.UserUUID)
		if err != nil {
			respondWithError(w, r, fmt.Sprintf("invalid user UUID: %v", err), http.StatusBadRequest)
			return
		}
	}

	companyUUID, err := uuid.Parse(requestBody.CompanyUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("invalid user UUID: %v", err), http.StatusBadRequest)
		return
	}

	role, err := v1.userToCompanyUseCase.Add(r.Context(), domain.UserToCompany{
		CompanyUUID: companyUUID,
		UserUUID:    userUUID,
		RoleType:    domain.Role(requestBody.RoleType),
		Approved:    *requestBody.Approved,
	}, assignerUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to create role mapping: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusCreated, role)

}

func (v1 httpV1Implem) ApproveRole(w http.ResponseWriter, r *http.Request, params ApproveRoleParams) {

	var user domain.User

	userUUID, err := uuid.Parse(params.UserUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing UUID: %v", err), http.StatusBadRequest)
		return
	}

	companyUUID, err := uuid.Parse(params.CompanyUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing UUID: %v", err), http.StatusBadRequest)
		return
	}

	user, err = v1.loadDomainUser(r)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	role, err := v1.userToCompanyUseCase.Approve(r.Context(), domain.UserToCompany{
		CompanyUUID: companyUUID,
		UserUUID:    userUUID,
	}, user.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to approve role mapping: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusOK, role)

}

func (v1 httpV1Implem) GetUserBy(w http.ResponseWriter, r *http.Request, params GetUserByParams) {

	var user domain.User

	if params.UserUUID != nil && params.UserEmail != nil {
		respondWithError(w, r, "only userUUID or userEmail can be defined", http.StatusBadRequest)
		return
	}

	if params.UserUUID != nil {
		id, err := uuid.Parse(*params.UserUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
			return
		}

		user, err = v1.userUseCase.FindByUUID(r.Context(), id)
		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
	} else if params.UserEmail != nil {
		var err error
		user, err = v1.userUseCase.FindByEmail(r.Context(), *params.UserEmail)
		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
	}

	respondWithJson(w, r, http.StatusOK, User{
		Created: user.CreatedAt.String(),
		Email:   user.Email,
		Uuid:    user.UUID.String(),
	})
}

func (v1 httpV1Implem) GetRolesBy(w http.ResponseWriter, r *http.Request, params GetRolesByParams) {

	var roles []domain.UserToCompany

	if params.CompanyUUID == nil && params.UserUUID == nil {
		respondWithError(w, r, "Must supply companyUUID or userUUID", http.StatusBadRequest)
		return
	} else if params.CompanyUUID != nil && params.UserUUID != nil {
		userUUID, err := uuid.Parse(*params.UserUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		companyUUID, err := uuid.Parse(*params.CompanyUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		role, err := v1.userToCompanyUseCase.FindByUUIDS(r.Context(), companyUUID, userUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		roles = append(roles, role)

	} else if params.UserUUID != nil {
		id, err := uuid.Parse(*params.UserUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		roles, err = v1.userToCompanyUseCase.FindByUserUUID(r.Context(), id)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
	} else if params.CompanyUUID != nil {
		id, err := uuid.Parse(*params.CompanyUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		roles, err = v1.userToCompanyUseCase.FindByCompanyUUID(r.Context(), id)
		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
	}

	respondWithJson(w, r, http.StatusOK, roles)
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
