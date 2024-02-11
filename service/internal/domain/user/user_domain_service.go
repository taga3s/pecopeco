package user

import (
	"context"
)

type userDomainService struct {
	userRepo UserRepository
}

func NewUserDomainService(userRepo UserRepository) UserDomainService {
	return &userDomainService{
		userRepo: userRepo,
	}
}

func (ds *userDomainService) Exists(ctx context.Context, user *User) (bool, error) {
	count, err := ds.userRepo.CountById(ctx, user.ID)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
