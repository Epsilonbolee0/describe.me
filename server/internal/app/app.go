// Package app configures and runs application.
package app

import (
	"fmt"
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

	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	// Repository
	pg, err := postgres.New(
		cfg.DSN,
		postgres.MaxIdleConnections(cfg.Postgres.MaxIdleConnections),
		postgres.MaxOpenConnections(cfg.Postgres.MaxOpenConnections),
	)

	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Service
	userService := service.New(
		repository.NewUserRepository(pg),
	)

	// HTTP Server
	handler := gin.New()
	http_v1.NewRouter(handler, l, userService)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
