package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if arr := json.NewEncoder(w).Encode(payload); arr != nil {
		slog.Error("Failed to write JSON response", "error", arr)
	}
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, ErrorResponse{Error: message})
}
