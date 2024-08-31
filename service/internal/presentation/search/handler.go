package search

import (
	"fmt"
	"net/http"

	"github.com/taga3s/pecopeco-service/internal/presentation/responder"
	"github.com/taga3s/pecopeco-service/internal/presentation/search/client"
)

type handler struct{}

func NewHandler() handler {
	return handler{}
}

func (h *handler) ListGenres(w http.ResponseWriter, r *http.Request) {
	var response ListGenresResponse

	if err := client.HttpClient("GET", "/genre/v1/", "&format=json", &response); err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	responder.ReturnStatusOK(w, response)
}

func (h *handler) ListRestaurantsByCityAndGenre(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	genre := r.URL.Query().Get("genre")

	var response ListRestaurantsByCityAndGenreResponse

	if err := client.HttpClient("GET", "/gourmet/v1/", fmt.Sprintf("&keyword=%s&genre=%s&count=100&format=json", city, genre), &response); err != nil {
		responder.ReturnStatusInternalServerError(w, err)
	}

	responder.ReturnStatusOK(w, response)
}
