/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/Seiya-Tagami/pecopeco-cli/cmd"
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
