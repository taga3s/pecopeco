package genre

type ListResponse struct {
	Results struct {
		ResultsAvailable int `json:"results_available"`
		Genre            []struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"genre"`
	}
}
