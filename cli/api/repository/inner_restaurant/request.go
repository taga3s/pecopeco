package restaurant

type PostRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	StationName string `json:"nearest_station"`
	GenreName   string `json:"genre"`
	URL         string `json:"url"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
