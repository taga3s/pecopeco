package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		viper.Set("line_notify_api_token", token)
		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("Error writing config file: %s\n", err)
			return
		}
		fmt.Println("Updated line_notify_api_token:", viper.GetString("line_notify_api_token"))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("token", "t", "", "LINE notify token")
}
