/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Seiya-Tagami/pecopeco-cli/cmd"
	"github.com/Seiya-Tagami/pecopeco-cli/config"
)

func main() {
	config.Load()
	cmd.Execute()
}
