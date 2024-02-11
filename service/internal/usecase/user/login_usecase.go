package user

import (
	"context"

	"github.com/Seiya-Tagami/pecopeco-service/internal/db"
	userDomain "github.com/Seiya-Tagami/pecopeco-service/internal/domain/user"
)

type LoginUsecase struct {
	userDomainService userDomain.UserDomainService
	userRepo          userDomain.UserRepository
}

func NewLoginUsecase(
	userDomainService userDomain.UserDomainService,
	userRepo userDomain.UserRepository,
) *LoginUsecase {
	return &LoginUsecase{
		userDomainService: userDomainService,
		userRepo:          userRepo,
	}
}

type LoginUseCaseInputDto struct {
	ID    string
	Name  string
	Email string
}

type LoginUseCaseOutputDto struct {
	ID    string
	Name  string
	Email string
}

func (uc *LoginUsecase) Run(ctx context.Context, dto LoginUseCaseInputDto) (*LoginUseCaseOutputDto, error) {
	user, err := userDomain.NewUser(dto.ID, dto.Name, dto.Email)
	if err != nil {
		return nil, err
	}

	// 既存のユーザーかどうかチェックする
	exists, err := uc.userDomainService.Exists(ctx, user)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, nil
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.SaveWithTx(ctx, tx, user); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return &LoginUseCaseOutputDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.ID,
	}, nil
}
