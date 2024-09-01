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
			fmt.Println("Error writing config file:", err)
			return
		}
		fmt.Println("Updated LINE Notify API Token:", viper.GetString(config.LINE_NOTIFY_API_TOKEN))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("token", "t", "", "LINE Notify API Token")
}
