package restaurant

import "time"

type RestaurantResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Genre          string    `json:"genre"`
	NearestStation string    `json:"nearest_station"`
	Address        string    `json:"address"`
	URL            string    `json:"url"`
	CreatedAt      time.Time `json:"created_at"`
}

type ListRestaurantsResponse struct {
	Restaurants []RestaurantResponse `json:"restaurants"`
}

type PostRestaurantResponse = RestaurantResponse
