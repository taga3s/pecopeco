package main

import (
	"context"
	"fmt"

	"github.com/ayanami77/pecopeco-service/internal/db"
	"github.com/ayanami77/pecopeco-service/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := db.NewMySQL()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.CloseDB()

	server.Run(ctx)
}
