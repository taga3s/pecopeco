package health

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/model"
	"github.com/Seiya-Tagami/pecopeco-cli/api/repository/health"
)

type factory struct {
	repository health.Repository
}

func CreateFactory() model.HealthFactory {
	repository := health.New()
	return &factory{repository}
}

func (f *factory) HealthCheck() (model.Health, error) {
	res, err := f.repository.HealthCheck()
	if err != nil {
		err := fmt.Errorf("Failed to implement health check: %v", err)
		return model.Health{}, err
	}
	health := model.Health{
		Status:  res.Status,
		Message: res.Message,
	}
	return health, nil
}
