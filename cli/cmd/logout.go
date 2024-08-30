package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taga3s/pecopeco-cli/config"
	uiutil "github.com/taga3s/pecopeco-cli/ui/util"
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
