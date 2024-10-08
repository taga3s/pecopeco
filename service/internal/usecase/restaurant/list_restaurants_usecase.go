package restaurant

import (
	"context"
	"time"

	restaurantDomain "github.com/taga3s/pecopeco-service/internal/domain/restaurant"
)

type ListRestaurantsUseCase struct {
	restaurantRepo restaurantDomain.RestaurantRepository
}

func NewListRestaurantsUseCase(
	restaurantRepo restaurantDomain.RestaurantRepository,
) *ListRestaurantsUseCase {
	return &ListRestaurantsUseCase{
		restaurantRepo: restaurantRepo,
	}
}

type ListRestaurantsUseCaseDto struct {
	ID             string
	Name           string
	Genre          string
	NearestStation string
	Address        string
	URL            string
	PostedBy       string
	CreatedAt      time.Time
}

func (uc *ListRestaurantsUseCase) Run(ctx context.Context) ([]*ListRestaurantsUseCaseDto, error) {
	restaurants, err := uc.restaurantRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]*ListRestaurantsUseCaseDto, 0, len(restaurants))

	for _, v := range restaurants {
		dtos = append(dtos, &ListRestaurantsUseCaseDto{
			ID:             v.ID,
			Name:           v.Name,
			Genre:          v.Genre,
			NearestStation: v.NearestStation,
			Address:        v.Address,
			URL:            v.URL,
			PostedBy:       v.PostedBy,
			CreatedAt:      v.CreatedAt,
		})
	}
	return dtos, nil
}
