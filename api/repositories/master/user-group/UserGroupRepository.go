package usergrouprepo

import (
	usergroupentities "user-services/api/entities/master/user-group"
	"user-services/api/exceptions"
	approverpayloads "user-services/api/payloads/master/approver"
	usergrouppayloads "user-services/api/payloads/master/user-group"

	"gorm.io/gorm"
)

type UserGroupRepository interface {
	Create(*gorm.DB, usergrouppayloads.CreateUserGroupRequest) (bool, *exceptions.BaseErrorResponse)
	Get(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
	Update(*gorm.DB, int, usergrouppayloads.UpdateUserGroupRequest) (bool, *exceptions.BaseErrorResponse)
	Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
	GetAll(*gorm.DB, int) (*usergroupentities.UserGroup, *exceptions.BaseErrorResponse)
	GetByName(*gorm.DB, []string) (*[]usergroupentities.UserGroup, *exceptions.BaseErrorResponse)
	GetLeaderForApproval(*gorm.DB, usergrouppayloads.GetLeaderForApprovalRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse)
}
