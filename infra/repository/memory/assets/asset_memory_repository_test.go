package assets

import (
	"assets_manager/domain/entities/asset"
	"assets_manager/domain/entities/group"
	"assets_manager/domain/entities/user"
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
	as := GetAssets()

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
	newData := UpdateAssetDto{
		name: "Notebook Lenovo i5 12",
	}

	err := UpdateAsset(aux_asset.ID, newData)

	if err != nil {
		t.Errorf("Error when updating asset: %v", err)
	}

	if newData.name != aux_asset.Name {
		t.Errorf("Name did not match. Name(%q, %q)", newData.name, aux_asset.Name)
	}
}

func TestUpdateStatus(t *testing.T) {
	newStatus := asset.REPAIR

	err := UpdateStatus(aux_asset.ID, newStatus)

	if err != nil {
		t.Errorf("Error when updating asset status: %v", err)
	}

	if aux_asset.Status != newStatus {
		t.Errorf("Status did not match. Status(%q, %q)", newStatus, aux_asset.Status)
	}
}

func TestUpdateCurrentuser(t *testing.T) {
	newUser, err := user.New("John", "john@mail.com")

	if err != nil {
		t.Fatalf("Error when creating user for test asset: %v", err)
	}

	err2 := UpdateCurrentUser(aux_asset.ID, newUser)

	if err2 != nil {
		t.Errorf("Error when udpating current user: %v", err2)
	}

	if aux_asset.CurrentUser.ID != newUser.ID {
		t.Errorf("User was not updated")
	}
}

func TestDeleteAsset(t *testing.T) {
	err := DeleteAsset(aux_asset.ID)

	if err != nil {
		t.Errorf("Error when deleting asset: %v", err)
	}
}
