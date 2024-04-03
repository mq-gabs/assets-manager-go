package users

import (
	"assets_manager/domain/entities/user"
	e "assets_manager/utils/exception"
	"net/http"
)

var db_users []*user.User

func Save(user *user.User) *e.Exception {
	db_users = append(db_users, user)

	return nil
}

func GetUsers() []*user.User {
	return db_users
}

func GetUserById(id uint16) (*user.User, *e.Exception) {

	for _, v := range db_users {
		if v.ID == id {
			return v, nil
		}
	}

	return &user.User{}, e.New("User Not Found", http.StatusNotFound)
}

type UpdateUserType struct {
	name  string
	email string
}

func UpdateUser(id uint16, data UpdateUserType) (*user.User, *e.Exception) {

	if data.email == "" && data.name == "" {
		return &user.User{}, e.New("No data to update", http.StatusBadRequest)
	}

	foundUser, err := GetUserById(id)

	if err != nil {
		return &user.User{}, e.New("Cannot find user with this id", http.StatusNotFound)
	}

	if data.name != "" {
		err := user.UpdateName(foundUser, data.name)

		if err != nil {
			return &user.User{}, e.New(err.Error(), http.StatusBadRequest)
		}
	}

	if data.email != "" {
		err := user.UpdateEmail(foundUser, data.email)

		if err != nil {
			return &user.User{}, e.New(err.Error(), http.StatusBadRequest)
		}
	}

	var new_db_users []*user.User

	for _, vv := range db_users {
		if vv.ID == id {
			new_db_users = append(new_db_users, foundUser)
		} else {
			new_db_users = append(new_db_users, vv)
		}
	}

	db_users = new_db_users

	return foundUser, nil

}

func DeleteUser(id uint16) *e.Exception {
	for _, v := range db_users {
		if v.ID == id {
			var new_db_users []*user.User

			for _, vv := range db_users {
				if vv.ID != id {
					new_db_users = append(new_db_users, vv)
				}
			}

			db_users = new_db_users

			return nil
		}
	}

	return e.New("Cannot find user with this id", http.StatusNotFound)
}
