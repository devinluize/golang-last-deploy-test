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

type MenuAccessServiceImpl struct {
	MenuAccessRepository menurepo.MenuAccessRepository
	UserRepository       userrepo.UserRepository
	DB                   *gorm.DB
}

func NewMenuAccessService(
	MenuAccessRepository menurepo.MenuAccessRepository,
	UserRepository userrepo.UserRepository,
	db *gorm.DB,
) menuservices.MenuAccessService {
	return &MenuAccessServiceImpl{
		DB:                   db,
		MenuAccessRepository: MenuAccessRepository,
		UserRepository:       UserRepository,
	}
}

func (service *MenuAccessServiceImpl) Create(menuAccessRequest menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := service.MenuAccessRepository.IsByRoleDuplicate(tx, menuAccessRequest)
	if err != nil {
		return false, err
	}

	create, err := service.MenuAccessRepository.Create(tx, menuAccessRequest)

	if err != nil {
		return false, err
	}

	return create, nil
}

func (service *MenuAccessServiceImpl) IsUserHaveAccess(menuAccessRequest menupayloads.CreateMenuAccessParamRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	check, err := service.MenuAccessRepository.IsUserHaveAccess(tx, menuAccessRequest)

	if err != nil {
		return check, err
	}

	return check, nil
}

func (service *MenuAccessServiceImpl) GetByCompanyAndUserID(companyID int, userID int) ([]*menupayloads.GetMenuByRoleIDResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	roleID, err := service.UserRepository.GetRoleByCompanyAndUserID(tx, companyID, userID)

	if err != nil {
		return nil, err
	}

	get, err := service.MenuAccessRepository.GetByRoleID(tx, roleID)

	if err != nil {
		return get, err
	}

	return get, nil
}

func (service *MenuAccessServiceImpl) Get(menuAccessID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.MenuAccessRepository.Get(tx, menuAccessID)

	if err != nil {
		return get, err
	}

	return get, nil
}

func (service *MenuAccessServiceImpl) Delete(menuAccessID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.MenuAccessRepository.Get(tx, menuAccessID)
	if err != nil {
		return get, err
	}
	delete, err := service.MenuAccessRepository.Delete(tx, menuAccessID)

	if err != nil {
		return delete, err
	}

	return delete, nil
}
