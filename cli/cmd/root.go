package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/taga3s/pecopeco-cli/config"
	"github.com/taga3s/pecopeco-cli/ui/util"
)

var rootCmd = &cobra.Command{
	Use:   "pecopeco-cli",
	Short: "A restaurant searching application",
	Long:  "The pecopeco-cli enables you to search restaurants and register it as a favorite.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Runの前に必ず実行される
func init() {
	util.Clear()
	cobra.OnInitialize(config.Load)
}
