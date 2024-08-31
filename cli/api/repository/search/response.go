package search

type ListRestaurantsByCityAndGenreResponse struct {
	Results struct {
		Shop []struct {
			Name           string `json:"name"`
			Address        string `json:"address"`
			NearestStation string `json:"station_name"`
			Genre          struct {
				Name  string `json:"name"`
				Catch string `json:"catch"`
			} `json:"genre"`
			URLs struct {
				PC string `json:"pc"`
			} `json:"urls"`
		} `json:"shop"`
	} `json:"results"`
}

type ListGenresResponse struct {
	Results struct {
		ResultsAvailable int `json:"results_available"`
		Genre            []struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"genre"`
	}
}
