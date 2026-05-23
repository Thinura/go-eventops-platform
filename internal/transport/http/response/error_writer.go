package response

import (
	"net/http"

	"github.com/Thinura/go-eventops-platform/internal/infrastructure/logger"
	"github.com/Thinura/go-eventops-platform/internal/transport/http/httperror"
)

func AppError(w http.ResponseWriter, err error) {
	mappedErr := httperror.Map(err)

	if mappedErr.StatusCode >= http.StatusInternalServerError {
		logger.Error("request failed", "error", err)
	}

	Error(w, mappedErr.StatusCode, mappedErr.Code, mappedErr.Message)
}
