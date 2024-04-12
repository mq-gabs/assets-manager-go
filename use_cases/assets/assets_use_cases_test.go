package assets_use_cases

import (
	"assets_manager/domain/entities/asset"
	groups_use_cases "assets_manager/use_cases/groups"
	users_use_cases "assets_manager/use_cases/users"
	"assets_manager/utils/query"
	"testing"
)

var aux_asset *asset.Asset

func TestCreateAsset(t *testing.T) {
	groupData := groups_use_cases.CreateGroupDto{
		Name: "Notebook",
	}

	g, err := groups_use_cases.CreateGroup(&groupData)

	if err != nil {
		t.Fatalf("Cannot create group to create asset. Error: %v", err)
	}

	assetData := CreateAssetDto{
		Name:    "Notebook Lenovo",
		GroupId: g.ID,
	}

	a, err2 := CreateAsset(&assetData)

	if err2 != nil {
		t.Errorf("Some error occurred: %v", err2)
	}

	aux_asset = a
}

func TestFindAssets(t *testing.T) {
	gs := FindAssets(&query.IQuery{})

	if len(gs) == 0 {
		t.Errorf("Assets list is empty but it should not.")
	}
}

func TestFindAssetById(t *testing.T) {
	a, err := FindAssetById(aux_asset.ID)

	if err != nil {
		t.Errorf("Some error occurred: %v", err)
	}

	if a.ID != aux_asset.ID {
		t.Errorf("IDs do not match. ID(%q, %q)", a.ID, aux_asset.ID)
	}
}

func TestUpdateAsset(t *testing.T) {
	newGroupData := groups_use_cases.CreateGroupDto{
		Name: "Eletrônicos",
	}

	newg, err := groups_use_cases.CreateGroup(&newGroupData)

	if err != nil {
		t.Fatalf("Cannot create new group to update asset. Error: %v", err)
	}

	updatedData := UpdateAssetDto{
		Name:    "Notebook Lenovo",
		Status:  asset.INACTIVE,
		GroupId: newg.ID,
	}

	a, err2 := UpdateAsset(aux_asset.ID, &updatedData)

	if err2 != nil {
		t.Errorf("Some error occurred: %v", err2)
	}

	if a.Name != updatedData.Name {
		t.Errorf("Names do not match. Names(%q, %q)", a.Name, updatedData.Name)
	}
}

func TestChangeStatus(t *testing.T) {
	newData := ChangeStatusDto{
		Status: asset.RESERVED,
	}

	a, err := ChangeStatus(aux_asset.ID, &newData)

	if err != nil {
		t.Fatalf("Some error occured: %v", err)
	}

	if a.Status != asset.RESERVED {
		t.Errorf("Status do not match. Status(%q, %q)", a.Status, asset.RESERVED)
	}
}

func TestSetCurrentUser(t *testing.T) {
	userData := users_use_cases.CreateUserDto{
		Name:  "John",
		Email: "john@gmail.com",
	}

	u, err := users_use_cases.CreateUser(&userData)

	if err != nil {
		t.Fatalf("Cannot create user to set new asset user. Error: %v", err)
	}

	newData := SetCurrentUserDto{
		UserId: u.ID,
	}

	a, err2 := SetCurrentUser(aux_asset.ID, &newData)

	if err2 != nil {
		t.Errorf("Some error occured: %v", err2)
	}

	if a.CurrentUser.ID != u.ID {
		t.Errorf("IDs do not match. IDs(%q, %q)", a.CurrentUser.ID, u.ID)
	}
}

func TestDeleteAsset(t *testing.T) {
	err := DeleteAsset(aux_asset.ID)

	if err != nil {
		t.Errorf("Some error occured: %v", err)
	}

	a, err := FindAssetById(aux_asset.ID)

	if err == nil {
		t.Errorf("Asset found but it should not. IDs(%q, %q)", a.ID, aux_asset.ID)
	}
}
