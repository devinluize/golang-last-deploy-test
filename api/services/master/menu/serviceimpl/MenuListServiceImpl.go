package menuservicesimpl

import (
	"net/http"
	"user-services/api/exceptions"
	"user-services/api/helper"
	menupayloads "user-services/api/payloads/master/menu"
	menurepo "user-services/api/repositories/master/menu"
	menuservices "user-services/api/services/master/menu"

	"gorm.io/gorm"
)

type MenuListServiceImpl struct {
	MenuListRepository menurepo.MenuListRepository
	DB                 *gorm.DB
}

func NewMenuListService(
	MenuListRepository menurepo.MenuListRepository,
	db *gorm.DB,
) menuservices.MenuListService {
	return &MenuListServiceImpl{
		MenuListRepository: MenuListRepository,
		DB:                 db,
	}
}

func (service *MenuListServiceImpl) Create(role int, request []menupayloads.CreateMenuListRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	menuTitles := []string{}
	for i := range request {
		menuTitles = append(menuTitles, request[i].MenuTitle)
	}
	if role != 1 {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusUnauthorized,
		}
	}

	checkDuplicates, err := service.MenuListRepository.IsDuplicate(tx, menuTitles)

	if err != nil {
		return checkDuplicates, err
	}

	create, err := service.MenuListRepository.Create(tx, request)

	if err != nil {
		return create, err
	}

	return create, nil
}

func (service *MenuListServiceImpl) GetDropDown() ([]menupayloads.DropDownListForMenuListResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.MenuListRepository.GetDropDown(tx)

	if err != nil {
		return get, err
	}

	return get, nil
}
