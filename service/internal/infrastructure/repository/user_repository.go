package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/ayanami77/pecopeco-service/internal/domain/user"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) SaveWithTx(ctx context.Context, tx *sql.Tx, user *user.User) error {
	fields := []string{
		"id",
		"name",
		"email",
	}
	query := fmt.Sprintf(
		"insert into users (%s) values (?, ?, ?)",
		strings.Join(fields, ", "),
	)
	_, err := tx.ExecContext(ctx, query, user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindById(ctx context.Context, id string) (*user.User, error) {
	fields := []string{
		"id",
		"name",
		"email",
	}
	query := fmt.Sprintf(
		"select %s from users where id = ?",
		strings.Join(fields, ", "),
	)

	row := ur.db.QueryRowContext(ctx, query, id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	user := &user.User{}
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) CountById(ctx context.Context, id string) (int, error) {
	query := "select count(*) from users where id = ?"

	rows, err := ur.db.QueryContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	var count int

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			break
		}
	}

	defer rows.Close()

	return count, nil
}
