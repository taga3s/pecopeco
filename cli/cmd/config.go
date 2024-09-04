package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/taga3s/pecopeco-cli/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure application settings",
	Long:  "Configure application settings",
	Run: func(cmd *cobra.Command, args []string) {
		// for setting LINE Notify API Token
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println(err)
			return
		}
		if token != "" {
			viper.Set(config.LINE_NOTIFY_API_TOKEN, token)
			if err := viper.WriteConfig(); err != nil {
				fmt.Println("Error writing config file:", err)
				return
			}
			fmt.Println("Updated LINE Notify API Token:", viper.GetString(config.LINE_NOTIFY_API_TOKEN))
		}

		// for setting Username
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			fmt.Println(err)
			return
		}
		if username != "" {
			viper.Set(config.USERNAME, username)
			if err := viper.WriteConfig(); err != nil {
				fmt.Println("Error writing config file:", err)
				return
			}
			fmt.Println("Updated Username:", viper.GetString(config.USERNAME))
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("token", "t", "", "LINE Notify API Token")
	configCmd.Flags().StringP("username", "u", "", "Username")
}
