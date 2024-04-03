package users

import (
	"assets_manager/domain/entities/user"
	e "assets_manager/utils/exception"
	"assets_manager/utils/query"
	"net/http"
)

var db_users []*user.User

func Save(user *user.User) *e.Exception {
	db_users = append(db_users, user)

	return nil
}

func GetUsers(query *query.IQuery) []*user.User {
	return db_users
}

func GetUserById(id uint16) (*user.User, *e.Exception) {

	for _, v := range db_users {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, e.New("User Not Found", http.StatusNotFound)
}

func UpdateUser(id uint16, u *user.User) *e.Exception {
	var new_db_users []*user.User

	for _, v := range db_users {
		if v.ID == id {
			new_db_users = append(new_db_users, u)
		} else {
			new_db_users = append(new_db_users, v)
		}
	}

	db_users = new_db_users

	return nil
}

func DeleteUser(id uint16) *e.Exception {
	var new_db_users []*user.User
	found := false

	for _, v := range db_users {
		if v.ID != id {
			new_db_users = append(new_db_users, v)
		} else {
			found = true
		}
	}

	if found {
		db_users = new_db_users
		return nil
	}

	return e.New("User does not exists", http.StatusNotFound)
}
