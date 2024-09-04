package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/taga3s/pecopeco-service/internal/domain/restaurant"
)

type restaurantRepository struct {
	db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) restaurant.RestaurantRepository {
	return &restaurantRepository{db: db}
}

func (rr *restaurantRepository) List(ctx context.Context) ([]*restaurant.Restaurant, error) {
	fields := []string{
		"id",
		"name",
		"genre",
		"nearest_station",
		"address",
		"url",
		"posted_by",
		"created_at",
	}

	query := fmt.Sprintf(
		"select %s from restaurants",
		strings.Join(fields, ", "),
	)

	rows, err := rr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	restaurants := []*restaurant.Restaurant{}

	for rows.Next() {
		restaurant := &restaurant.Restaurant{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Genre,
			&restaurant.NearestStation,
			&restaurant.Address,
			&restaurant.URL,
			&restaurant.PostedBy,
			&restaurant.CreatedAt,
		)
		if err != nil {
			break
		}
		restaurants = append(restaurants, restaurant)
	}

	defer rows.Close()

	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (rr *restaurantRepository) SaveWithTx(ctx context.Context, tx *sql.Tx, restaurant *restaurant.Restaurant) error {
	fields := []string{
		"id",
		"name",
		"genre",
		"nearest_station",
		"address",
		"url",
		"posted_by",
	}
	query := fmt.Sprintf(
		"insert into restaurants (%s) values (?, ?, ?, ?, ?, ?, ?)",
		strings.Join(fields, ", "),
	)
	_, err := tx.ExecContext(
		ctx,
		query,
		restaurant.ID,
		restaurant.Name,
		restaurant.Genre,
		restaurant.NearestStation,
		restaurant.Address,
		restaurant.URL,
		restaurant.PostedBy,
	)
	if err != nil {
		return err
	}

	return nil
}

func (rr *restaurantRepository) DeleteByIDWithTx(ctx context.Context, tx *sql.Tx, id string) error {
	query := "delete from restaurants where id = ?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
