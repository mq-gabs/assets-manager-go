package users_use_cases

import (
	"assets_manager/domain/entities/user"
	"testing"
)

var aux_user *user.User

func TestCreateUser(t *testing.T) {
	data := CreateUserDto{
		Name:  "John",
		Email: "john@gmail.com",
	}

	u, err := CreateUser(&data)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	aux_user = u
}

func TestFindUsers(t *testing.T) {
	u, err := FindUserById(aux_user.ID)

	if err != nil {
		t.Errorf("Some errors occurred: %v", err)
	}

	if u.ID != aux_user.ID {
		t.Errorf("Ids does not match. IDS(%q, %q)", u.ID, aux_user.ID)
	}
}

func TestUpdateUser(t *testing.T) {
	data := UpdateUserDto{
		Name:  "Mike",
		Email: "mike@gmail.com",
	}

	u, err := UpdateUser(aux_user.ID, &data)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if u.Name != data.Name {
		t.Errorf("Names does not match. Names(%q, %q)", data.Name, u.Name)
	}

	if u.Email != data.Email {
		t.Errorf("Emails does not match. Emails(%q, %q)", data.Email, u.Email)
	}
}

func TestDeleteUser(t *testing.T) {
	if err := DeleteUser(aux_user.ID); err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	u, err := FindUserById(aux_user.ID)

	if err == nil {
		t.Errorf("User found but it should not. IDs(%q, %q)", aux_user.ID, u.ID)
	}
}
