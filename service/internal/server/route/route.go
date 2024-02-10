package route

import (
	"github.com/Seiya-Tagami/pecopeco-service/internal/db"
	"github.com/Seiya-Tagami/pecopeco-service/internal/domain/user"
	"github.com/Seiya-Tagami/pecopeco-service/internal/infrastructure/repository"
	hh "github.com/Seiya-Tagami/pecopeco-service/internal/presentation/health_handler"
	uh "github.com/Seiya-Tagami/pecopeco-service/internal/presentation/user"
	mymiddleware "github.com/Seiya-Tagami/pecopeco-service/internal/server/middleware"
	uu "github.com/Seiya-Tagami/pecopeco-service/internal/usecase/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRoute(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health-check", hh.HealthCheck)
		userRoute(r)
	})
}

func userRoute(r chi.Router) chi.Router {
	db := db.GetDB()
	userRepository := repository.NewUserRepository(db)
	h := uh.NewHandler(
		uu.NewFindUserUseCase(userRepository),
		uu.NewLoginUsecase(
			user.NewUserDomainService(userRepository),
			userRepository,
		),
	)
	return r.Route("/users", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.With(mymiddleware.Auth).Get("/{id}", h.FindUser)
	})
}
