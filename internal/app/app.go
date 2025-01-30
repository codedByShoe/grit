package app

import (
	"fmt"
	"net/http"
	"time"
)

func Run() error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	r := InitializeRouter()

	// setup Server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// run server
	fmt.Printf("🚀 Application running at %s:%d \n", cfg.ServerConfig.Url, cfg.ServerConfig.Port)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
