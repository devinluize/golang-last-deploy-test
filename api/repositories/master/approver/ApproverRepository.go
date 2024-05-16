package approverrepo

import (
	approverentities "user-services/api/entities/master/approver"
	"user-services/api/exceptions"
	approverpayloads "user-services/api/payloads/master/approver"

	"gorm.io/gorm"
)

type ApproverRepository interface {
	CheckApprovalLevelByUserID(*gorm.DB, approverpayloads.GetByUserIDRequest) (*int, *exceptions.BaseErrorResponse)
	GetApprovers(*gorm.DB, int) ([]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse)
	GetAll(*gorm.DB) ([]approverentities.Approver, *exceptions.BaseErrorResponse)
	GetByLevel(*gorm.DB, approverpayloads.GetByLevelRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse)
	GetByUserID(*gorm.DB, approverpayloads.GetByUserIDRequest) (approverpayloads.GetApproverForApprovalByUserIDResponse, *exceptions.BaseErrorResponse)
	Get(*gorm.DB, int) (approverpayloads.CreateApprover, *exceptions.BaseErrorResponse)
	Create(*gorm.DB, approverpayloads.CreateApprover) (bool, *exceptions.BaseErrorResponse)
	Update(*gorm.DB, int, approverpayloads.UpdateApprover) (bool, *exceptions.BaseErrorResponse)
	Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
}
