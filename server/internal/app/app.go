// Package app configures and runs application.
package app

import (
	jwt_auth "describe.me/pkg/jwt-auth"
	"fmt"
	"github.com/go-chi/chi/v5"
	"os"
	"os/signal"
	"syscall"

	"describe.me/config"
	"describe.me/internal/controller/http_v1"
	"describe.me/internal/service"
	"describe.me/internal/service/repository"
	"describe.me/pkg/httpserver"
	"describe.me/pkg/logger"
	"describe.me/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	log := logger.New(cfg.Logger.Level)

	// Repository
	pg, err := postgres.New(
		cfg.DSN,
		postgres.MaxIdleConnections(cfg.Postgres.MaxIdleConnections),
		postgres.MaxOpenConnections(cfg.Postgres.MaxOpenConnections),
	)

	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Service
	userService := service.New(
		repository.NewUserRepository(pg),
	)

	// Auth
	jwtAuth := jwt_auth.NewJWTAuth(cfg.Shelter.Secret)

	// HTTP Server
	handler := chi.NewRouter()
	http_v1.NewRouter(
		handler,
		log,
		jwtAuth,
		userService,
	)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
