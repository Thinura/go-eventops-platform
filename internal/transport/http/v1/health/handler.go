package health

import (
	response "github.com/Thinura/go-eventops-platform/internal/transport/http/response"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "eventops-api",
		"version": "v1",
	})

}
