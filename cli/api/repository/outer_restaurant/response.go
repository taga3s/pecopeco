package restaurant

// type ListFromServiceResponse struct {
// 	Restaurants struct {
// 		ID             string `json:"id"`
// 		Name           string `json:"name"`
// 		Genre          string `json:"genre"`
// 		NearestStation string `json:"nearest_station"`
// 		Address        string `json:"address"`
// 		URL            string `json:"url"`
// 	} `json:"restaurants"`
// }

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

type NotifyToLINEResponse struct {
	Status int `json:"status"`
}
