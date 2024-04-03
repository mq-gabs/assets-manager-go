package asset

import (
	"assets_manager/domain/entities/asset"
	"assets_manager/domain/entities/group"
	"assets_manager/domain/entities/user"
)

type UpdateAssetDataDto struct {
	name string
}

type AssetRepository interface {
	Save() error
	GetAssets() []asset.Asset
	GetAssetById(int64) asset.Asset
	UpdateAssetData(int64, UpdateAssetDataDto) error
	UpdateStatus(int64, asset.Status) error
	UpdateGroup(int64, group.Group) error
	UpdateCurrentUser(int64, user.User) error
}
