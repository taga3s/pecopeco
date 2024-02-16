package health

import (
	"github.com/ayanami77/pecopeco-cli/api/client/app"
)

type Repository interface {
	HealthCheck() (HealthCheckResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) HealthCheck() (HealthCheckResponse, error) {
	healthCheckResponse := HealthCheckResponse{}
	if err := app.HttpClient("GET", "/health-check", nil, &healthCheckResponse); err != nil {
		return healthCheckResponse, nil
	}
	return healthCheckResponse, nil
}
