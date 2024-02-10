package user

import (
	"context"

	userDomain "github.com/Seiya-Tagami/pecopeco-service/internal/domain/user"
)

type FindUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserUseCase(
	userRepo userDomain.UserRepository,
) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

type FindUserUseCaseDto struct {
	ID    string
	Name  string
	Email string
}

func (uc *FindUserUseCase) Run(ctx context.Context, dto FindUserUseCaseDto) (*userDomain.User, error) {
	user, err := uc.userRepo.FindById(ctx, dto.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
