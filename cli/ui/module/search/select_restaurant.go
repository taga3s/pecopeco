package search

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"github.com/taga3s/pecopeco-cli/api/model"
	"github.com/taga3s/pecopeco-cli/config"
	uiutil "github.com/taga3s/pecopeco-cli/ui/util"
)

type selectRestaurantResult struct {
	Restaurant model.Restaurant
	Share      bool
	Notify     bool
}

func SelectRestaurant(restaurantList []model.Restaurant) (selectRestaurantResult, error) {
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
	uiutil.TextGreen().Printf("---------------------\n[店名] %s\n[住所] %s\n[最寄り駅] %s\n[ジャンル] %s\n[URL] %s\n---------------------\n",
		restaurant.Name,
		restaurant.Address,
		restaurant.NearestStation,
		restaurant.Genre,
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
		result.Restaurant = restaurant

		promptForShare := promptui.Select{
			Label: "Want to share for other users?",
			Items: []string{"Yes", "No"},
		}
		_, share, err := promptForShare.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return selectRestaurantResult{}, err
		}
		if share == "Yes" {
			result.Share = true
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
				uiutil.TextBlue().Println("Sorry, you have not set your personal token to notify your LINE app yet. To notify your LINE app, you can use following command.\n\n```\n> pecopeco config --token <your personal token>\n```")
				time.Sleep(time.Second * 1)
			} else {
				result.Notify = true
			}
		}
		return result, nil
	} else {
		return SelectRestaurant(restaurantList)
	}
}
