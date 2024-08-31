package restaurant

import (
	"context"
	"database/sql"
)

type RestaurantRepository interface {
	SaveWithTx(ctx context.Context, tx *sql.Tx, restaurant *Restaurant) error
	List(ctx context.Context) ([]*Restaurant, error)
	DeleteByIDWithTx(ctx context.Context, tx *sql.Tx, id string) error
}
