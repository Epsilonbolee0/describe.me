// Package http_v1 implements routing paths. Each services in own file.
package http_v1

import (
	"net/http"

	"describe.me/internal/controller/http_v1/handlers"
	"describe.me/internal/service"
	"describe.me/pkg/logger"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter -.
// Swagger spec:
// @title       describe.me API
// @description describe.me app server
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(
	router *chi.Mux,
	log logger.Interface,
	auth *jwtauth.JWTAuth,
	userService *service.UserService) {

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	// Prometheus
	router.Use(chiprometheus.NewMiddleware("describe.me"))

	// Swagger
	router.Mount("/swagger", httpSwagger.WrapHandler)

	// Public routes
	router.Group(func(r chi.Router) {
		handlers.NewUserHandler(router, userService, log)
	})

	// Private routes
	router.Group(func(r chi.Router) {
		router.Use(jwtauth.Verifier(auth))
		router.Use(jwtauth.Authenticator)

		router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		router.Handle("/metrics", promhttp.Handler())
	})

}
