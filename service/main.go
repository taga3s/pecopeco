package main

import (
	"context"

	"github.com/Seiya-Tagami/pecopeco-service/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.Run(ctx)
}
