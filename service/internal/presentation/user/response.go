package user

type FindCurrentlyLoginUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}
