package approvalrepo

import (
	approvalentities "user-services/api/entities/master/approval"
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type ApprovalRepository interface {
	CheckDataExists(*gorm.DB, string) (bool, *exceptions.BaseErrorResponse)
	Get(*gorm.DB, int) (approvalentities.Approval, *exceptions.BaseErrorResponse)
	GetAll(*gorm.DB, utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	Create(*gorm.DB, approvalpayloads.CreateApproval) (bool, *exceptions.BaseErrorResponse)
	Update(*gorm.DB, int, approvalpayloads.UpdateApproval) (bool, *exceptions.BaseErrorResponse)
	Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
}
