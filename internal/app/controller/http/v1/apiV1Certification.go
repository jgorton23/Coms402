package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

// GetCertificationBy
// Find certification by *
// (GET /certification)
func (v1 httpV1Implem) GetCertificationBy(w http.ResponseWriter, r *http.Request, params GetCertificationByParams) {

	domainUser, err := v1.loadDomainUser(r)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to load user: %v", err), http.StatusBadRequest)
		return
	}

	companyUUID, err := uuid.Parse(*params.CompanyUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	if params.CompanyUUID != nil && params.CertificationUUID == nil {

		itemBatches, err := v1.certificationUseCase.GetByCompanyUUID(r.Context(), companyUUID, domainUser.UUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		respondWithJson(w, r, http.StatusFound, itemBatches)
		return
	} else if params.CompanyUUID != nil && params.CertificationUUID != nil {
		certificationUUID, err := uuid.Parse(*params.CertificationUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		itemBatches, err := v1.certificationUseCase.GetByUUID(r.Context(), companyUUID, domainUser.UUID, certificationUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		respondWithJson(w, r, http.StatusFound, itemBatches)
		return
	} else {
		respondWithError(w, r, "invalid prams", http.StatusBadRequest)
		return
	}
}

// AddCertification
// Create a new item batch
// (POST /certification)
func (v1 httpV1Implem) AddCertification(w http.ResponseWriter, r *http.Request) {
	var requestBody AddCertificationJSONRequestBody

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

	cid, err := uuid.Parse(*&requestBody.CompanyUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	iid, err := uuid.Parse(*requestBody.ImageUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	ibid, err := uuid.Parse(requestBody.ItemBatchUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "userUUID", domainUser.UUID.String())
	itemBatch, err := v1.certificationUseCase.Create(ctx, domain.Certification{
		PrimaryAttribute: requestBody.PrimaryAttribute,
		ImageUUID:        iid,
		ItemBatchUUID:    ibid,
		CompanyUUID:      cid,
	}, domainUser.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to create certification: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusCreated, itemBatch)
}

// UpdateCertification
// Update an existing item batch
// (PUT /certification)
func (v1 httpV1Implem) UpdateCertification(w http.ResponseWriter, r *http.Request) {
	var requestBody UpdateCertificationJSONRequestBody

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

	cid, err := uuid.Parse(*&requestBody.CompanyUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	iid, err := uuid.Parse(*requestBody.ImageUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	ibid, err := uuid.Parse(requestBody.ItemBatchUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	tid, err := uuid.Parse(*requestBody.TemplateUUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(requestBody.Uuid)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "userUUID", domainUser.UUID.String())
	err = v1.certificationUseCase.Update(ctx, domain.Certification{
		UUID:             id,
		PrimaryAttribute: requestBody.PrimaryAttribute,
		ImageUUID:        iid,
		ItemBatchUUID:    ibid,
		TemplateUUID:     tid,
		CompanyUUID:      cid,
	}, domainUser.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to update certification: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusAccepted, requestBody)
}
