package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	restaurantFactory "github.com/Seiya-Tagami/pecopeco-cli/api/factory/restaurant"
	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/Seiya-Tagami/pecopeco-cli/config"
	"github.com/Seiya-Tagami/pecopeco-cli/ui"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Items: []string{"Search restaurants", "Show favorites", "Exit"},
	}

	_, mode, err := promptForMode.Run()
	if err != nil {
		fmt.Printf("Prompt failed%v\n", err)
		return
	}
	switch mode {
	case "Search restaurants":
		factory := restaurantFactory.CreateFactory()

		searchRestaurantInput := getSearchRestaurantInput()
		params := restaurantFactory.ListRestaurantsParams{
			City: searchRestaurantInput.city,
			Food: searchRestaurantInput.food,
		}
		restaurantList, err := factory.ListRestaurants(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(restaurantList) == 0 {
			ui.TextBlue().Println("Sorry, there is no data. Please try to change the input parameters.")
			time.Sleep(1 * time.Second)
			run()
		}

		selectRestaurantResult, err := selectRestaurant(restaurantList)
		if err != nil {
			fmt.Println(err)
			return
		}

		if selectRestaurantResult.addToFavorites {
			ui.TextGreen().Println("Add to favorites!")
		}
		if selectRestaurantResult.notify {
			params := restaurantFactory.NotifyRestaurantToLINEParams{
				Restaurant: selectRestaurantResult.restaurant,
			}
			err := factory.NotifyRestaurantToLINE(params)
			if err != nil {
				fmt.Println(err)
			} else {
				ui.TextGreen().Println("Notify to your line app!")
			}
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

type searchRestaurantInput struct {
	city string
	food string
}

func getSearchRestaurantInput() searchRestaurantInput {
	promptForCity := promptui.Prompt{
		Label: "> Which city?",
		Validate: func(input string) error {
			if utf8.RuneCountInString(input) == 0 {
				return errors.New("Please enter a city.")
			}
			if strings.TrimSpace(input) == "" {
				return errors.New("City cannot be only whitespace.")
			}
			if strings.Contains(input, " ") {
				return errors.New("City cannot contain whitespace.")
			}
			return nil
		},
		Templates: ui.DefaultPromptTemplate(),
	}
	city, err := promptForCity.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchRestaurantInput{}
	}

	promptForRestaurant := promptui.Prompt{
		Label: "> What restaurant?",
		Validate: func(input string) error {
			if utf8.RuneCountInString(input) == 0 {
				return errors.New("Please enter restaurant.")
			}
			if strings.TrimSpace(input) == "" {
				return errors.New("Restaurant cannot be only whitespace.")
			}
			if strings.Contains(input, " ") {
				return errors.New("Restaurant cannot contain whitespace.")
			}
			return nil
		},
		Templates: ui.DefaultPromptTemplate(),
	}
	restaurant, err := promptForRestaurant.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return searchRestaurantInput{}
	}

	return searchRestaurantInput{city: strings.TrimSpace(city), food: strings.TrimSpace(restaurant)}
}

type selectRestaurantResult struct {
	restaurant     model.Restaurant
	addToFavorites bool
	notify         bool
}

func selectRestaurant(restaurantList []model.Restaurant) (selectRestaurantResult, error) {
	restaurantMap := map[string]model.Restaurant{}
	options := make([]string, 0, len(restaurantList))

	for _, v := range restaurantList {
		restaurantMap[v.Name] = v
		options = append(options, v.Name)
	}

	promptForOptions := promptui.Select{
		Label: "Please select following restaurants",
		Items: options,
	}

	_, option, err := promptForOptions.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return selectRestaurantResult{}, err
	}

	restaurant := restaurantMap[option]
	ui.TextGreen().Printf("---------------------\n[店名] %s\n[住所] %s\n[最寄り駅] %s\n[ジャンル] %s\n[URL] %s\n---------------------\n",
		restaurant.Name,
		restaurant.Address,
		restaurant.StationName,
		restaurant.GenreName,
		restaurant.URL,
	)

	promptForDecision := promptui.Select{
		Label: "Select this restaurant?",
		Items: []string{"Yes", "No"},
	}

	_, decision, err := promptForDecision.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return selectRestaurantResult{}, err
	}

	if decision == "Yes" {
		result := selectRestaurantResult{}
		result.restaurant = restaurant

		promptForAddToFavorites := promptui.Select{
			Label: "Add to favorites?",
			Items: []string{"Yes", "No"},
		}
		_, addToFavorites, err := promptForAddToFavorites.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return selectRestaurantResult{}, err
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
			return selectRestaurantResult{}, err
		}
		if notify == "Yes" {
			// トークンがセットされていない場合、ここで弾くようにする。
			if viper.GetString(config.LINE_NOTIFY_API_TOKEN) == "" {
				ui.TextBlue().Println("Sorry, you have not set your personal token to notify your line app yet. To notify your line app, you can use following command.\n> pecopeco config --token <your personal token>\nFor more info, you can reach https://github.com/Seiya-Tagami/pecopeco")
			} else {
				result.notify = true
			}
		}
		return result, nil
	} else {
		return selectRestaurant(restaurantList)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
