package menurepoimpl

import (
	"net/http"
	menuentities "user-services/api/entities/master/menu"
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"
	menurepo "user-services/api/repositories/master/menu"

	"gorm.io/gorm"
)

type MenuListRepositoryImpl struct {
}

func NewMenuListRepository() menurepo.MenuListRepository {
	return &MenuListRepositoryImpl{}
}

func (*MenuListRepositoryImpl) Create(tx *gorm.DB, request []menupayloads.CreateMenuListRequest) (bool, *exceptions.BaseErrorResponse) {
	menuListEntities := []menuentities.MenuList{}

	for i := range request {
		menuListEntities = append(menuListEntities, menuentities.MenuList{
			// Title:     request[i].MenuTitle,
			MenuUrlID: request[i].MenuUrlID,
			ParentID:  request[i].ParentID,
			Image:     request[i].MenuImage,
		})
	}

	rows, err := tx.
		Create(&menuListEntities).
		Rows()

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer rows.Close()

	return true, nil
}

func (*MenuListRepositoryImpl) IsDuplicate(tx *gorm.DB, menuTitle []string) (bool, *exceptions.BaseErrorResponse) {
	var menu menuentities.Menu
	var exists bool

	rows, err := tx.Model(&menu).
		Select("count(id)").
		Where("title in (?)", menuTitle).
		Find(&exists).
		Rows()

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if exists {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}

	defer rows.Close()

	return true, nil
}

func (*MenuListRepositoryImpl) GetDropDown(tx *gorm.DB) ([]menupayloads.DropDownListForMenuListResponse, *exceptions.BaseErrorResponse) {
	var menu menuentities.Menu
	var response []menupayloads.DropDownListForMenuListResponse

	tempRows := tx.Model(&menu).
		Select(
			"id",
			"title",
		)
	rows, err := tempRows.
		Order("id asc").
		Scan(&response).
		Rows()

	if err != nil {
		return response, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer rows.Close()

	return response, nil
}
