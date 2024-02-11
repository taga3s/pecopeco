package user

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	SaveWithTx(ctx context.Context, tx *sql.Tx, user *User) error
	FindById(ctx context.Context, id string) (*User, error)
	CountById(ctx context.Context, id string) (int, error)
}
