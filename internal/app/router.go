package app

import (
	"net/http"

	"github.com/codedbyshoe/grit/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *Application) InitializeRoutes() *chi.Mux {
	r := chi.NewRouter()

	wh := handler.NewWelcomeHandler(a.Inertia)

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
	r.Mount("/", wh)

	fileServer := http.FileServer(http.Dir("./dist"))
	r.Handle("/dist/*", http.StripPrefix("/dist/", fileServer))

	return r
}
