package menuservices

import (
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"
)

type MenuUrlService interface {
	Create([]menupayloads.CreateMenuUrlRequest) (bool, *exceptions.BaseErrorResponse)
	GetByCompanyAndUser(int, int) ([]string, *exceptions.BaseErrorResponse)
	GetByName([]string) (menupayloads.GetMenuUrlByName, *exceptions.BaseErrorResponse)
}
