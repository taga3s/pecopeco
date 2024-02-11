package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

// 200
func ReturnStatusOK[T any](w http.ResponseWriter, body T) {
	returnResponse[T](w, body, http.StatusOK)
}

// 400
func ReturnStatusBadRequest(w http.ResponseWriter, err error) {
	response := ErrorResponse{
		Message: err.Error(),
	}
	returnResponse[ErrorResponse](w, response, http.StatusBadRequest)
}

// 401
func ReturnStatusUnauthorized(w http.ResponseWriter, err error) {
	response := ErrorResponse{
		Message: err.Error(),
	}
	returnResponse[ErrorResponse](w, response, http.StatusUnauthorized)
}

// 500
func ReturnStatusInternalServerError(w http.ResponseWriter, err error) {
	response := ErrorResponse{
		Message: err.Error(),
	}
	returnResponse[ErrorResponse](w, response, http.StatusInternalServerError)
}

// 503
func ReturnStatusUnavailable(w http.ResponseWriter, err error) {
	response := ErrorResponse{
		Message: err.Error(),
	}
	returnResponse[ErrorResponse](w, response, http.StatusServiceUnavailable)
}

func returnResponse[T any](w http.ResponseWriter, body T, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response by error '%#v'", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
