package genre

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/Seiya-Tagami/pecopeco-cli/api/repository/genre"
)

type GenreFactory interface {
	ListGenres() ([]model.Genre, error)
}

type factory struct {
	repository genre.Repository
}

func CreateFactory() GenreFactory {
	repository := genre.New()
	return &factory{repository}
}

func (f *factory) ListGenres() ([]model.Genre, error) {
	res, err := f.repository.List()
	if err != nil {
		err := fmt.Errorf("Failed to implement List: %v", err)
		return []model.Genre{}, err
	}

	genresList := make([]model.Genre, 0, res.Results.ResultsAvailable)

	for _, v := range res.Results.Genre {
		genre := model.Genre{
			Name: v.Name,
			Code: v.Code,
		}
		genresList = append(genresList, genre)
	}
	return genresList, nil
}
