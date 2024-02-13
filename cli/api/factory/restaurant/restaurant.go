package restaurant

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	innerRestaurant "github.com/Seiya-Tagami/pecopeco-cli/api/repository/inner_restaurant"
	outerRestaurant "github.com/Seiya-Tagami/pecopeco-cli/api/repository/outer_restaurant"
)

type RestaurantFactory interface {
	ListRestaurants(params ListRestaurantsParams) ([]model.Restaurant, error)
	NotifyRestaurantToLINE(params NotifyRestaurantToLINEParams) error
	ListFavorites() ([]model.Restaurant, error)
	PostRestaurant(params PostRestaurantParams) (model.Restaurant, error)
	DeleteRestaurant(params DeleteRestaurantParams) error
}

type factory struct {
	innerRepository innerRestaurant.Repository
	outerRepository outerRestaurant.Repository
}

func CreateFactory() RestaurantFactory {
	innerRepository := innerRestaurant.New()
	outerRepository := outerRestaurant.New()
	return &factory{
		innerRepository: innerRepository,
		outerRepository: outerRepository,
	}
}

func (f *factory) ListRestaurants(params ListRestaurantsParams) ([]model.Restaurant, error) {
	request := outerRestaurant.ListRequest{
		City:  params.City,
		Genre: params.Genre,
	}
	res, err := f.outerRepository.List(request)
	if err != nil {
		err := fmt.Errorf("Failed to implement List: %v", err)
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

func (f *factory) NotifyRestaurantToLINE(params NotifyRestaurantToLINEParams) error {
	request := outerRestaurant.NotifyToLINERequest{
		Name:        params.Restaurant.Name,
		Address:     params.Restaurant.Address,
		StationName: params.Restaurant.StationName,
		GenreName:   params.Restaurant.GenreName,
		URL:         params.Restaurant.URL,
	}
	err := f.outerRepository.NotifyToLINE(request)
	if err != nil {
		err := fmt.Errorf("Failed to implement NotifyToLINE: %v", err)
		return err
	}
	return nil
}

func (f *factory) ListFavorites() ([]model.Restaurant, error) {
	res, err := f.innerRepository.List()
	if err != nil {
		err := fmt.Errorf("Failed to implement List: %v", err)
		return []model.Restaurant{}, err
	}

	restaurantList := make([]model.Restaurant, 0, len(res.Restaurants))

	for _, v := range res.Restaurants {
		restaurant := model.Restaurant{
			ID:          v.ID,
			Name:        v.Name,
			Address:     v.Address,
			StationName: v.NearestStation,
			GenreName:   v.Genre,
			URL:         v.URL,
		}
		restaurantList = append(restaurantList, restaurant)
	}
	return restaurantList, nil
}

func (f *factory) PostRestaurant(params PostRestaurantParams) (model.Restaurant, error) {
	request := innerRestaurant.PostRequest{
		Name:        params.Name,
		Address:     params.Address,
		StationName: params.StationName,
		GenreName:   params.GenreName,
		URL:         params.URL,
	}
	res, err := f.innerRepository.Post(request)
	if err != nil {
		err := fmt.Errorf("Failed to implement List: %v", err)
		return model.Restaurant{}, err
	}
	return model.Restaurant{
		ID:          res.ID,
		Name:        res.Name,
		Address:     res.Address,
		StationName: res.NearestStation,
		GenreName:   res.Genre,
		URL:         res.URL,
	}, nil
}

func (f *factory) DeleteRestaurant(params DeleteRestaurantParams) error {
	request := innerRestaurant.DeleteRequest{
		ID: params.ID,
	}
	err := f.innerRepository.Delete(request)
	if err != nil {
		return err
	}
	return err
}
