package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	foodFactory "github.com/Seiya-Tagami/pecopeco-cli/api/factory/food"
	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run pecopeco CLI",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
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

		food, label, err := selectFood(foodList)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch label {
		case "LINE":
			fmt.Println(food)
		case "favorites":
			fmt.Println(food)
		case "cancel":
			run()
		}

	case "Show favorites":
		fmt.Printf("%s called\n", mode)
		run()
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

func selectFood(foodList []model.Food) (model.Food, string, error) {
	foodMap := map[string]model.Food{}
	options := make([]string, 0, len(foodList))

	for _, v := range foodList {
		foodMap[v.Name] = v
		options = append(options, v.Name)
	}

	promptForOptions := promptui.Select{
		Label: "Please select following food",
		Items: options,
	}

	_, option, err := promptForOptions.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return model.Food{}, "", err
	}

	food := foodMap[option]

	c := color.New(color.FgHiGreen)
	c.Printf("---------------------\n[店名] %s\n[住所] %s\n[最寄り駅] %s\n[ジャンル] %s\n[URL] %s\n---------------------\n",
		food.Name,
		food.Address,
		food.StationName,
		food.GenreName,
		food.URL,
	)

	promptForDecision := promptui.Select{
		Label: "Select this food?",
		Items: []string{"Yes", "No"},
	}

	_, decision, err := promptForDecision.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return model.Food{}, "", err
	}

	if decision == "Yes" {
		promptForNextAction := promptui.Select{
			Label: "What do you do?",
			Items: []string{"Notify your LINE app", "Add to favorites", "Cancel"},
		}

		_, nextAction, err := promptForNextAction.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return model.Food{}, "", err
		}
		if nextAction == "Notify your LINE app" {
			return food, "LINE", nil
		} else if nextAction == "Add to favorites" {
			return food, "favorites", nil
		} else {
			return food, "cancel", nil
		}
	} else {
		return selectFood(foodList)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
