package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/petaki/inertia-go"
)

func InitializeRoutes(app *fiber.App, inertia *inertia.Inertia) {
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(csrf.New())
	app.Get("/", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inertia.Render(w, r, "Welcome", nil)
	}))
	app.Get("/metrics", monitor.New())

	app.Static("/dist", "dist")
}
