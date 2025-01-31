package app

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"github.com/petaki/inertia-go"
	"gorm.io/gorm"
)

type Application struct {
	Config  *Config
	DB      *gorm.DB
	Inertia *inertia.Inertia
	Ctx     context.Context
}

func NewApplication() *Application {
	config, err := LoadConfig()
	if err != nil {
		panic("could not load config")
	}
	db := InitializeDatabase(config.Database.Name)
	inertia := inertia.New(config.Server.Url, config.Server.Template, "")

	return &Application{
		Config:  config,
		DB:      db,
		Inertia: inertia,
		Ctx:     context.Background(),
	}
}

func (a *Application) Run() error {
	a.Inertia.ShareFunc("assets", func() template.HTML {
		return template.HTML(`
        <link rel="stylesheet" href="/dist/css/style.css"></link>
        <script src="/dist/main.js" defer></script>
        <script src="/dist/manifest.js" defer></script>
        <script src="/dist/vendor.js" defer></script>
      `)
	})

	CSRF := csrf.Protect([]byte(a.Config.Session.ApplicationKey))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", a.Config.Server.Port),
		Handler:      CSRF(a.InitializeRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("🚀 Application running at http://localhost:%s", server.Addr)
	return server.ListenAndServe()
}
