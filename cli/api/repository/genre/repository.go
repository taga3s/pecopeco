package genre

import (
	"github.com/taga3s/pecopeco-cli/api/client/hotpepper"
)

type Repository interface {
	List() (ListResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) List() (ListResponse, error) {
	listResponse := ListResponse{}
	if err := hotpepper.HttpClient("GET", "/genre/v1/", "&format=json", &listResponse); err != nil {
		return ListResponse{}, err
	}
	return listResponse, nil
}
