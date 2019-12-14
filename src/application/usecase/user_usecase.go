package usecase

import (
	"context"
	"map-friend/src/domain/user"
)

type IUserUseCase interface {
	GetUserByID(context.Context, uint) (*user.User, error)
}

type userUseCase struct {
	userRepo user.IUserRepository
}

func NewIUserUseCase(
	user user.IUserRepository,
) IUserUseCase {
	return &userUseCase{user}
}

func (u *userUseCase) GetUserByID(ctx context.Context, id uint) (*user.User, error) {
	fUser, err := u.userRepo.FindOneByID(ctx, id)
	if err != nil {
		return nil, err
	}
	fUser.Name = user.ToSnakeName(fUser)
	return fUser, nil
}
