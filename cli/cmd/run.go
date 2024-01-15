package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	foodFactory "github.com/Seiya-Tagami/pecopeco-cli/api/factory/food"
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
	promptForMode := promptui.Select{
		Label: "What would you like to do?",
		Items: []string{"Search food", "Show favorites", "Exit"},
	}

	_, mode, err := promptForMode.Run()

	if err != nil {
		fmt.Printf("Prompt failed%v\n", err)
		return
	}

	switch mode {
	case "Search food":
		factory := foodFactory.CreateFactory()

		searchFoodInput := getSearchFoodInput()
		params := foodFactory.ListFoodParams{
			City: searchFoodInput.city,
			Food: searchFoodInput.food,
		}

		foodList, err := factory.ListFood(params)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range foodList {
			fmt.Printf("-----------------\n店名: %s\n住所: %s\n最寄り駅: %s\nジャンル: %s\nURL: %s\n", v.Name, v.Address, v.StationName, v.GenreName, v.URL)
		}

		selectOption()
	case "Show favorites":
		fmt.Printf("%s called\n", mode)
		selectOption()
	case "Exit":
		fmt.Print("Bye!\n")
		os.Exit(1)
	}
}

type searchFoodInput struct {
	city string
	food string
}

func getSearchFoodInput() searchFoodInput {
	promptForCity := promptui.Prompt{
		Label: "Which city?",
		Validate: func(input string) error {
			if utf8.RuneCountInString(input) == 0 {
				return errors.New("Please enter a city.")
			}
			if strings.TrimSpace(input) == "" {
				return errors.New("City name cannot be only whitespace.")
			}
			return nil
		},
	}
	city, err := promptForCity.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchFoodInput{}
	}

	promptForFood := promptui.Prompt{
		Label: "What food?",
		Validate: func(input string) error {
			if utf8.RuneCountInString(input) == 0 {
				return errors.New("Please enter food.")
			}
			if strings.TrimSpace(input) == "" {
				return errors.New("Food name cannot be only whitespace")
			}
			return nil
		},
	}
	food, err := promptForFood.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchFoodInput{}
	}

	return searchFoodInput{city: strings.TrimSpace(city), food: strings.TrimSpace(food)}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
