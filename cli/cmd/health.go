package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check server status.",
	Long:  "Check server status when dev env.",
	Run: func(cmd *cobra.Command, args []string) {
		env := os.Getenv("GO_ENV")
		if env == "dev" {
			fmt.Println("health api called")
		} else {
			panic("health: command not found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
