package user

import (
	"context"
	"net/mail"
	"unicode/utf8"

	errDomain "github.com/taga3s/pecopeco-service/internal/domain/error"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(
	id string,
	name string,
	email string,
) (*User, error) {
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("名前の値が不正です。")
	}
	// メールアドレスのバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("メールアドレスの値が不正です。")
	}
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}

const (
	nameLengthMax = 255
	nameLengthMin = 1
)

type UserDomainService interface {
	Exists(ctx context.Context, user *User) (bool, error)
}
