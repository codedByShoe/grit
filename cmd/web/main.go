package main

import (
	"context"
	"log"
	"time"

	"github.com/codedbyshoe/grit/internal/config"
	"github.com/codedbyshoe/grit/internal/handler"
	"github.com/codedbyshoe/grit/internal/repository"
	"github.com/codedbyshoe/grit/internal/server"
	"github.com/codedbyshoe/grit/internal/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.StartTimeout(20*time.Second),
		fx.Provide(
			config.NewConfig,
			server.NewDatabase,
			fx.Annotate(
				server.NewRouter,
				fx.ParamTags(`group:"routes"`),
			),
			server.NewHTTPServer,
			repository.NewUserRepository,
			service.NewUserService,
			server.AsRoute(handler.NewUserHandler),
			server.AsRoute(handler.NewHealthHandler),
		),
		fx.Invoke(server.RegisterHTTPServer),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
