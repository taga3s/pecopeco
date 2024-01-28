package restaurant

import "github.com/Seiya-Tagami/pecopeco-cli/api/model"

type ListRestaurantsParams struct {
	City string
	Food string
}

type NotifyRestaurantToLINEParams struct {
	Restaurant model.Restaurant
}
