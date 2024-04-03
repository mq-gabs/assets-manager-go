package groups

import (
	"assets_manager/domain/entities/group"
	repo "assets_manager/infra/repository/memory/groups"
	e "assets_manager/utils/exception"
	"assets_manager/utils/query"
	"net/http"
)

func CreateGroup(name string) *e.Exception {
	g, err := group.New(name)

	if err != nil {
		return e.New(err.Error(), http.StatusBadRequest)
	}

	if err := repo.Save(g); err != nil {
		return err
	}

	return nil
}

func FindGroups(query *query.IQuery) []*group.Group {
	gs := repo.GetGroups(query)

	return gs
}

func FindGroupById(id uint16) (*group.Group, *e.Exception) {
	g, err := repo.GetGroupById(id)

	if err != nil {
		return nil, err
	}

	return g, nil
}

type UpdateGroupDto struct {
	Name string
}

func UpdateGroup(id uint16, data *UpdateGroupDto) *e.Exception {
	g, err := FindGroupById(id)

	if err != nil {
		return err
	}

	if data.Name != "" {
		if err := group.UpdateName(g, data.Name); err != nil {
			return e.New(err.Error(), http.StatusBadRequest)
		}
	}

	if err := repo.UpdateGroup(id, g); err != nil {
		return err
	}

	return nil
}

func DeleteGroup(id uint16) *e.Exception {
	if err := repo.DeleteGroup(id); err != nil {
		return err
	}

	return nil
}
