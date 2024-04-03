package group

import (
	"errors"
	"math/rand"
)

type Group struct {
	ID   uint16 `json:"id"`
	Name string `json:"name"`
}

func New(name string) (*Group, error) {
	if err := checkName(name); err != nil {
		return nil, err
	}

	group := Group{
		ID:   uint16(rand.Uint32()),
		Name: name,
	}

	return &group, nil
}

func UpdateName(group *Group, newName string) error {
	if err := checkName(newName); err != nil {
		return err
	}

	group.Name = newName

	return nil
}

func checkName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}
