package search

import (
	"fmt"

	"github.com/taga3s/pecopeco-cli/api/client/app"
)

type Repository interface {
	ListRestaurantsByCityAndGenre(request ListRestaurantsByCityAndGenreRequest) (ListRestaurantsByCityAndGenreResponse, error)
	ListGenres() (ListGenresResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) ListRestaurantsByCityAndGenre(request ListRestaurantsByCityAndGenreRequest) (ListRestaurantsByCityAndGenreResponse, error) {
	listResponse := ListRestaurantsByCityAndGenreResponse{}
	if err := app.HttpClient("GET", fmt.Sprintf("/search/restaurants?city=%s&genre=%s", request.City, request.Genre), nil, &listResponse); err != nil {
		return ListRestaurantsByCityAndGenreResponse{}, err
	}
	return listResponse, nil
}

func (r *repository) ListGenres() (ListGenresResponse, error) {
	listResponse := ListGenresResponse{}
	if err := app.HttpClient("GET", "/search/genres", nil, &listResponse); err != nil {
		return ListGenresResponse{}, err
	}
	return listResponse, nil
}
