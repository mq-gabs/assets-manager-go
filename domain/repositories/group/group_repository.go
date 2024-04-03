package group

import "assets_manager/domain/entities/group"

type UpdateGroupDto struct {
	name string
}

type GroupRepository interface {
	Save(group.Group) error
	GetGroups() []group.Group
	GetGroupById(uint16) (group.Group, error)
	UpdateGroup(uint16, UpdateGroupDto) error
	DeleteGroup(uint16) error
}
