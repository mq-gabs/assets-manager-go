package group

import "testing"

var group *Group

func TestNew(t *testing.T) {
	name := "Notebook"

	newGroup, err := New(name)

	if err != nil {
		t.Fatalf("Some error occurred: %v", err)
	}

	group = newGroup

	if newGroup.Name != name {
		t.Errorf("Name did not match: Name(%q, %q)", name, newGroup.Name)
	}
}

func TestUpdateName(t *testing.T) {
	newName := "Eletronic"

	err := UpdateName(group, newName)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if group.Name != newName {
		t.Errorf("Name did not match. Name(%q, %q)", newName, group.Name)
	}
}
