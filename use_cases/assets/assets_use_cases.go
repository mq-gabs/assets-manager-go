package assets_use_cases

import (
	"assets_manager/domain/entities/asset"
	repo "assets_manager/infra/repository/memory/assets"
	"assets_manager/infra/repository/memory/users"
	groups_use_cases "assets_manager/use_cases/groups"
	e "assets_manager/utils/exception"
	"assets_manager/utils/query"
	"net/http"
)

type CreateAssetDto struct {
	Name    string `json:"name"`
	GroupId uint16 `json:"groupId"`
}

func CreateAsset(data *CreateAssetDto) *e.Exception {
	g, err := groups_use_cases.FindGroupById(data.GroupId)

	if err != nil {
		return err
	}

	a, err2 := asset.New(data.Name, g)

	if err2 != nil {
		return e.New(err2.Error(), http.StatusBadRequest)
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
	Name    string       `json:"name"`
	Status  asset.Status `json:"status"`
	GroupId uint16       `json:"groupId"`
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

	if data.GroupId != 0 {
		g, err2 := groups_use_cases.FindGroupById(data.GroupId)

		if err2 != nil {
			return err2
		}

		if err := asset.UpdateGroup(a, g); err != nil {
			return e.New(err.Error(), http.StatusBadRequest)
		}
	}

	if data.Status != "" {
		if err := asset.UpdateStatus(a, data.Status); err != nil {
			return e.New(err.Error(), http.StatusBadRequest)
		}
	}

	return repo.UpdateAsset(id, a)
}

func DeleteAsset(id uint16) *e.Exception {
	return repo.DeleteAsset(id)
}

type ChangeStatusDto struct {
	Status asset.Status `json:"status"`
}

func ChangeStatus(id uint16, data *ChangeStatusDto) *e.Exception {
	a, err := FindAssetById(id)

	if err != nil {
		return err
	}

	if err := asset.UpdateStatus(a, data.Status); err != nil {
		return e.New(err.Error(), http.StatusBadRequest)
	}

	return repo.UpdateAsset(id, a)
}

type SetCurrentUserDto struct {
	UserId uint16 `json:"userId"`
}

func SetCurrentUser(id uint16, data *SetCurrentUserDto) *e.Exception {
	user, err := users.GetUserById(data.UserId)

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

	asset.UpdateStatus(a, asset.BUSY)

	return repo.UpdateAsset(id, a)
}
