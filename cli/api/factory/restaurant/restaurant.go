package restaurant

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/Seiya-Tagami/pecopeco-cli/api/repository/restaurant"
)

type RestaurantFactory interface {
	ListRestaurants(params ListRestaurantsParams) ([]model.Restaurant, error)
}

type factory struct {
	repository restaurant.Repository
}

func CreateFactory() RestaurantFactory {
	repository := restaurant.New()
	return &factory{repository}
}

func (f *factory) ListRestaurants(params ListRestaurantsParams) ([]model.Restaurant, error) {
	request := restaurant.ListRequest{
		City: params.City,
		Food: params.Food,
	}
	res, err := f.repository.List(request)
	if err != nil {
		err := fmt.Errorf("Failed to implement Get restaurantList: %v", err)
		return []model.Restaurant{}, err
	}

	restaurantList := make([]model.Restaurant, 0, len(res.Results.Shop))

	for _, v := range res.Results.Shop {
		restaurant := model.Restaurant{
			Name:        v.Name,
			Address:     v.Address,
			StationName: v.StationName,
			GenreName:   v.Genre.Name,
			URL:         v.URLs.PC,
		}
		restaurantList = append(restaurantList, restaurant)
	}
	return restaurantList, nil
}
