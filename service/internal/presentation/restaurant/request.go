package restaurant

type PostRestaurantParams struct {
	Name           string `json:"name" validate:"required"`
	Genre          string `json:"genre" validate:"required"`
	NearestStation string `json:"nearest_station" validate:"required"`
	Address        string `json:"address" validate:"required"`
	URL            string `json:"url" validate:"required"`
	PostedBy       string `json:"posted_by" validate:"required"`
}
