package health

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Repository interface {
	HealthCheck() (HealthCheckResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) HealthCheck() (HealthCheckResponse, error) {
	uri := os.Getenv("API_URI")
	req, _ := http.NewRequest("GET", uri+"/health-check", nil)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return HealthCheckResponse{}, err
	}
	defer res.Body.Close()

	byteArray, _ := io.ReadAll(res.Body)

	healthCheckResponse := HealthCheckResponse{}
	if err := json.Unmarshal(byteArray, &healthCheckResponse); err != nil {
		return HealthCheckResponse{}, err
	}

	return healthCheckResponse, nil
}
