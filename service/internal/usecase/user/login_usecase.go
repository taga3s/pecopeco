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

type LoginUseCaseDto struct {
	ID    int
	Name  string
	Email string
}

func (uc *LoginUsecase) Run(ctx context.Context, dto LoginUseCaseDto) (*userDomain.User, error) {
	user, err := userDomain.NewUser(dto.ID, dto.Name, dto.Email)

	exists, err := uc.userDomainService.Exists(ctx, user)
	if err != nil {
		return nil, err
	}
	if exists {
		return user, nil
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.SaveWithTx(ctx, tx, user); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return user, nil
}
