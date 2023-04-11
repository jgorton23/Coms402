package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

// Find certificationPDF by *
// (GET /certification/pdf)
func (v1 httpV1Implem) GetCertificationPDFBy(w http.ResponseWriter, r *http.Request, params GetCertificationPDFByParams) {
	if params.CertificationPDFUUID != nil {

		uuid, err := uuid.Parse(*params.CertificationPDFUUID)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		reader, err := v1.certificationPDFUseCase.GetByUUID(r.Context(), uuid)

		if err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-type", "application/pdf")

		if _, err := io.Copy(w, reader); err != nil {
			respondWithError(w, r, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
		return
	}

	respondWithError(w, r, "unimplemented", http.StatusNotImplemented)
}

// Create a new certification pdf
// (POST /certification/pdf)
func (v1 httpV1Implem) AddCertificationPDF(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(1 << 20); err != nil {
		respondWithError(w, r, "Max pdf size if 1MB", http.StatusRequestEntityTooLarge)
		return
	}
	file, _, err := r.FormFile("upload.pdf")

	if err != nil {
		respondWithError(w, r, "file needs to be called upload.pdf", http.StatusBadRequest)
		return
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		respondWithError(w, r, "upload server error", http.StatusInternalServerError)
	}

	domainUser, err := v1.loadDomainUser(r)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to load user: %v", err), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "userUUID", domainUser.UUID.String())
	c, err := v1.certificationPDFUseCase.Create(ctx, buf)

	if err != nil {
		respondWithError(w, r, fmt.Sprintf("unable to create certification: %v", err), http.StatusBadRequest)
		return
	}

	respondWithJson(w, r, http.StatusCreated, struct {
		Status string
		UUID   string
	}{Status: "ok", UUID: c.String()})
}
