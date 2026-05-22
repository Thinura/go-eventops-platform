package http

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok", "service": "eventops-api"}
	WriteJSON(w, http.StatusOK, response)
}
