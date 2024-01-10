package food

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/Seiya-Tagami/pecopeco-cli/api/repository/food"
)

type factory struct {
	repository food.Repository
}

func CreateFactory() model.FoodFactory {
	repository := food.New()
	return &factory{repository}
}

func (f *factory) ListFood(request food.ListRequest) ([]model.Food, error) {
	res, err := f.repository.List(request)
	if err != nil {
		err := fmt.Errorf("Failed to implement Get FoodList: %v", err)
		return []model.Food{}, err
	}

	foodList := []model.Food{}

	for _, v := range res.Results.Shop {
		food := model.Food{
			Name:        v.Name,
			Address:     v.Address,
			StationName: v.StationName,
			GenreName:   v.Genre.Name,
			URL:         v.URLs.PC,
		}
		foodList = append(foodList, food)
	}
	return foodList, nil
}
