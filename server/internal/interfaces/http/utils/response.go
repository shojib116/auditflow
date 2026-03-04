package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSON handles writing a JSON response to the wire.
func SendJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// SendError is a helper for consistent error shapes
func SendError(w http.ResponseWriter, statusCode int, message string) {
	SendJSON(w, statusCode, map[string]string{"error": message})
}
