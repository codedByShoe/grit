package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/codedbyshoe/grit/internal/config"
	"github.com/codedbyshoe/grit/internal/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schema
	db.AutoMigrate(&model.User{})

	return db, nil
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func NewRouter(routes []Route) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Routes
	for _, route := range routes {
		r.Group(func(r chi.Router) {
			r.Route(route.Pattern(), func(r chi.Router) {
				route.Mount(r)
			})
		})
	}

	return r
}

func NewHTTPServer(cfg *config.Config, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerPort),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func RegisterHTTPServer(lifecycle fx.Lifecycle, srv *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return fmt.Errorf("failed to listen: %w", err)
			}
			fmt.Printf("🚀 Server running at http://127.0.0.1:%s", srv.Addr)

			if err := srv.Serve(ln); err != http.ErrServerClosed {
				fmt.Printf("Http server error %v\n", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
