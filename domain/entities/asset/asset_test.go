package asset

import (
	"assets_manager/domain/entities/group"
	"assets_manager/domain/entities/user"
	"testing"
)

var asset *Asset

func TestNew(t *testing.T) {
	name := "Notebook"
	group, _ := group.New("MyGroup")

	newAsset, err := New(name, group)

	if err != nil {
		t.Fatalf("Some error occurred: %v", err)
	}

	asset = newAsset

	if newAsset.Name != name {
		t.Errorf("Name did not match. Name(%q, %q)", name, newAsset.Name)
	}

	if newAsset.Group != group {
		t.Errorf("Group did not match")
	}
}

func TestUpdateName(t *testing.T) {
	newName := "Notebook Lenovo"

	err := UpdateName(asset, newName)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if newName != asset.Name {
		t.Errorf("Name did not match. Name(%q, %q)", newName, asset.Name)
	}
}

func TestUpdateGroup(t *testing.T) {
	newGroup, _ := group.New("NewGroup")

	err := UpdateGroup(asset, newGroup)

	if err != nil {
		t.Errorf("Some error occured: %v", err)
	}

	if asset.Group != newGroup {
		t.Errorf("Gruop did not match")
	}
}

func TestUpdateCurrentuser(t *testing.T) {
	newUser, _ := user.New("User", "user@gmail.com")

	err := UpdateCurrentUser(asset, newUser)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if newUser != asset.CurrentUser {
		t.Errorf("User did not match")
	}
}

func TestUpdateStatus(t *testing.T) {
	UpdateStatus(asset, INACTIVE)

	if asset.Status != INACTIVE {
		t.Errorf("Status did not match. Status(%q, %q)", INACTIVE, asset.Status)
	}
}
