package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Pattern() string {
	return "/health"

}

func (h *HealthHandler) Mount(r chi.Router) {
	r.Get("/", h.Check)
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "healthy"}); err != nil {
		w.Write([]byte("Error parsing json"))
	}
}
