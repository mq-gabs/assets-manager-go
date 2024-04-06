package asset

import (
	"assets_manager/domain/entities/group"
	"assets_manager/domain/entities/user"
	"errors"
	"math/rand"
)

type Status string

const (
	ACTIVE   Status = "ACTIVE"
	BUSY     Status = "BUSY"
	REPAIR   Status = "REPAIR"
	RESERVED Status = "RESERVED"
	INACTIVE Status = "INACTIVE"
)

type Asset struct {
	ID          uint16       `json:"id"`
	Name        string       `json:"name"`
	Status      Status       `json:"status"`
	Group       *group.Group `json:"group"`
	CurrentUser *user.User   `json:"user"`
}

func New(name string, group *group.Group) (*Asset, error) {
	if err := checkName(name); err != nil {
		return nil, err
	}

	if err := checkGroup(group); err != nil {
		return nil, err
	}

	asset := Asset{
		ID:     uint16(rand.Uint32()),
		Name:   name,
		Group:  group,
		Status: ACTIVE,
	}

	return &asset, nil
}

func UpdateName(asset *Asset, newName string) error {
	if err := checkName(newName); err != nil {
		return err
	}

	asset.Name = newName

	return nil
}

func UpdateStatus(asset *Asset, newStatus Status) error {
	if newStatus != ACTIVE && newStatus != BUSY && newStatus != INACTIVE && newStatus != REPAIR && newStatus != RESERVED {
		return errors.New("this status is not available")
	}

	asset.Status = newStatus

	return nil
}

func UpdateGroup(asset *Asset, group *group.Group) error {
	if err := checkGroup(group); err != nil {
		return err
	}

	asset.Group = group

	return nil
}

func UpdateCurrentUser(asset *Asset, newUser *user.User) error {
	if err := checkUser(newUser); err != nil {
		return err
	}

	asset.CurrentUser = newUser

	return nil
}

func checkName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func checkGroup(group *group.Group) error {
	if group.ID == 0 {
		return errors.New("invalid group id")
	}

	return nil
}

func checkUser(user *user.User) error {
	if user.ID == 0 {
		return errors.New("invalid user id")
	}

	return nil
}
