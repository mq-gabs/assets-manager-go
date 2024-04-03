package user

import (
	"errors"
	"math/rand"
)

type User struct {
	ID    uint16 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func New(name string, email string) (*User, error) {
	if err := checkName(name); err != nil {
		return &User{}, err
	}

	if err := checkEmail(email); err != nil {
		return &User{}, err
	}

	user := User{
		ID:    uint16(rand.Uint32()),
		Name:  name,
		Email: email,
	}

	return &user, nil
}

func UpdateEmail(user *User, newEmail string) error {
	if err := checkEmail(newEmail); err != nil {
		return err
	}

	user.Email = newEmail

	return nil
}

func UpdateName(user *User, newName string) error {
	if err := checkName(newName); err != nil {
		return err
	}
	user.Name = newName

	return nil
}

func checkName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func checkEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}
