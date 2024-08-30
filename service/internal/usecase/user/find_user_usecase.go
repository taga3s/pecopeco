package user

import (
	"context"

	userDomain "github.com/taga3s/pecopeco-service/internal/domain/user"
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

func (uc *FindUserUseCase) Run(ctx context.Context, id string) (*FindUserUseCaseDto, error) {
	user, err := uc.userRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
