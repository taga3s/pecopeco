package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/taga3s/pecopeco-service/internal/db"
	"github.com/taga3s/pecopeco-service/internal/infrastructure/repository"
	hh "github.com/taga3s/pecopeco-service/internal/presentation/health_handler"
	rh "github.com/taga3s/pecopeco-service/internal/presentation/restaurant"
	sh "github.com/taga3s/pecopeco-service/internal/presentation/search"
	ru "github.com/taga3s/pecopeco-service/internal/usecase/restaurant"
)

func InitRoute(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health-check", hh.HealthCheck)
		searchRoute(r)
		restaurantRoute(r)
	})
}

func searchRoute(r chi.Router) chi.Router {
	h := sh.NewHandler()
	return r.Route("/search", func(r chi.Router) {
		r.Get("/genres", h.ListGenres)
		r.Get("/restaurants", h.ListRestaurantsByCityAndGenre)
	})
}

func restaurantRoute(r chi.Router) chi.Router {
	db := db.GetDB()
	restaurantRepository := repository.NewRestaurantRepository(db)
	h := rh.NewHandler(
		ru.NewListRestaurantsUseCase(restaurantRepository),
		ru.NewSaveRestaurantUseCase(restaurantRepository),
		ru.NewDeleteRestaurantsUseCase(restaurantRepository),
	)
	return r.Route("/restaurants", func(r chi.Router) {
		r.Get("/", h.ListRestaurants)
		r.Post("/", h.SaveRestaurant)
	})
}
