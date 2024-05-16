package approvalservices

import (
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
)

type ApprovalMappingService interface {
	Get(approvalpayloads.GetApprovalMappingRequest) (approvalpayloads.GetApprovalMappingResponse, *exceptions.BaseErrorResponse)
	Create(approvalpayloads.CreateApprovalMapping) (bool, *exceptions.BaseErrorResponse)
	Update(approvalpayloads.UpdateApprovalMapping) (bool, *exceptions.BaseErrorResponse)
	Delete(int) (bool, *exceptions.BaseErrorResponse)
}
