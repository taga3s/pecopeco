package route

import (
	healthhandler "github.com/Seiya-Tagami/pecopeco-service/internal/presentation/health_handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRoute(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health-check", healthhandler.HealthCheck)
	})
}
