package food

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Repository interface {
	List(request ListRequest) (ListResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) List(request ListRequest) (ListResponse, error) {
	uri := os.Getenv("HOTPEPPER_API_URL")
	key := os.Getenv("HOTPEPPER_API_KEY")

	queryParams := url.Values{}
	queryParams.Add("key", key)
	queryParams.Add("keyword", fmt.Sprintf("%s,%s", request.City, request.Food))
	queryParams.Add("count", "100")
	queryParams.Add("format", "json")

	fullURL := fmt.Sprintf("%s?%s", uri, queryParams.Encode())

	req, _ := http.NewRequest("GET", fullURL, nil)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return ListResponse{}, err
	}
	defer res.Body.Close()

	byteArray, _ := io.ReadAll(res.Body)
	listResponse := ListResponse{}
	if err := json.Unmarshal(byteArray, &listResponse); err != nil {
		return ListResponse{}, err
	}
	return listResponse, nil
}
