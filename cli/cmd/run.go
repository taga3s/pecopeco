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
		Items: []string{"search food", "show favorites", "exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed%v\n", err)
		return
	}

	switch result {
	case "search food":
		fmt.Printf("%s called\n", result)
	case "show your favorites":
		fmt.Printf("%s called\n", result)
	case "exit":
		fmt.Print("bye!\n")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
