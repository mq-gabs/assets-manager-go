package assets

import (
	"assets_manager/domain/entities/asset"
	"assets_manager/domain/entities/group"
	"testing"
)

var aux_asset *asset.Asset

func TestSave(t *testing.T) {
	g, err1 := group.New("Notebook")

	if err1 != nil {
		t.Fatalf("Cannot create group")
	}

	a, err2 := asset.New("Notebook Lenovo", g)

	if err2 != nil {
		t.Fatalf("Cannot create asset")
	}

	aux_asset = a

	err3 := Save(a)

	if err3 != nil {
		t.Errorf("Error when saving asset: %v", err3)
	}
}

func TestGetAsset(t *testing.T) {
	as := GetAssets(nil)

	if len(as) == 0 {
		t.Errorf("Array is empry")
	}
}

func TestGetAssetById(t *testing.T) {
	_, err := GetAssetById(aux_asset.ID)

	if err != nil {
		t.Errorf("Error when getting asset by id: %v", err)
	}
}

func TestUpdateAsset(t *testing.T) {
	asset.UpdateName(aux_asset, "Notebook Lenovo")

	if err := UpdateAsset(aux_asset.ID, aux_asset); err != nil {
		t.Errorf("Some error occurred: %v", err)
	}
}

func TestDeleteAsset(t *testing.T) {
	err := DeleteAsset(aux_asset.ID)

	if err != nil {
		t.Errorf("Error when deleting asset: %v", err)
	}
}
