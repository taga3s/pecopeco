package healthhandler

type HealthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"string,omitempty"`
}
