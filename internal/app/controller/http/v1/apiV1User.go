package http

import (
	"fmt"
	"net/http"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"github.com/google/uuid"
)

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
