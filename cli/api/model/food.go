package model

import "github.com/Seiya-Tagami/pecopeco-cli/api/repository/food"

type Food struct {
	Name        string
	Address     string
	StationName string
	GenreName   string
	URL         string
}

type FoodFactory interface {
	GetFoodList(request food.GetFoodListRequest) ([]Food, error)
}
