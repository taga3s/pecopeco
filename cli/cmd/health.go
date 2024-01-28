package cmd

import (
	"fmt"
	"os"

	"github.com/Seiya-Tagami/pecopeco-cli/api/factory/health"
	"github.com/Seiya-Tagami/pecopeco-cli/ui"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check server status.",
	Long:  "Check server status when dev env.",
	Run: func(cmd *cobra.Command, args []string) {
		env := os.Getenv("GO_ENV")
		if env == "dev" {
			factory := health.CreateFactory()
			health, err := factory.HealthCheck()
			if err != nil {
				fmt.Println(err)
				return
			}
			ui.TextGreen().Printf("service health check... status: %d, message: %s\n", health.Status, health.Message)
		} else {
			panic("health: command not found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
