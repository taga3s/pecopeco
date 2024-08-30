package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/taga3s/pecopeco-cli/cmd"
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
