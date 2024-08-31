package restaurant

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/taga3s/pecopeco-service/internal/presentation/responder"
	"github.com/taga3s/pecopeco-service/internal/presentation/util/httputil"
	restaurantUseCase "github.com/taga3s/pecopeco-service/internal/usecase/restaurant"
	"github.com/taga3s/pecopeco-service/pkg/validator"
)

type handler struct {
	listRestaurantsUseCase  *restaurantUseCase.ListRestaurantsUseCase
	saveRestaurantUsecase   *restaurantUseCase.SaveRestaurantUseCase
	deleteRestaurantUseCase *restaurantUseCase.DeleteRestaurantUseCase
}

func NewHandler(
	listRestaurantsUseCase *restaurantUseCase.ListRestaurantsUseCase,
	saveRestaurantUsecase *restaurantUseCase.SaveRestaurantUseCase,
	deleteRestaurantUseCase *restaurantUseCase.DeleteRestaurantUseCase,
) handler {
	return handler{
		listRestaurantsUseCase:  listRestaurantsUseCase,
		saveRestaurantUsecase:   saveRestaurantUsecase,
		deleteRestaurantUseCase: deleteRestaurantUseCase,
	}
}

func (h *handler) ListRestaurants(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	dtos, err := h.listRestaurantsUseCase.Run(ctx)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	// 何も取得できなかった場合は、長さ0の配列で返すようにする
	if len(dtos) == 0 {
		responder.ReturnStatusOK(w, ListRestaurantsResponse{Restaurants: []RestaurantResponse{}})
		return
	}

	var response ListRestaurantsResponse

	for _, v := range dtos {
		response.Restaurants = append(
			response.Restaurants,
			RestaurantResponse{
				ID:             v.ID,
				Name:           v.Name,
				Genre:          v.Genre,
				NearestStation: v.NearestStation,
				Address:        v.Address,
				URL:            v.URL,
				CreatedAt:      v.CreatedAt,
			},
		)
	}
	responder.ReturnStatusOK(w, response)
}

func (h *handler) SaveRestaurant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	params := PostRestaurantParams{}
	if err := httputil.ParseJSONRequestBody(r, &params); err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	validate := validator.Get()
	if err := validate.Struct(&params); err != nil {
		responder.ReturnStatusBadRequest(w, err)
		return
	}

	inputDto := restaurantUseCase.SaveRestaurantUseCaseInputDto{
		Name:           params.Name,
		Genre:          params.Genre,
		NearestStation: params.NearestStation,
		Address:        params.Address,
		URL:            params.URL,
	}

	outputDto, err := h.saveRestaurantUsecase.Run(ctx, inputDto)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}

	response := PostRestaurantResponse{
		ID:             outputDto.ID,
		Name:           outputDto.Name,
		Genre:          outputDto.Genre,
		NearestStation: outputDto.NearestStation,
		Address:        outputDto.Address,
		URL:            outputDto.URL,
	}
	responder.ReturnStatusOK(w, response)
}

func (h *handler) DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	restaurantID := chi.URLParam(r, "id")

	err := h.deleteRestaurantUseCase.Run(ctx, restaurantID)
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}
	responder.ReturnStatusNoContent(w)
}
