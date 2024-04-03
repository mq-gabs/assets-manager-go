package groups

import (
	"assets_manager/domain/entities/group"
	"testing"
)

var g *group.Group

func TestSave(t *testing.T) {
	newGroup, err := group.New("Notebook")

	if err != nil {
		t.Errorf("Error creating group: %v", err)
	}

	g = newGroup

	errr := Save(newGroup)

	if errr != nil {
		t.Errorf("Some error occurred: %v", errr)
	}
}

func TestGetGroups(t *testing.T) {
	groups := GetGroups()

	if len(groups) == 0 {
		t.Errorf("Repo is empty.")
	}
}

func TestGetGroupById(t *testing.T) {
	group, err := GetGroupById(g.ID)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if group.ID != g.ID {
		t.Errorf("IDs did not match. IDS(%q, %q)", g.ID, group.ID)
	}
}

func TestUpdateGroup(t *testing.T) {
	udpateDate := UpdateGroupDto{
		name: "Device",
	}

	err := UpdateGroup(g.ID, &udpateDate)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	group, errr := GetGroupById(g.ID)

	if errr != nil {
		t.Errorf("Error to get group: %v", errr)
	}

	if group.Name != udpateDate.name {
		t.Errorf("Names did not match. Name(%q, %q)", udpateDate.name, group.Name)
	}
}

func TestDeleteGroup(t *testing.T) {
	err := DeleteGroup(g.ID)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	_, errr := GetGroupById(g.ID)

	if errr == nil {
		t.Errorf("Group was not deleted")
	}
}
