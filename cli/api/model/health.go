package model

type Health struct {
	Status  int
	Message string
}

type HealthFactory interface {
	HealthCheck() (Health, error)
}
