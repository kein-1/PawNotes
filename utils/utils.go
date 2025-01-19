package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, content any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(content)
}

func WriteError(w http.ResponseWriter, status int, content any) {
	WriteJSON(w, status, content)
}
