package restaurant

import (
	"context"

	"github.com/ayanami77/pecopeco-service/internal/db"
	"github.com/ayanami77/pecopeco-service/internal/domain/restaurant"
)

type DeleteRestaurantUseCase struct {
	restaurantRepo restaurant.RestaurantRepository
}

func NewDeleteRestaurantsUseCase(
	restaurantRepo restaurant.RestaurantRepository,
) *DeleteRestaurantUseCase {
	return &DeleteRestaurantUseCase{
		restaurantRepo: restaurantRepo,
	}
}

func (uc *DeleteRestaurantUseCase) Run(ctx context.Context, id string) error {
	db := db.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = uc.restaurantRepo.DeleteByIDWithTx(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
