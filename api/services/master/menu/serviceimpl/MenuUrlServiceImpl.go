package menuservicesimpl

import (
	"user-services/api/exceptions"
	"user-services/api/helper"
	menupayloads "user-services/api/payloads/master/menu"
	menurepo "user-services/api/repositories/master/menu"
	userrepo "user-services/api/repositories/user"
	menuservices "user-services/api/services/master/menu"

	"gorm.io/gorm"
)

type MenuUrlServiceImpl struct {
	MenuUrlRepository menurepo.MenuUrlRepository
	UserRepository    userrepo.UserRepository
	DB                *gorm.DB
}

func NewMenuUrlService(
	MenuUrlRepository menurepo.MenuUrlRepository,
	UserRepository userrepo.UserRepository,
	db *gorm.DB,
) menuservices.MenuUrlService {
	return &MenuUrlServiceImpl{
		DB:                db,
		MenuUrlRepository: MenuUrlRepository,
		UserRepository:    UserRepository,
	}
}

func (service *MenuUrlServiceImpl) Create(request []menupayloads.CreateMenuUrlRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	menuUrlPaths := []string{}
	for i := range request {
		menuUrlPaths = append(menuUrlPaths, request[i].MenuUrlPath)
	}

	_, err := service.MenuUrlRepository.GetByName(tx, menuUrlPaths)
	if err != nil {
		return false, err
	}

	create, err := service.MenuUrlRepository.Create(tx, request)

	if err != nil {
		return false, err
	}

	return create, nil
}
func (service *MenuUrlServiceImpl) GetByCompanyAndUser(companyID int, userID int) ([]string, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	roleID, err := service.UserRepository.GetRoleByCompanyAndUserID(tx, companyID, userID)

	if err != nil {
		return []string{}, err
	}

	get, err := service.MenuUrlRepository.GetByRoleID(tx, roleID)

	if err != nil {
		return get, err
	}

	return get, nil
}
func (service *MenuUrlServiceImpl) GetByName(menuUrlPath []string) (menupayloads.GetMenuUrlByName, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.MenuUrlRepository.GetByName(tx, menuUrlPath)

	if err != nil {
		return get, err
	}

	return get, nil
}
