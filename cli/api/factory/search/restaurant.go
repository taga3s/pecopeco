package search

import (
	"fmt"

	"github.com/taga3s/pecopeco-cli/api/model"
	search "github.com/taga3s/pecopeco-cli/api/repository/search"
)

type SearchFactory interface {
	ListRestaurantsByCityAndGenre(params ListRestaurantsByCityAndGenreParams) ([]model.Restaurant, error)
	ListGenres() ([]model.Genre, error)
}

type factory struct {
	repository search.Repository
}

func CreateFactory() SearchFactory {
	repository := search.New()
	return &factory{
		repository: repository,
	}
}

func (f *factory) ListRestaurantsByCityAndGenre(params ListRestaurantsByCityAndGenreParams) ([]model.Restaurant, error) {
	request := search.ListRestaurantsByCityAndGenreRequest{
		City:  params.City,
		Genre: params.Genre,
	}
	res, err := f.repository.ListRestaurantsByCityAndGenre(request)
	if err != nil {
		err := fmt.Errorf("error: %v", err)
		return []model.Restaurant{}, err
	}

	restaurantList := make([]model.Restaurant, 0, len(res.Results.Shop))

	for _, v := range res.Results.Shop {
		restaurant := model.Restaurant{
			Name:           v.Name,
			Address:        v.Address,
			NearestStation: v.NearestStation,
			Genre:          v.Genre.Name,
			URL:            v.URLs.PC,
		}
		restaurantList = append(restaurantList, restaurant)
	}

	return restaurantList, nil
}

func (f *factory) ListGenres() ([]model.Genre, error) {
	res, err := f.repository.ListGenres()
	if err != nil {
		err := fmt.Errorf("error: %v", err)
		return []model.Genre{}, err
	}

	genreList := make([]model.Genre, 0, len(res.Results.Genre))

	for _, v := range res.Results.Genre {
		genre := model.Genre{
			Name: v.Name,
			Code: v.Code,
		}
		genreList = append(genreList, genre)
	}
	return genreList, nil
}
