/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/Seiya-Tagami/pecopeco-cli/cmd"
	"github.com/Seiya-Tagami/pecopeco-cli/config"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
