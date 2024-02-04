package cmd

import (
	"fmt"
	"os"

	"github.com/Seiya-Tagami/pecopeco-cli/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pecopeco-cli",
	Short: "A restaurant searching application",
	Long:  "The pecopeco-cli enables you to search restaurants and register it as a favorite.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Runの前に必ず実行される
	cobra.OnInitialize(config.Load)
}
