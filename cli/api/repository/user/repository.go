package user

import (
	"github.com/taga3s/pecopeco-cli/api/client/app"
)

type Repository interface {
	Get() (FindUserResponse, error)
	Create(request LoginRequest) (LoginResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) Get() (FindUserResponse, error) {
	findUserResponse := FindUserResponse{}
	if err := app.HttpClient("GET", "/users/me", nil, &findUserResponse); err != nil {
		return findUserResponse, err
	}
	return findUserResponse, nil
}

func (r *repository) Create(request LoginRequest) (LoginResponse, error) {
	loginResponse := LoginResponse{}
	if err := app.HttpClient("POST", "/users/login", request, &loginResponse); err != nil {
		return loginResponse, err
	}
	return loginResponse, nil
}
