package menurepoimpl

import (
	"net/http"
	menuentities "user-services/api/entities/master/menu"
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"
	menurepo "user-services/api/repositories/master/menu"

	"gorm.io/gorm"
)

type MenuAccessRepositoryImpl struct {
}

func NewMenuAccessRepository() menurepo.MenuAccessRepository {
	return &MenuAccessRepositoryImpl{}
}

func (*MenuAccessRepositoryImpl) IsUserHaveAccess(tx *gorm.DB, menuAccessRequest menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse) {
	var menuAccess menuentities.MenuAccess

	err := tx.
		Preload("MenuListByAccess", menuAccessRequest.MenuListID).
		Find(&menuAccess, "role_id = ?", menuAccessRequest.RoleID).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if menuAccess.ID != 0 {
		return true, nil
	}

	return true, nil
}

func (*MenuAccessRepositoryImpl) IsByRoleDuplicate(tx *gorm.DB, menuAccessRequest menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse) {
	var menuAccess menuentities.MenuAccess

	err := tx.
		Preload("MenuListByAccess", menuAccessRequest.MenuListID).
		Find(&menuAccess, "role_id = ?", menuAccessRequest.RoleID).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if menuAccess.ID != 0 {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}

	return true, nil
}

func (*MenuAccessRepositoryImpl) GetByRoleID(tx *gorm.DB, roleID int) ([]*menupayloads.GetMenuByRoleIDResponse, *exceptions.BaseErrorResponse) {
	var menuAccess menuentities.MenuAccess
	var response []*menupayloads.GetMenuByRoleIDResponse

	err := tx.
		Preload("MenuListByAccess.MenuUrl").
		Preload("MenuListByAccess.MenuMaster").
		Preload("MenuListByAccess.MenuUrl.MenuList").
		Find(&menuAccess, "role_id = ?", roleID).
		Error

	if err != nil {
		return response, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	for i := range menuAccess.MenuListByAccess {
		response = append(response, &menupayloads.GetMenuByRoleIDResponse{
			MenuID:      menuAccess.MenuListByAccess[i].MenuUrl.MenuList.ID,
			MenuTitle:   menuAccess.MenuListByAccess[i].MenuMaster.Title,
			MenuUrlPath: menuAccess.MenuListByAccess[i].MenuUrl.Path,
			ParentID:    menuAccess.MenuListByAccess[i].ParentID,
		})
	}

	menuMap := make(map[int]*menupayloads.GetMenuByRoleIDResponse)
	var rootItems []*menupayloads.GetMenuByRoleIDResponse

	for _, item := range response {
		menuMap[int(item.MenuID)] = item

		if item.ParentID == 0 {
			rootItems = append(rootItems, item)
		} else {
			parent, exists := menuMap[int(item.ParentID)]
			if exists {
				parent.Children = append(parent.Children, item)
			} else {
				parentList, found := menuMap[int(item.ParentID)-1]
				if !found {
					parentList = &menupayloads.GetMenuByRoleIDResponse{Children: []*menupayloads.GetMenuByRoleIDResponse{}}
					menuMap[int(item.ParentID)-1] = parentList
				}
				parentList.Children = append(parentList.Children, item)
			}
		}
	}

	// Process the list of items waiting for their parents
	for _, waitingItems := range menuMap {
		if waitingItems.ParentID < 0 {
			parent, exists := menuMap[int(waitingItems.ParentID)-1]
			if exists {
				parent.Children = append(parent.Children, waitingItems.Children...)
			}
		}
	}

	// Filter out items that were waiting for their parents
	finalItems := make([]*menupayloads.GetMenuByRoleIDResponse, 0)
	for _, item := range rootItems {
		if item.ParentID >= 0 {
			finalItems = append(finalItems, item)
		}
	}
	return finalItems, nil
}

func (*MenuAccessRepositoryImpl) Create(tx *gorm.DB, menuAccessRequest menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse) {
	var menuAccessEntities []menuentities.MenuAccess

	for i := range menuAccessRequest.MenuListID {
		menuAccessEntities = append(menuAccessEntities, menuentities.MenuAccess{
			RoleID: menuAccessRequest.RoleID,
		})
		menuAccessEntities[i].MenuListByAccess = append(menuAccessEntities[i].MenuListByAccess, menuentities.MenuList{
			ID: menuAccessRequest.MenuListID[i],
		})
	}

	err := tx.
		Create(&menuAccessEntities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}
func (*MenuAccessRepositoryImpl) Get(tx *gorm.DB, menuAccessID int) (bool, *exceptions.BaseErrorResponse) {
	var menuAccessEntities menuentities.MenuAccess

	rows, err := tx.Model(&menuAccessEntities).
		Where(menuentities.MenuAccess{
			ID: menuAccessID,
		}).
		Scan(&menuAccessEntities).
		Rows()

	if menuAccessEntities.ID == 0 {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer rows.Close()

	return true, nil
}
func (*MenuAccessRepositoryImpl) Delete(tx *gorm.DB, menuAccessID int) (bool, *exceptions.BaseErrorResponse) {
	var menuAccess menuentities.MenuAccess

	err := tx.
		Model(&menuAccess).
		Delete(&menuAccess, menuAccessID).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

func (*MenuAccessRepositoryImpl) CheckMenuAccess(tx *gorm.DB, roleID int, menuUrlPath string) (string, *exceptions.BaseErrorResponse) {
	menuAccess := menuentities.MenuAccess{}

	err := tx.
		Model(&menuAccess).
		Preload("MenuListByAccess", "role_id = ?", roleID).
		Preload("MenuListByAccess.MenuUrl", "path = ?", menuUrlPath).
		Error

	if err != nil {

		return "", &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return menuAccess.MenuListByAccess[0].MenuUrl.Path, nil
}
