package favorites

// import (
// 	"fmt"
// 	"time"

// 	"github.com/manifoldco/promptui"
// 	"github.com/spf13/viper"
// 	"github.com/taga3s/pecopeco-cli/api/model"
// 	"github.com/taga3s/pecopeco-cli/config"
// 	uiutil "github.com/taga3s/pecopeco-cli/ui/util"
// )

// type selectRestaurantResult struct {
// 	Restaurant model.Restaurant
// 	Notify     bool
// 	Delete     bool
// 	Exit       bool
// }

// func SelectRestaurant(restaurantList []model.Restaurant) (selectRestaurantResult, error) {
// 	restaurantMap := map[string]model.Restaurant{}
// 	options := make([]string, 0, len(restaurantList)+1)

// 	for _, v := range restaurantList {
// 		restaurantMap[v.Name] = v
// 		options = append(options, v.Name)
// 	}
// 	options = append(options, "Back to Menu")

// 	promptForOptions := promptui.Select{
// 		Label: "Your favorite restaurants. Please select to show details",
// 		Items: options,
// 	}

// 	_, option, err := promptForOptions.Run()
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		return selectRestaurantResult{}, err
// 	}

// 	if option == "Back to Menu" {
// 		result := selectRestaurantResult{
// 			Exit: true,
// 		}
// 		return result, nil
// 	}

// 	restaurant := restaurantMap[option]
// 	uiutil.TextGreen().Printf("---------------------\n[店名] %s\n[住所] %s\n[最寄り駅] %s\n[ジャンル] %s\n[URL] %s\n---------------------\n",
// 		restaurant.Name,
// 		restaurant.Address,
// 		restaurant.NearestStation,
// 		restaurant.Genre,
// 		restaurant.URL,
// 	)

// 	promptForDecision := promptui.Select{
// 		Label: "What would you like to do?",
// 		Items: []string{"Notify your LINE app", "Delete from Favorites", "Back to Favorites"},
// 	}

// 	_, decision, err := promptForDecision.Run()
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		return selectRestaurantResult{}, err
// 	}

// 	result := selectRestaurantResult{}
// 	result.Restaurant = restaurant

// 	if decision == "Notify your LINE app" {
// 		// トークンがセットされていない場合、ここで弾くようにする。
// 		if viper.GetString(config.LINE_NOTIFY_API_TOKEN) == "" {
// 			uiutil.TextBlue().Println("Sorry, you have not set your personal token to notify your line app yet. To notify your line app, you can use following command.\n\n```\n> pecopeco config --token <your personal token>\n```")
// 			time.Sleep(time.Second * 1)
// 		} else {
// 			result.Notify = true
// 		}
// 		return result, nil
// 	}

// 	if decision == "Delete from Favorites" {
// 		promptForConfirm := promptui.Select{
// 			Label: "Do you confirm?",
// 			Items: []string{"Yes", "No"},
// 		}
// 		_, answer, err := promptForConfirm.Run()
// 		if err != nil {
// 			return result, err
// 		}
// 		if answer == "Yes" {
// 			result.Delete = true
// 		}
// 		return result, nil
// 	}
// 	return result, nil
// }
