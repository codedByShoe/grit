package app

import (
	"fmt"
	"html/template"

	"github.com/codedbyshoe/grit/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/petaki/inertia-go"
)

func Run() error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	app := fiber.New()

	db := InitializeDatabase(cfg.Database.Url)

	db.AutoMigrate(&model.User{}, &model.Session{})

	inertiaManager := inertia.New(cfg.Server.Url, cfg.Server.Template, "")

	inertiaManager.ShareFunc("assets", func() template.HTML {
		return template.HTML(`
        <link rel="stylesheet" href="/dist/css/style.css"></link>
        <script src="/dist/main.js" defer></script>
        <script src="/dist/manifest.js" defer></script>
        <script src="/dist/vendor.js" defer></script>
      `)
	})

	InitializeRoutes(app, inertiaManager)

	return app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
