package health

import (
	"fmt"

	"github.com/ayanami77/pecopeco-cli/api/model"
	"github.com/ayanami77/pecopeco-cli/api/repository/health"
)

type HealthFactory interface {
	HealthCheck() (model.Health, error)
}

type factory struct {
	repository health.Repository
}

func CreateFactory() HealthFactory {
	repository := health.New()
	return &factory{repository}
}

func (f *factory) HealthCheck() (model.Health, error) {
	res, err := f.repository.HealthCheck()
	if err != nil {
		err := fmt.Errorf("Error: %v", err)
		return model.Health{}, err
	}
	health := model.Health{
		Status:  res.Status,
		Message: res.Message,
	}
	return health, nil
}
