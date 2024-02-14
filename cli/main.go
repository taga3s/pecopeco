package main

import (
	"log"
	"os"

	"github.com/Seiya-Tagami/pecopeco-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	cmd.Execute()
}

func loadEnv() {
	if os.Getenv("GO_ENV") == "dev" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}
}
