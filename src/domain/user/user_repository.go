package user

import "context"

type IUserRepository interface {
	FindOneByID(context.Context, uint) (*User, error)
}
