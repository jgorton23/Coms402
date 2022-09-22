package v1

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func jsonResponse(w http.ResponseWriter, code int, msg interface{}) {
	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
