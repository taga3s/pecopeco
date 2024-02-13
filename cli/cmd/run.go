package cmd

import (
	"fmt"
	"os"
	"time"

	genrefactory "github.com/Seiya-Tagami/pecopeco-cli/api/factory/genre"
	restaurantfactory "github.com/Seiya-Tagami/pecopeco-cli/api/factory/restaurant"
	"github.com/Seiya-Tagami/pecopeco-cli/config"
	"github.com/Seiya-Tagami/pecopeco-cli/ui/module/favorites"
	"github.com/Seiya-Tagami/pecopeco-cli/ui/module/search"
	uiutil "github.com/Seiya-Tagami/pecopeco-cli/ui/util"
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
		Items: []string{"Search Restaurants", "Show Favorites", "Exit"},
	}

	_, mode, err := promptForMode.Run()
	if err != nil {
		fmt.Printf("Prompt failed%v\n", err)
		return
	}
	switch mode {
	case "Search Restaurants":
		searchRestaurants()
	case "Show Favorites":
		showFavorites()
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
	searchRestaurantInput := search.GetSearchRestaurantInput(genreList)
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
	// お気に入り登録処理
	if selectRestaurantResult.AddToFavorites {
		params := restaurantfactory.PostRestaurantParams{
			Name:        selectRestaurantResult.Restaurant.Name,
			Address:     selectRestaurantResult.Restaurant.Address,
			StationName: selectRestaurantResult.Restaurant.StationName,
			GenreName:   selectRestaurantResult.Restaurant.GenreName,
			URL:         selectRestaurantResult.Restaurant.URL,
		}
		_, err := restaurantFactory.PostRestaurant(params)
		if err != nil {
			fmt.Println(err)
		} else {
			uiutil.TextGreen().Println("Added to favorites!")
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
			uiutil.TextGreen().Println("Notify to your line app!")
		}
	}

	// 退出処理
	time.Sleep(time.Second * 1)
	run()
}

// -- お気に入り機能
func showFavorites() {
	if !config.IsLogin() {
		uiutil.TextBlue().Println("Sorry, to add to favorites, you have to login first. Please login with following command.\n> pecopeco login")
		time.Sleep(time.Second * 1)
		run()
		return
	}

	restaurantFactory := restaurantfactory.CreateFactory()
	restaurantList, err := restaurantFactory.ListFavorites()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(restaurantList) == 0 {
		uiutil.TextBlue().Println("Sorry, there is no data. Please try to add restaurants by searching.")
		time.Sleep(time.Second * 1)
		run()
		return
	}

	selectedRestaurant, err := favorites.SelectRestaurant(restaurantList)
	if err != nil {
		fmt.Println(err)
		return
	}

	// LINEへ通知処理
	if selectedRestaurant.Notify {
		params := restaurantfactory.NotifyRestaurantToLINEParams{
			Restaurant: selectedRestaurant.Restaurant,
		}
		err := restaurantFactory.NotifyRestaurantToLINE(params)
		if err != nil {
			fmt.Println(err)
		} else {
			uiutil.TextGreen().Println("Notify to your line app!")
		}
	}

	// 削除処理
	if selectedRestaurant.Delete {
		params := restaurantfactory.DeleteRestaurantParams{
			ID: selectedRestaurant.Restaurant.ID,
		}
		err := restaurantFactory.DeleteRestaurant(params)
		if err != nil {
			fmt.Println(err)
		} else {
			uiutil.TextGreen().Println("Securely delete from Favorites.")
		}
	}

	// 退出処理
	if selectedRestaurant.Exit {
		run()
		return
	}

	// 退出処理がない限り、基本お気に入りリストを表示し続ける
	showFavorites()
}
func init() {
	rootCmd.AddCommand(runCmd)
}
