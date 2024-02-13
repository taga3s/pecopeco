package restaurant

type RestaurantResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Genre          string `json:"genre"`
	NearestStation string `json:"nearest_station"`
	Address        string `json:"address"`
	URL            string `json:"url"`
}

type ListRestaurantsResponse struct {
	Restaurants []RestaurantResponse `json:"restaurants"`
}

type PostRestaurantResponse = RestaurantResponse
