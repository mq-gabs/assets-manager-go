package assets

import (
	"assets_manager/domain/entities/asset"
	"assets_manager/domain/entities/group"
	"assets_manager/domain/entities/user"
	e "assets_manager/utils/exception"
	"net/http"
)

var db_assets []*asset.Asset

func Save(a *asset.Asset) *e.Exception {
	db_assets = append(db_assets, a)

	return nil
}

func GetAssets() []*asset.Asset {
	return db_assets
}

func GetAssetById(id uint16) (*asset.Asset, *e.Exception) {
	for _, v := range db_assets {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, e.New("Asset not found", http.StatusNotFound)
}

type UpdateAssetDto struct {
	name  string
	group *group.Group
}

func UpdateAsset(id uint16, data UpdateAssetDto) *e.Exception {
	a, err := GetAssetById(id)

	if err != nil {
		return err
	}

	if data.name != "" {
		err1 := asset.UpdateName(a, data.name)

		if err1 != nil {
			return e.New(err1.Error(), http.StatusBadRequest)
		}
	}

	if data.group != nil {
		err2 := asset.UpdateGroup(a, data.group)

		if err2 != nil {
			return e.New(err2.Error(), http.StatusBadRequest)
		}
	}

	var new_db_assets []*asset.Asset

	for _, v := range db_assets {
		if v.ID == id {
			new_db_assets = append(new_db_assets, a)
		} else {
			new_db_assets = append(new_db_assets, v)
		}
	}

	db_assets = new_db_assets

	return nil
}

func DeleteAsset(id uint16) *e.Exception {
	_, err := GetAssetById(id)

	if err != nil {
		return err
	}

	var new_db_assets []*asset.Asset

	for _, v := range db_assets {
		if v.ID != id {
			new_db_assets = append(new_db_assets, v)
		}
	}

	db_assets = new_db_assets

	return nil
}

func UpdateStatus(id uint16, status asset.Status) *e.Exception {
	a, err := GetAssetById(id)

	if err != nil {
		return err
	}

	errr := asset.UpdateStatus(a, status)

	if errr != nil {
		return e.New(errr.Error(), http.StatusBadRequest)
	}

	return nil
}

func UpdateCurrentUser(id uint16, user *user.User) *e.Exception {
	a, err := GetAssetById(id)

	if err != nil {
		return err
	}

	errr := asset.UpdateCurrentUser(a, user)

	if errr != nil {
		return err
	}

	return nil
}
