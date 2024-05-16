package approvalservices

import (
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
)

type ApprovalLevelService interface {
	GetByApprovalApproverID(int, int, int) (approvalpayloads.GetApprovalLevelResponse, *exceptions.BaseErrorResponse)
	GetByUserID(int, int, int) (approvalpayloads.GetApprovalLevelResponse, *exceptions.BaseErrorResponse)
}
