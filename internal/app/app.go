package app

import (
	"fmt"

	"github.com/codedbyshoe/grit/internal/model"
	"github.com/gofiber/fiber/v2"
)

func Run() error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	app := fiber.New()

	db := InitializeDatabase(cfg.Database.Url)

	db.AutoMigrate(&model.User{}, &model.Session{})

	InitializeRoutes(app)

	return app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
