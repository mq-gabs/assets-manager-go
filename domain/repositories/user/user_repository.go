package user

import (
	"assets_manager/domain/entities/user"
	"assets_manager/utils/query"
)

type UpdateUserDto struct {
	Name  string
	Email string
}

type UserRepository interface {
	Save(user.User) error
	GetUsers(query query.IQuery) []user.User
	GetUserById() (user.User, error)
	UpdateUser(int64, UpdateUserDto) error
	DeleteUser(int64) error
}
