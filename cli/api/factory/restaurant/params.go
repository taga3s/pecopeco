package restaurant

import "github.com/Seiya-Tagami/pecopeco-cli/api/model"

type ListRestaurantsParams struct {
	City  string
	Genre string
}

type NotifyRestaurantToLINEParams struct {
	Restaurant model.Restaurant
}

type PostRestaurantParams struct {
	Name           string
	Address        string
	NearestStation string
	Genre          string
	URL            string
}

type DeleteRestaurantParams struct {
	ID string
}
