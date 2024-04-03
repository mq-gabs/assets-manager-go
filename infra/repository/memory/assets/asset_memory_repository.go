package assets

import (
	"assets_manager/domain/entities/asset"
	e "assets_manager/utils/exception"
	"assets_manager/utils/query"
	"net/http"
)

var db_assets []*asset.Asset

func Save(a *asset.Asset) *e.Exception {
	db_assets = append(db_assets, a)

	return nil
}

func GetAssets(query *query.IQuery) []*asset.Asset {
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

func UpdateAsset(id uint16, a *asset.Asset) *e.Exception {
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
