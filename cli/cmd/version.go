package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of pecopeco-cli",
	Long:  "All software has versions. This is pecopeco-cli's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
