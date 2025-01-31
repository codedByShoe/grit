package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/petaki/inertia-go"
)

type WelcomeHandler struct {
	*chi.Mux
	inertia *inertia.Inertia
}

func NewWelcomeHandler(inertia *inertia.Inertia) *WelcomeHandler {
	h := WelcomeHandler{
		Mux:     chi.NewMux(),
		inertia: inertia,
	}

	h.Route("/", func(r chi.Router) {
		r.Get("/", h.Welcome)
	})

	return &h
}

func (h *WelcomeHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	_ = h.inertia.Render(w, r, "Welcome", nil)
}
