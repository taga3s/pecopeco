package cmd

import (
	"time"

	"github.com/Seiya-Tagami/pecopeco-cli/api/factory/user"
	"github.com/Seiya-Tagami/pecopeco-cli/config"
	uiutil "github.com/Seiya-Tagami/pecopeco-cli/ui/util"
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
	if !config.IsLogin() {
		uiutil.TextBlue().Println(errorMsg)
		time.Sleep(time.Second * 1)
		return
	}

	factory := user.CreateFactory()
	user, err := factory.FindUser()
	if err != nil {
		uiutil.TextBlue().Println(errorMsg)
		time.Sleep(time.Second * 1)
		return
	}
	uiutil.TextGreen().Printf("Logged in as %s, %s\n", user.Name, user.Email)
}

const errorMsg = "Sorry, you may have not logged in yet. Please login with following command.\n\n```\n> pecopeco login\n```"

func init() {
	rootCmd.AddCommand(userCmd)
}
