package approvalservices

import (
	approvalentities "user-services/api/entities/master/approval"
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
	"user-services/api/utils"
)

type ApprovalService interface {
	Get(int) (approvalentities.Approval, *exceptions.BaseErrorResponse)
	GetAll(utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	Create(approvalpayloads.CreateApproval) (bool, *exceptions.BaseErrorResponse)
	Update(int, approvalpayloads.UpdateApproval) (bool, *exceptions.BaseErrorResponse)
	Delete(int) (bool, *exceptions.BaseErrorResponse)
}
