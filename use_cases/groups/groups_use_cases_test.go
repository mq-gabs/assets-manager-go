package groups_use_cases

import (
	"assets_manager/domain/entities/group"
	"assets_manager/utils/query"
	"testing"
)

var aux_group *group.Group

func TestCreateGroup(t *testing.T) {
	data := CreateGroupDto{
		Name: "Notebook",
	}

	g, err := CreateGroup(&data)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	aux_group = g
}

func TestFindGroups(t *testing.T) {
	gs := FindGroups(&query.IQuery{})

	if len(gs) == 0 {
		t.Error("Groups list is empty but it should not.")
	}
}

func TestFindGroupById(t *testing.T) {
	g, err := FindGroupById(aux_group.ID)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if g.ID != aux_group.ID {
		t.Errorf("IDs do not match. IDs(%q, %q)", g.ID, aux_group.ID)
	}
}

func TestUpdateGroup(t *testing.T) {
	data := UpdateGroupDto{
		Name: "notebooks",
	}

	g, err := UpdateGroup(aux_group.ID, &data)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if g.Name != data.Name {
		t.Errorf("Names do not mathc. Names(%q, %q)", g.Name, data.Name)
	}
}

func TestDeleteGroup(t *testing.T) {
	if err := DeleteGroup(aux_group.ID); err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	g, err := FindGroupById(aux_group.ID)

	if err == nil {
		t.Errorf("Group found but it should not. IDs(%q, %q)", g.ID, aux_group.ID)
	}
}
