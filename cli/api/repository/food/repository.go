package food

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/client/hotpepper"
)

type Repository interface {
	List(request ListRequest) (ListResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) List(request ListRequest) (ListResponse, error) {
	listResponse := ListResponse{}
	if err := hotpepper.HttpClient("GET", fmt.Sprintf("&keyword=%s,%s&count=100&format=json", request.City, request.Food), &listResponse); err != nil {
		return ListResponse{}, err
	}
	return listResponse, nil
}
