package server

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/codedbyshoe/grit/internal/config"
	"github.com/codedbyshoe/grit/internal/database"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/petaki/inertia-go"
)

type Server struct {
	cfg            *config.Config
	db             database.Service
	session        *sessions.CookieStore
	inertiaManager *inertia.Inertia
}

func NewServer(assets embed.FS) *http.Server {

	config := config.Load()
	database := database.New(config)
	session := NewStore(config.SecretKey)
	CSRF := csrf.Protect([]byte(config.SecretKey), csrf.Path("/"))

	inertia := inertia.NewWithFS(config.AppUrl, config.ViewsUrl, "", assets)

	NewServer := &Server{
		cfg:            config,
		db:             database,
		session:        session,
		inertiaManager: inertia,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", NewServer.cfg.Port),
		Handler:      CSRF(NewServer.RegisterRoutes(assets)),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
