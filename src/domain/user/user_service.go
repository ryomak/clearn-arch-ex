package user

import (
	"github.com/iancoleman/strcase"
)

func ToSnakeName(user *User) string {
	return strcase.ToSnake(user.Name)
}
