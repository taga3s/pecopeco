package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run pecopeco CLI",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		selectOption()
	},
}

func selectOption() {
	prompt := promptui.Select{
		Label: "What would you like to do?",
		Items: []string{"Search food", "Show favorites", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed%v\n", err)
		return
	}

	switch result {
	case "Search food":
		fmt.Printf("%s called\n", result)
	case "Show favorites":
		fmt.Printf("%s called\n", result)
	case "Exit":
		fmt.Print("Bye!\n")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
