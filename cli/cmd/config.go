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
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println(err)
			return
		}
		viper.Set(config.LINE_NOTIFY_API_TOKEN, token)
		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("Error writing config file: %s\n", err)
			return
		}
		fmt.Println("Updated line_notify_api_token:", viper.GetString(config.LINE_NOTIFY_API_TOKEN))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("token", "t", "", "LINE notify token")
}
