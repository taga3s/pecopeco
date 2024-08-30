package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/taga3s/pecopeco-cli/api/factory/health"
	uiutil "github.com/taga3s/pecopeco-cli/ui/util"
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
			uiutil.TextGreen().Printf("service health check... status: %d, message: %s\n", health.Status, health.Message)
		} else {
			panic("health: command not found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
