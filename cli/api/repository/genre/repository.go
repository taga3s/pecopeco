package genre

import "github.com/taga3s/pecopeco-cli/api/client/app"

type Repository interface {
	List() (ListResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) List() (ListResponse, error) {
	listResponse := ListResponse{}
	if err := app.HttpClient("GET", "/search/genres", nil, &listResponse); err != nil {
		return ListResponse{}, err
	}
	return listResponse, nil
}
