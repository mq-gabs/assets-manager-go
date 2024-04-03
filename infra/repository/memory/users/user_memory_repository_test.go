package users

import (
	"assets_manager/domain/entities/user"
	"testing"
)

var aux_user *user.User

func TestSave(t *testing.T) {
	newUser, err := user.New("John", "john@mail.com")

	aux_user = newUser

	if err != nil {
		t.Error("Failed to create user for testing")
	}

	errr := Save(newUser)

	if errr != nil {
		t.Errorf("Some error occured: %v", err)
	}
}

func TestGetUsers(t *testing.T) {
	users := GetUsers()

	if len(users) == 0 {
		t.Errorf("Array not empty")
	}

	if users[0].Name != "John" {
		t.Errorf("Name did not match. Name(%q, %q)", "John", users[0].Name)
	}
}

func TestGetUserById(t *testing.T) {
	user, err := GetUserById(aux_user.ID)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if user.Name != aux_user.Name {
		t.Errorf("Name did not match. Name(%q, %q)", aux_user.Name, user.Name)
	}
}

func TestUpdateUser(t *testing.T) {
	newName := "Mike"

	updatedUser, err := UpdateUser(aux_user.ID, UpdateUserType{name: newName})

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if updatedUser.Name != newName {
		t.Errorf("Name did not match. Name(%q, %q)", newName, updatedUser.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser(aux_user.ID)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	_, errr := GetUserById(aux_user.ID)

	if errr == nil {
		t.Errorf("Some error occurred: %v", err)
	}

}
