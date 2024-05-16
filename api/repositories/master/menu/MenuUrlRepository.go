package menurepo

import (
	"user-services/api/exceptions"
	menupayloads "user-services/api/payloads/master/menu"

	"gorm.io/gorm"
)

type MenuUrlRepository interface {
	Create(*gorm.DB, []menupayloads.CreateMenuUrlRequest) (bool, *exceptions.BaseErrorResponse)
	GetByName(*gorm.DB, []string) (menupayloads.GetMenuUrlByName, *exceptions.BaseErrorResponse)
	GetByRoleID(*gorm.DB, int) ([]string, *exceptions.BaseErrorResponse)
}
