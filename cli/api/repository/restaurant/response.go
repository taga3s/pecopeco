package restaurant

type ListResponse struct {
	Results struct {
		Shop []struct {
			Name        string `json:"name"`
			Address     string `json:"address"`
			StationName string `json:"station_name"`
			Genre       struct {
				Name  string `json:"name"`
				Catch string `json:"catch"`
			} `json:"genre"`
			URLs struct {
				PC string `json:"pc"`
			} `json:"urls"`
		} `json:"shop"`
	} `json:"results"`
}
