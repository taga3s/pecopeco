package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Seiya-Tagami/pecopeco-service/internal/config"
	"github.com/Seiya-Tagami/pecopeco-service/internal/server/route"
	"github.com/go-chi/chi/v5"
)

func Run(ctx context.Context) {
	router := chi.NewRouter()
	route.InitRoute(router)

	address := ":" + config.GetServerConfig().Port
	log.Printf("Starting server on localhost%s...\n", address)
	srv := &http.Server{
		Addr:              address,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}
