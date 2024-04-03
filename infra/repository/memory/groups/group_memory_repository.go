package groups

import (
	"assets_manager/domain/entities/group"
	e "assets_manager/utils/exception"
	"assets_manager/utils/query"
)

var db_groups []*group.Group

func Save(g *group.Group) *e.Exception {
	db_groups = append(db_groups, g)

	return nil
}

func GetGroups(query *query.IQuery) []*group.Group {
	return db_groups
}

func GetGroupById(id uint16) (*group.Group, *e.Exception) {
	for _, v := range db_groups {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, e.New("Group not found", 404)
}

func UpdateGroup(id uint16, g *group.Group) *e.Exception {
	var new_db_groups []*group.Group

	for _, v := range db_groups {
		if v.ID == id {
			new_db_groups = append(new_db_groups, g)
		} else {
			new_db_groups = append(new_db_groups, v)
		}
	}

	db_groups = new_db_groups

	return nil
}

func DeleteGroup(id uint16) *e.Exception {
	var newGroups []*group.Group

	for _, v := range db_groups {
		if v.ID != id {
			newGroups = append(newGroups, v)
		}
	}

	db_groups = newGroups

	return nil
}
