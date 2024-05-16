package approverservices

import (
	approverentities "user-services/api/entities/master/approver"
	"user-services/api/exceptions"
	approverpayloads "user-services/api/payloads/master/approver"
)

type ApproverService interface {
	GetApprovers(int) ([]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse)
	GetByLevel(approverpayloads.GetByLevelRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse)
	GetByUserID(approverpayloads.GetByUserIDRequest) (approverpayloads.GetApproverForApprovalByUserIDResponse, *exceptions.BaseErrorResponse)
	GetAll() ([]approverentities.Approver, *exceptions.BaseErrorResponse)
	Get(int) (approverpayloads.CreateApprover, *exceptions.BaseErrorResponse)
	Create(approverpayloads.CreateApprover) (bool, *exceptions.BaseErrorResponse)
	Update(int, approverpayloads.UpdateApprover) (bool, *exceptions.BaseErrorResponse)
	Delete(int) (bool, *exceptions.BaseErrorResponse)
}
