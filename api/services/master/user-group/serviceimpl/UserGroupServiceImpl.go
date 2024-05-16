package usergroupservicesimpl

import (
	usergroupentities "user-services/api/entities/master/user-group"
	"user-services/api/exceptions"
	"user-services/api/helper"
	approverpayloads "user-services/api/payloads/master/approver"
	usergrouppayloads "user-services/api/payloads/master/user-group"
	usergrouprepo "user-services/api/repositories/master/user-group"
	usergroupservices "user-services/api/services/master/user-group"

	"gorm.io/gorm"
)

type UserGroupServiceImpl struct {
	UserGroupRepository usergrouprepo.UserGroupRepository
	DB                  *gorm.DB
}

// Create implements usergroupservices.UserGroupService.
func (service *UserGroupServiceImpl) Create(request usergrouppayloads.CreateUserGroupRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	create, err := service.UserGroupRepository.Create(tx, request)

	if err != nil {
		return false, err
	}

	return create, nil
}

// Update implements usergroupservices.UserGroupService.
func (service *UserGroupServiceImpl) Update(userGroupID int, request usergrouppayloads.UpdateUserGroupRequest) (bool, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// Delete implements usergroupservices.UserGroupService.
func (service *UserGroupServiceImpl) Delete(userGroupID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserGroupRepository.Get(tx, userGroupID)
	if err != nil {
		return get, err
	}
	delete, err := service.UserGroupRepository.Delete(tx, userGroupID)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

// GetAll implements usergroupservices.UserGroupService.
func (service *UserGroupServiceImpl) GetAll(int) (usergroupentities.UserGroup, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// GetByUserGroupName implements usergroupservices.UserGroupService.
func (service *UserGroupServiceImpl) GetByUserGroupName([]string) ([]usergroupentities.UserGroup, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// GetLeaderForApproval implements usergroupservices.UserGroupService.
func (service *UserGroupServiceImpl) GetLeaderForApproval(userGroupReq usergrouppayloads.GetLeaderForApprovalRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserGroupRepository.GetLeaderForApproval(tx, userGroupReq)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetUserGroupByName implements usergroupservices.UserGroupService.
func (service *UserGroupServiceImpl) GetUserGroupByName([]string) ([]usergroupentities.UserGroup, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

func NewUserGroupService(userGroupRepository usergrouprepo.UserGroupRepository, db *gorm.DB) usergroupservices.UserGroupService {
	return &UserGroupServiceImpl{
		UserGroupRepository: userGroupRepository,
		DB:                  db,
	}
}
