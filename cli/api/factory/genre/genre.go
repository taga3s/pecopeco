package genre

import (
	"fmt"

	"github.com/taga3s/pecopeco-cli/api/model"
	"github.com/taga3s/pecopeco-cli/api/repository/genre"
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
