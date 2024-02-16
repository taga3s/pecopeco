package restaurant

import (
	"context"

	restaurantDomain "github.com/ayanami77/pecopeco-service/internal/domain/restaurant"
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
	UserID         string
}

func (uc *ListRestaurantsUseCase) Run(ctx context.Context, userID string) ([]*ListRestaurantsUseCaseDto, error) {
	restaurants, err := uc.restaurantRepo.ListByUserID(ctx, userID)
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
			UserID:         v.UserID,
		})
	}
	return dtos, nil
}
