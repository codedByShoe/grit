package app

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)


func InitializeRouter() *chi.Mux {
	r := chi.NewRouter()
	// global middleware on all requests
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r
}