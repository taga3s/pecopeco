package restaurant

import (
	"context"

	"github.com/Seiya-Tagami/pecopeco-service/internal/db"
	restaurantDomain "github.com/Seiya-Tagami/pecopeco-service/internal/domain/restaurant"
)

type SaveRestaurantUseCase struct {
	restaurantRepo restaurantDomain.RestaurantRepository
}

func NewSaveRestaurantUseCase(
	restaurantRepo restaurantDomain.RestaurantRepository,
) *SaveRestaurantUseCase {
	return &SaveRestaurantUseCase{
		restaurantRepo: restaurantRepo,
	}
}

type SaveRestaurantUseCaseInputDto struct {
	Name           string
	Genre          string
	NearestStation string
	Address        string
	URL            string
	UserID         string
}

type SaveRestaurantUseCaseOutputDto struct {
	ID             string
	Name           string
	Genre          string
	NearestStation string
	Address        string
	URL            string
	UserID         string
}

func (uc *SaveRestaurantUseCase) Run(ctx context.Context, dto SaveRestaurantUseCaseInputDto) (*SaveRestaurantUseCaseOutputDto, error) {
	restaurant, err := restaurantDomain.NewRestaurant(dto.Name, dto.Genre, dto.NearestStation, dto.Address, dto.URL, dto.UserID)
	if err != nil {
		return nil, err
	}

	db := db.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	if err = uc.restaurantRepo.SaveWithTx(ctx, tx, restaurant); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return &SaveRestaurantUseCaseOutputDto{
		ID:             restaurant.ID,
		Name:           restaurant.Name,
		Genre:          restaurant.Genre,
		NearestStation: restaurant.NearestStation,
		Address:        restaurant.Address,
		URL:            restaurant.URL,
		UserID:         restaurant.UserID,
	}, nil
}
