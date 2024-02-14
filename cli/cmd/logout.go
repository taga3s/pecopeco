package cmd

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/config"
	uiutil "github.com/Seiya-Tagami/pecopeco-cli/ui/util"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the pecopeco-cli",
	Long:  "Logout from the pecopeco-cli",
	Run: func(cmd *cobra.Command, args []string) {
		logout()
	},
}

func logout() {
	if err := config.Revoke(config.PECOPECO_API_TOKEN); err != nil {
		fmt.Println(err)
	}
	uiutil.TextGreen().Println("Successfully logout!")
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
