package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponseWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SendErrorWithError(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	http.Error(w, msg, statusCode)
}
