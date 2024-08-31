package restaurant

import (
	"github.com/taga3s/pecopeco-cli/api/client/app"
)

type Repository interface {
	List() (ListResponse, error)
	Post(request PostRequest) (PostResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) List() (ListResponse, error) {
	response := ListResponse{}
	if err := app.HttpClient("GET", "/restaurants", nil, &response); err != nil {
		return ListResponse{}, err
	}
	return response, nil
}

func (r *repository) Post(request PostRequest) (PostResponse, error) {
	response := PostResponse{}
	if err := app.HttpClient("POST", "/restaurants", request, &response); err != nil {
		return PostResponse{}, err
	}
	return response, nil
}
