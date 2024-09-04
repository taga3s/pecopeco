package restaurant

import (
	"fmt"

	"github.com/taga3s/pecopeco-cli/api/model"
	notifyToLine "github.com/taga3s/pecopeco-cli/api/repository/notify_to_line"
	restaurant "github.com/taga3s/pecopeco-cli/api/repository/restaurant"
)

type RestaurantFactory interface {
	NotifyRestaurantToLINE(params NotifyRestaurantToLINEParams) error
	ListSharedRestaurants() ([]model.Restaurant, error)
	PostSharedRestaurant(params PostRestaurantParams) (model.Restaurant, error)
}

type factory struct {
	innerRepository restaurant.Repository
	outerRepository notifyToLine.Repository
}

func CreateFactory() RestaurantFactory {
	innerRepository := restaurant.New()
	outerRepository := notifyToLine.New()
	return &factory{
		innerRepository: innerRepository,
		outerRepository: outerRepository,
	}
}

func (f *factory) NotifyRestaurantToLINE(params NotifyRestaurantToLINEParams) error {
	request := notifyToLine.NotifyToLINERequest{
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
			PostedBy:       v.PostedBy,
			PostedAt:       v.PostedAt,
		}
		restaurantList = append(restaurantList, restaurant)
	}
	return restaurantList, nil
}

func (f *factory) PostSharedRestaurant(params PostRestaurantParams) (model.Restaurant, error) {
	request := restaurant.PostRequest{
		Name:           params.Name,
		Address:        params.Address,
		NearestStation: params.NearestStation,
		Genre:          params.Genre,
		URL:            params.URL,
		PostedBy:       params.PostedBy,
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
