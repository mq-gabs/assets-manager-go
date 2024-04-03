package assets

import (
	"assets_manager/domain/entities/asset"
	"assets_manager/domain/entities/group"
	repo "assets_manager/infra/repository/memory/assets"
	"assets_manager/infra/repository/memory/users"
	e "assets_manager/utils/exception"
	"assets_manager/utils/query"
	"net/http"
)

func CreateAsset(name string, g *group.Group) *e.Exception {
	a, err := asset.New(name, g)

	if err != nil {
		return e.New(err.Error(), http.StatusBadRequest)
	}

	return repo.Save(a)
}

func FindAssets(query *query.IQuery) []*asset.Asset {
	return repo.GetAssets(query)
}

func FindAssetById(id uint16) (*asset.Asset, *e.Exception) {
	a, err := repo.GetAssetById(id)

	if err != nil {
		return nil, err
	}

	return a, nil
}

type UpdateAssetDto struct {
	Name   string
	Status asset.Status
	Group  *group.Group
}

func UpdateAsset(id uint16, data *UpdateAssetDto) *e.Exception {
	a, err := FindAssetById(id)

	if err != nil {
		return err
	}

	if data.Name != "" && data.Name != a.Name {
		if err := asset.UpdateName(a, data.Name); err != nil {
			return e.New(err.Error(), http.StatusBadRequest)
		}
	}

	if data.Group != nil {
		if err := asset.UpdateGroup(a, data.Group); err != nil {
			return e.New(err.Error(), http.StatusBadRequest)
		}
	}

	if err := asset.UpdateStatus(a, data.Status); err != nil {
		return e.New(err.Error(), http.StatusBadRequest)
	}

	return repo.UpdateAsset(id, a)
}

func DeleteAsset(id uint16) *e.Exception {
	return repo.DeleteAsset(id)
}

func UpdateStatus(id uint16, status asset.Status) *e.Exception {
	a, err := FindAssetById(id)

	if err != nil {
		return err
	}

	if err := asset.UpdateStatus(a, status); err != nil {
		return e.New(err.Error(), http.StatusBadRequest)
	}

	return repo.UpdateAsset(id, a)
}

func UpdateCurrentUser(id uint16, userId uint16) *e.Exception {
	user, err := users.GetUserById(id)

	if err != nil {
		return err
	}

	a, err := FindAssetById(id)

	if err != nil {
		return err
	}

	if err := asset.UpdateCurrentUser(a, user); err != nil {
		return e.New(err.Error(), http.StatusBadRequest)
	}

	return repo.UpdateAsset(id, a)
}
