package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	foodFactory "github.com/Seiya-Tagami/pecopeco-cli/api/factory/food"
	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/Seiya-Tagami/pecopeco-cli/ui"
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
		if len(foodList) == 0 {
			ui.TextBlue().Println("Sorry, there is no data. Please try to change the input parameters.")
			time.Sleep(1 * time.Second)
			run()
		}

		selectFoodResult, err := selectFood(foodList)
		if err != nil {
			fmt.Println(err)
			return
		}

		if selectFoodResult.addToFavorites {
			ui.TextGreen().Println("Add to favorites!")
		}
		if selectFoodResult.notify {
			ui.TextGreen().Println("Notify to your line app!")
		}
		time.Sleep(2 * time.Second)
		run()
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
		Label: "> Which city?",
		Validate: func(input string) error {
			if utf8.RuneCountInString(input) == 0 {
				return errors.New("Please enter a city.")
			}
			if strings.TrimSpace(input) == "" {
				return errors.New("City name cannot be only whitespace.")
			}
			if strings.Contains(input, " ") {
				return errors.New("City name cannot contain whitespace.")
			}
			return nil
		},
		Templates: ui.DefaultPromptTemplate(),
	}
	city, err := promptForCity.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchFoodInput{}
	}

	promptForFood := promptui.Prompt{
		Label: "> What food?",
		Validate: func(input string) error {
			if utf8.RuneCountInString(input) == 0 {
				return errors.New("Please enter food.")
			}
			if strings.TrimSpace(input) == "" {
				return errors.New("Food name cannot be only whitespace.")
			}
			if strings.Contains(input, " ") {
				return errors.New("Food name cannot contain whitespace.")
			}
			return nil
		},
		Templates: ui.DefaultPromptTemplate(),
	}
	food, err := promptForFood.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchFoodInput{}
	}

	return searchFoodInput{city: strings.TrimSpace(city), food: strings.TrimSpace(food)}
}

type selectFoodResult struct {
	food           model.Food
	addToFavorites bool
	notify         bool
}

func selectFood(foodList []model.Food) (selectFoodResult, error) {
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
		return selectFoodResult{}, err
	}

	food := foodMap[option]
	ui.TextGreen().Printf("---------------------\n[店名] %s\n[住所] %s\n[最寄り駅] %s\n[ジャンル] %s\n[URL] %s\n---------------------\n",
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
		return selectFoodResult{}, err
	}

	if decision == "Yes" {
		result := selectFoodResult{}
		result.food = food

		promptForAddToFavorites := promptui.Select{
			Label: "Add to favorites?",
			Items: []string{"Yes", "No"},
		}
		_, addToFavorites, err := promptForAddToFavorites.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return selectFoodResult{}, err
		}
		if addToFavorites == "Yes" {
			result.addToFavorites = true
		}

		promptForNotify := promptui.Select{
			Label: "Notify your LINE app?",
			Items: []string{"Yes", "No"},
		}
		_, notify, err := promptForNotify.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return selectFoodResult{}, err
		}
		if notify == "Yes" {
			result.notify = true
		}

		return result, nil
	} else {
		return selectFood(foodList)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
