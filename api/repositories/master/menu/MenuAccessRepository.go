package menurepo

import (
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"

	"gorm.io/gorm"
)

type MenuAccessRepository interface {
	IsUserHaveAccess(*gorm.DB, menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse)
	IsByRoleDuplicate(*gorm.DB, menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse)
	GetByRoleID(*gorm.DB, int) ([]*menupayloads.GetMenuByRoleIDResponse, *exceptions.BaseErrorResponse)
	Create(*gorm.DB, menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse)
	Get(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
	Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
}
