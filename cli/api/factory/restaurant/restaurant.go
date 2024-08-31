package restaurant

import (
	"fmt"

	"github.com/taga3s/pecopeco-cli/api/model"
	innerRestaurant "github.com/taga3s/pecopeco-cli/api/repository/inner_restaurant"
	outerRestaurant "github.com/taga3s/pecopeco-cli/api/repository/outer_restaurant"
)

type RestaurantFactory interface {
	ListRestaurants(params ListRestaurantsParams) ([]model.Restaurant, error)
	NotifyRestaurantToLINE(params NotifyRestaurantToLINEParams) error
	ListSharedRestaurants() ([]model.Restaurant, error)
	PostSharedRestaurant(params PostRestaurantParams) (model.Restaurant, error)
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
		err := fmt.Errorf("error: %v", err)
		return []model.Restaurant{}, err
	}

	restaurantList := make([]model.Restaurant, 0, len(res.Results.Shop))

	for _, v := range res.Results.Shop {
		restaurant := model.Restaurant{
			Name:           v.Name,
			Address:        v.Address,
			NearestStation: v.NearestStation,
			Genre:          v.Genre.Name,
			URL:            v.URLs.PC,
		}
		restaurantList = append(restaurantList, restaurant)
	}

	return restaurantList, nil
}

func (f *factory) NotifyRestaurantToLINE(params NotifyRestaurantToLINEParams) error {
	request := outerRestaurant.NotifyToLINERequest{
		Name:           params.Restaurant.Name,
		Address:        params.Restaurant.Address,
		NearestStation: params.Restaurant.NearestStation,
		Genre:          params.Restaurant.Genre,
		URL:            params.Restaurant.URL,
	}
	err := f.outerRepository.NotifyToLINE(request)
	if err != nil {
		err := fmt.Errorf("error: %v", err)
		return err
	}
	return nil
}

func (f *factory) ListSharedRestaurants() ([]model.Restaurant, error) {
	res, err := f.innerRepository.List()
	if err != nil {
		err := fmt.Errorf("error: %v", err)
		return []model.Restaurant{}, err
	}

	restaurantList := make([]model.Restaurant, 0, len(res.Restaurants))

	for _, v := range res.Restaurants {
		restaurant := model.Restaurant{
			ID:             v.ID,
			Name:           v.Name,
			Address:        v.Address,
			NearestStation: v.NearestStation,
			Genre:          v.Genre,
			URL:            v.URL,
			PostedAt:       v.PostedAt,
		}
		restaurantList = append(restaurantList, restaurant)
	}
	return restaurantList, nil
}

func (f *factory) PostSharedRestaurant(params PostRestaurantParams) (model.Restaurant, error) {
	request := innerRestaurant.PostRequest{
		Name:           params.Name,
		Address:        params.Address,
		NearestStation: params.NearestStation,
		Genre:          params.Genre,
		URL:            params.URL,
	}
	res, err := f.innerRepository.Post(request)
	if err != nil {
		err := fmt.Errorf("error: %v", err)
		return model.Restaurant{}, err
	}
	return model.Restaurant{
		ID:             res.ID,
		Name:           res.Name,
		Address:        res.Address,
		NearestStation: res.NearestStation,
		Genre:          res.Genre,
		URL:            res.URL,
	}, nil
}
