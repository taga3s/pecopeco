package main

import (
	"context"

	"github.com/Seiya-Tagami/pecopeco-service/internal/db"
	"github.com/Seiya-Tagami/pecopeco-service/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db.NewMySQL()
	defer db.CloseDB()

	server.Run(ctx)
}
