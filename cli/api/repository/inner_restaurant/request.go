package restaurant

type PostRequest struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	NearestStation string `json:"nearest_station"`
	Genre          string `json:"genre"`
	URL            string `json:"url"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
