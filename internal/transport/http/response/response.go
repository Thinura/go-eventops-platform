package response

import (
	"encoding/json"
	"net/http"

	"github.com/Thinura/go-eventops-platform/internal/infrastructure/logger"
)

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error("failed to write json response", "error", err)
	}
}

func Error(w http.ResponseWriter, statusCode int, code string, message string) {
	JSON(w, statusCode, ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
		},
	})
}
