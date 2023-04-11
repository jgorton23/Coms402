package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

// GetRolesBy
// Get user roles
// (GET /role)
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

// ApproveRole
// Approve the given role
// (POST /role/approve)
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

	ctx := context.WithValue(r.Context(), "userUUID", user.UUID.String())
	role, err := v1.userToCompanyUseCase.Approve(ctx, domain.UserToCompany{
		CompanyUUID: companyUUID,
		UserUUID:    userUUID,
	}, user.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to approve role mapping: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusOK, role)

}
