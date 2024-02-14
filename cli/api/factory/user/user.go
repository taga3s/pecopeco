package user

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/Seiya-Tagami/pecopeco-cli/api/repository/user"
)

type UserFactory interface {
	FindUser() (model.User, error)
	Login(params LoginParams) (model.User, error)
}

type factory struct {
	repository user.Repository
}

func CreateFactory() UserFactory {
	repository := user.New()
	return &factory{repository}
}

func (f *factory) FindUser() (model.User, error) {
	response, err := f.repository.Get()
	if err != nil {
		err := fmt.Errorf("Error: %v", err)
		return model.User{}, err
	}
	user := model.User{
		ID:    response.ID,
		Name:  response.Name,
		Email: response.Email,
	}
	return user, nil
}

func (f *factory) Login(params LoginParams) (model.User, error) {
	request := user.LoginRequest{
		ID:    params.ID,
		Name:  params.Name,
		Email: params.Email,
	}
	response, err := f.repository.Create(request)
	if err != nil {
		err := fmt.Errorf("Error: %v", err)
		return model.User{}, err
	}
	user := model.User{
		ID:    response.ID,
		Name:  response.Name,
		Email: response.Email,
	}
	return user, nil
}
