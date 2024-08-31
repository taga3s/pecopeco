package restaurant

import "time"

type ListResponse struct {
	Restaurants []struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		Genre          string    `json:"genre"`
		NearestStation string    `json:"nearest_station"`
		Address        string    `json:"address"`
		URL            string    `json:"url"`
		PostedAt       time.Time `json:"created_at"`
	} `json:"restaurants"`
}

type PostResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Genre          string `json:"genre"`
	NearestStation string `json:"nearest_station"`
	Address        string `json:"address"`
	URL            string `json:"url"`
}
