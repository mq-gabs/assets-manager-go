package user

import "assets_manager/domain/entities/user"

type UpdateUserDto struct {
	name  string
	email string
}

type UserRepository interface {
	Save(user.User) error
	GetUsers() []user.User
	GetUserById() (user.User, error)
	UpdateUser(int64, UpdateUserDto) error
	DeleteUser(int64) error
}
