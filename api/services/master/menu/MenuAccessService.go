package menuservices

import (
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"
)

type MenuAccessService interface {
	IsUserHaveAccess(menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse)
	GetByCompanyAndUserID(int, int) ([]*menupayloads.GetMenuByRoleIDResponse, *exceptions.BaseErrorResponse)
	Get(int) (bool, *exceptions.BaseErrorResponse)
	Create(menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse)
	Delete(int) (bool, *exceptions.BaseErrorResponse)
}
