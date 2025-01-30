package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
