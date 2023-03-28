package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

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

	domainUser, err := v1.loadDomainUser(r)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to load user: %v", err), http.StatusBadRequest)
		return
	}

	assignerUUID = domainUser.UUID

	if requestBody.UserUUID == nil {
		userUUID = domainUser.UUID
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
		respondWithError(w, r, fmt.Sprintf("unable to create company: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusCreated, role)

}
