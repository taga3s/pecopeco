package health

type HealthCheckResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
