package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"github.com/google/uuid"
)

// Find itembatch by *
// (GET /itembatch)
func (v1 httpV1Implem) GetItemBatchBy(w http.ResponseWriter, r *http.Request, params GetItemBatchByParams) {

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

	if params.CompanyUUID != nil && params.ItemBatchUUID == nil {

		itemBatches, err := v1.itemBatchUseCase.GetByCompanyUUID(r.Context(), companyUUID, domainUser.UUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		respondWithJson(w, r, http.StatusCreated, itemBatches)
		return
	} else if params.CompanyUUID != nil && params.ItemBatchUUID != nil {
		itemBatchUUID, err := uuid.Parse(*params.ItemBatchUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		itemBatches, err := v1.itemBatchUseCase.GetByUUID(r.Context(), companyUUID, domainUser.UUID, itemBatchUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		respondWithJson(w, r, http.StatusCreated, itemBatches)
		return
	} else {
		respondWithError(w, r, "invalid prams", http.StatusBadRequest)
		return
	}
}

// Create a new item batch
// (POST /itembatch)
func (v1 httpV1Implem) AddItemBatch(w http.ResponseWriter, r *http.Request) {
	var requestBody AddItemBatchJSONRequestBody

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

	id, err := uuid.Parse(requestBody.CompanyUuid)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	itemBatch, err := v1.itemBatchUseCase.Create(r.Context(), domain.ItemBatch{
		ItemNumber:  requestBody.ItemNumber,
		CompanyUUID: id,
		Description: *requestBody.Description,
	}, domainUser.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to create itembatch: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusCreated, itemBatch)
}

// Update an existing item batch
// (PUT /itembatch)
func (v1 httpV1Implem) UpdateItemBatch(w http.ResponseWriter, r *http.Request) {
	var requestBody UpdateItemBatchJSONRequestBody

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

	cid, err := uuid.Parse(requestBody.CompanyUuid)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(requestBody.Uuid)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("error parsing uuid: %v", err), http.StatusBadRequest)
		return
	}

	err = v1.itemBatchUseCase.Update(r.Context(), domain.ItemBatch{
		UUID:        id,
		ItemNumber:  requestBody.ItemNumber,
		CompanyUUID: cid,
		Description: *requestBody.Description,
	}, domainUser.UUID)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to update itembatch: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusAccepted, requestBody)
}