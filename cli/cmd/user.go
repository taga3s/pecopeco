package cmd

import (
	"github.com/Seiya-Tagami/pecopeco-cli/api/factory/user"
	"github.com/Seiya-Tagami/pecopeco-cli/config"
	"github.com/Seiya-Tagami/pecopeco-cli/ui"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Print the logged in user",
	Long:  "Print the logged in user",
	Run: func(cmd *cobra.Command, args []string) {
		printLoggedInUser()
	},
}

func printLoggedInUser() {
	factory := user.CreateFactory()
	if !config.IsLogin() {
		ui.TextBlue().Println(errorMsg)
		return
	}
	user, err := factory.FindUser()
	if err != nil {
		ui.TextBlue().Println(errorMsg)
		return
	}
	ui.TextGreen().Printf("Logged in as %s, %s\n", user.Name, user.Email)
}

const errorMsg = "Sorry, you may have not logged in yet. Please login with following command.\n> pecopeco login\nFor more info, you can reach https://github.com/Seiya-Tagami/pecopeco"

func init() {
	rootCmd.AddCommand(userCmd)
}
