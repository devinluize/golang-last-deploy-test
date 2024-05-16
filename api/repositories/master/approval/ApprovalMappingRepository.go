package approvalrepo

import (
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"

	"gorm.io/gorm"
)

type ApprovalMappingRepository interface {
	Get(*gorm.DB, approvalpayloads.GetApprovalMappingRequest) (approvalpayloads.GetApprovalMappingResponse, *exceptions.BaseErrorResponse)
	Create(*gorm.DB, approvalpayloads.CreateApprovalMapping) (bool, *exceptions.BaseErrorResponse)
	Update(*gorm.DB, approvalpayloads.UpdateApprovalMapping) (bool, *exceptions.BaseErrorResponse)
	Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
}
