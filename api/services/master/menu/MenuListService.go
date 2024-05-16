package menuservices

import (
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"
)

type MenuListService interface {
	Create(int, []menupayloads.CreateMenuListRequest) (bool, *exceptions.BaseErrorResponse)
	GetDropDown() ([]menupayloads.DropDownListForMenuListResponse, *exceptions.BaseErrorResponse)
}
