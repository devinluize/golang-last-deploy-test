package usergroupservices

import (
	usergroupentities "user-services/api/entities/master/user-group"
	"user-services/api/exceptions"
	approverpayloads "user-services/api/payloads/master/approver"
	usergrouppayloads "user-services/api/payloads/master/user-group"
)

type UserGroupService interface {
	Create(usergrouppayloads.CreateUserGroupRequest) (bool, *exceptions.BaseErrorResponse)
	Update(int, usergrouppayloads.UpdateUserGroupRequest) (bool, *exceptions.BaseErrorResponse)
	Delete(int) (bool, *exceptions.BaseErrorResponse)
	GetAll(int) (usergroupentities.UserGroup, *exceptions.BaseErrorResponse)
	GetByUserGroupName([]string) ([]usergroupentities.UserGroup, *exceptions.BaseErrorResponse)
	GetLeaderForApproval(usergrouppayloads.GetLeaderForApprovalRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse)
}
