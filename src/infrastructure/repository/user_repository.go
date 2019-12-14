package repository

import (
	"context"
	"errors"

	"map-friend/src/domain/user"
)

func NewIUserRepository() user.IUserRepository {
	// mock
	users := map[uint]*user.User{
		1: &user.User{
			ID:   1,
			Name: "ryomak",
		},
		2: &user.User{
			ID:   2,
			Name: "test user",
		},
		3: &user.User{
			ID:   3,
			Name: "test user3",
		},
	}
	return &userRepository{users}
}

type userRepository struct {
	users map[uint]*user.User
}

func (r *userRepository) FindOneByID(ctx context.Context, id uint) (*user.User, error) {
	if val, ok := r.users[id]; !ok {
		return nil, errors.New("not found")
	} else {
		return val, nil
	}
}
