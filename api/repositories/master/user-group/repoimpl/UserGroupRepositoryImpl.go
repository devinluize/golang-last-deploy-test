package usergrouprepoimpl

import (
	"fmt"
	"net/http"
	usergroupentities "user-services/api/entities/master/user-group"
	"user-services/api/exceptions"
	approverpayloads "user-services/api/payloads/master/approver"
	usergrouppayloads "user-services/api/payloads/master/user-group"
	usergrouprepo "user-services/api/repositories/master/user-group"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type UserGroupRepositoryImpl struct {
}

// CreateUserGroup implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) Create(tx *gorm.DB, userGroupReq usergrouppayloads.CreateUserGroupRequest) (bool, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// Get implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) Get(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// GetByName implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) GetByName(*gorm.DB, []string) (*[]usergroupentities.UserGroup, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// Update implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) Update(*gorm.DB, int, usergrouppayloads.UpdateUserGroupRequest) (bool, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// Delete implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// GetAllUserGroup implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) GetAll(tx *gorm.DB, companyID int) (*usergroupentities.UserGroup, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

// GetLeaderForApproval implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) GetLeaderForApproval(tx *gorm.DB, userGroupReq usergrouppayloads.GetLeaderForApprovalRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse) {
	var leaderResponse []approverpayloads.GetApproverForApprovalResponse
	var user usergroupentities.UserGroup

	err := tx.
		Select(
			"level",
		).
		Joins("ApproverDetails", tx.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver", tx.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver.ApprovalApprover",
			tx.Select(
				"1",
			)).
		Joins("ApproverDetails.Approver.ApprovalApprover.ApprovalLevel",
			tx.Select(
				"1",
			)).
		Model(&user).
		Scan(&leaderResponse).
		Error

	if err != nil {
		return &leaderResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if leaderResponse == nil {
		return nil, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        fmt.Errorf("%s %s", "User Group", utils.GetDataNotFound),
		}
	}

	return &leaderResponse, nil

}

// GetUserGroupByName implements usergrouprepo.UserGroupRepository.
func (*UserGroupRepositoryImpl) GetUserGroupByName(tx *gorm.DB, names []string) (*[]usergroupentities.UserGroup, *exceptions.BaseErrorResponse) {
	panic("unimplemented")
}

func NewUserGroupRepository() usergrouprepo.UserGroupRepository {
	return &UserGroupRepositoryImpl{}
}
