package menurepo

import (
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"

	"gorm.io/gorm"
)

type MenuListRepository interface {
	Create(*gorm.DB, []menupayloads.CreateMenuListRequest) (bool, *exceptions.BaseErrorResponse)
	IsDuplicate(*gorm.DB, []string) (bool, *exceptions.BaseErrorResponse)
	GetDropDown(*gorm.DB) ([]menupayloads.DropDownListForMenuListResponse, *exceptions.BaseErrorResponse)
}
