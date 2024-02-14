package restaurant

type ListResponse struct {
	Restaurants []struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Genre          string `json:"genre"`
		NearestStation string `json:"nearest_station"`
		Address        string `json:"address"`
		URL            string `json:"url"`
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
