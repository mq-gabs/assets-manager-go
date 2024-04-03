package groups

import (
	"assets_manager/domain/entities/group"
	e "assets_manager/utils/exception"
)

var groups []*group.Group

func Save(g *group.Group) *e.Exception {
	groups = append(groups, g)

	return nil
}

func GetGroups() []*group.Group {
	return groups
}

func GetGroupById(id uint16) (*group.Group, *e.Exception) {
	for _, v := range groups {
		if v.ID == id {
			return v, nil
		}
	}

	return &group.Group{}, e.New("Group not found", 404)
}

type UpdateGroupDto struct {
	name string
}

func UpdateGroup(id uint16, data *UpdateGroupDto) *e.Exception {
	g, err := GetGroupById(id)

	if err != nil {
		return err
	}

	errr := group.UpdateName(g, data.name)

	if errr != nil {
		return e.New(errr.Error(), 400)
	}

	var newGroups []*group.Group

	for _, v := range groups {
		if v.ID == id {
			newGroups = append(newGroups, g)
		} else {
			newGroups = append(newGroups, v)
		}
	}

	groups = newGroups

	return nil
}

func DeleteGroup(id uint16) *e.Exception {
	var newGroups []*group.Group

	for _, v := range groups {
		if v.ID != id {
			newGroups = append(newGroups, v)
		}
	}

	groups = newGroups

	return nil
}
