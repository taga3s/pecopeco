package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	genrefactory "github.com/taga3s/pecopeco-cli/api/factory/genre"
	restaurantfactory "github.com/taga3s/pecopeco-cli/api/factory/restaurant"
	"github.com/taga3s/pecopeco-cli/ui/module/search"
	"github.com/taga3s/pecopeco-cli/ui/module/share"
	uiutil "github.com/taga3s/pecopeco-cli/ui/util"
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
		Items: []string{"Search Restaurants", "Check Shared Restaurants", "Exit"},
	}

	_, mode, err := promptForMode.Run()
	if err != nil {
		fmt.Printf("Prompt failed%v\n", err)
		return
	}
	switch mode {
	case "Search Restaurants":
		searchRestaurants()
	case "Check Shared Restaurants":
		showSharedRestaurants()
	case "Exit":
		fmt.Print("Bye!\n")
		os.Exit(1)
	}
}

// -- 検索機能
func searchRestaurants() {
	genreFactory := genrefactory.CreateFactory()
	restaurantFactory := restaurantfactory.CreateFactory()

	genreList, err := genreFactory.ListGenres()
	if err != nil {
		fmt.Println(err)
		return
	}
	searchRestaurantInput := search.GetConditionInput(genreList)
	params := restaurantfactory.ListRestaurantsParams{
		City:  searchRestaurantInput.City,
		Genre: searchRestaurantInput.Genre,
	}
	restaurantList, err := restaurantFactory.ListRestaurants(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(restaurantList) == 0 {
		uiutil.TextBlue().Println("Sorry, there is no data. Please try to change the input parameters.")
		time.Sleep(time.Second * 1)
		run()
	}

	selectRestaurantResult, err := search.SelectRestaurant(restaurantList)
	if err != nil {
		fmt.Println(err)
		return
	}
	// レストラン共有掲示板への追加
	if selectRestaurantResult.Share {
		params := restaurantfactory.PostRestaurantParams{
			Name:           selectRestaurantResult.Restaurant.Name,
			Address:        selectRestaurantResult.Restaurant.Address,
			NearestStation: selectRestaurantResult.Restaurant.NearestStation,
			Genre:          selectRestaurantResult.Restaurant.Genre,
			URL:            selectRestaurantResult.Restaurant.URL,
		}
		_, err := restaurantFactory.PostSharedRestaurant(params)
		if err != nil {
			fmt.Println(err)
		} else {
			uiutil.TextGreen().Println("Successfully shared the restaurant!")
		}
	}

	// LINEへ通知処理
	if selectRestaurantResult.Notify {
		params := restaurantfactory.NotifyRestaurantToLINEParams{
			Restaurant: selectRestaurantResult.Restaurant,
		}
		err := restaurantFactory.NotifyRestaurantToLINE(params)
		if err != nil {
			fmt.Println(err)
		} else {
			uiutil.TextGreen().Println("Notify to your LINE app!")
		}
	}

	// 退出処理
	time.Sleep(time.Second * 1)
	uiutil.Clear()
	run()
}

// -- レストランシェア掲示板
func showSharedRestaurants() {
	restaurantFactory := restaurantfactory.CreateFactory()
	restaurantList, err := restaurantFactory.ListSharedRestaurants()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(restaurantList) == 0 {
		uiutil.TextBlue().Println("Sorry, there is no data.")
		time.Sleep(time.Second * 1)
		run()
		return
	}

	selectedSharedRestaurant, err := share.SelectRestaurant(restaurantList)
	if err != nil {
		fmt.Println(err)
		return
	}

	// LINEへ通知処理
	if selectedSharedRestaurant.Notify {
		params := restaurantfactory.NotifyRestaurantToLINEParams{
			Restaurant: selectedSharedRestaurant.Restaurant,
		}
		err := restaurantFactory.NotifyRestaurantToLINE(params)
		if err != nil {
			fmt.Println(err)
		} else {
			uiutil.TextGreen().Println("Notify to your LINE app!")
		}
	}

	// 退出処理
	if selectedSharedRestaurant.Exit {
		uiutil.Clear()
		run()
		return
	}

	// 退出処理がない限り、再度レストランシェア掲示板を表示
	showSharedRestaurants()
}

func init() {
	rootCmd.AddCommand(runCmd)
}
