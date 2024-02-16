package main

import (
	"log"

	"github.com/ayanami77/pecopeco-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	cmd.Execute()
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}
