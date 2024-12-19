package server

import "github.com/go-chi/chi"

type Route interface {
	Mount(r chi.Router)
	Pattern() string
}
