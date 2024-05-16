package approvalrepo

import (
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"

	"gorm.io/gorm"
)

type ApprovalLevelRepository interface {
	GetByApprovalApproverID(*gorm.DB, int, int, int) (approvalpayloads.GetApprovalLevelResponse, *exceptions.BaseErrorResponse)
	GetByUserID(*gorm.DB, int, int, int) (approvalpayloads.GetApprovalLevelResponse, *exceptions.BaseErrorResponse)
}
