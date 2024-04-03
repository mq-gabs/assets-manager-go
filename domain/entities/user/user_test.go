package user

import "testing"

var user *User

func TestNew(t *testing.T) {
	name := "John"
	email := "john@mail.com"

	newUser, err := New(name, email)

	if err != nil {
		t.Fatalf("Some error occurred: %v", err)
	}

	user = newUser

	if newUser.Name != name || newUser.Email != email {
		t.Errorf("Name or email did not match. Name: (%q, %q), Email: (%q, %q)", name, newUser.Name, email, newUser.Email)
	}
}

func TestUpdateName(t *testing.T) {
	newName := "Mike"

	err := UpdateName(user, newName)

	if err != nil {
		t.Errorf("Some error ocurred: %v", err)
	}

	if user.Name != newName {
		t.Errorf("Name was not updated. Name(%q,  %q)", newName, user.Name)
	}
}

func TestUpdateEmail(t *testing.T) {
	newEmail := "mike@mail.com"

	err := UpdateEmail(user, newEmail)

	if err != nil {
		t.Errorf("Some error ocurred: %v", err)
	}

	if newEmail != user.Email {
		t.Errorf("Email was not updated. Email(%q, %q)", newEmail, user.Email)
	}
}
