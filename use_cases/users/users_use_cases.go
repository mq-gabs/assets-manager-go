package users_use_cases

import (
	"assets_manager/domain/entities/user"
	repo "assets_manager/infra/repository/memory/users"
	"net/http"

	e "assets_manager/utils/exception"
	q "assets_manager/utils/query"
)

type CreateUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(data *CreateUserDto) *e.Exception {
	u, err := user.New(data.Name, data.Email)

	if err != nil {
		return e.New(err.Error(), http.StatusBadRequest)
	}

	if err := repo.Save(u); err != nil {
		return err
	}

	return nil
}

func FindUsers(query *q.IQuery) []*user.User {

	us := repo.GetUsers(query)

	return us
}

func FindUserById(id uint16) (*user.User, *e.Exception) {
	u, err := repo.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return u, nil
}

type UpdateUserDto struct {
	Name  string
	Email string
}

func UpdateUser(id uint16, data *UpdateUserDto) *e.Exception {
	u, err := FindUserById(id)

	if err != nil {
		return err
	}

	if data.Name != "" {
		if err := user.UpdateName(u, data.Name); err != nil {
			return e.New(err.Error(), http.StatusBadRequest)
		}
	}

	if data.Email != "" {
		if err := user.UpdateEmail(u, data.Email); err != nil {
			return e.New(err.Error(), http.StatusBadRequest)
		}
	}

	if err := repo.UpdateUser(id, u); err != nil {
		return err
	}

	return nil
}

func DeleteUser(id uint16) *e.Exception {
	if err := repo.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
