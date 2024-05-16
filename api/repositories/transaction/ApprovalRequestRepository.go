package approvalrequestrepo

import (
	"user-services/api/exceptions"
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type ApprovalRequestRepository interface {
	Get(*gorm.DB, int) (approvalrequestpayloads.GetApprovalRequest, *exceptions.BaseErrorResponse)
	GetByApprover(*gorm.DB, int, utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	GetDetailsByApprover(*gorm.DB, int, int, utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	GetByRequester(*gorm.DB, int, utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	GetDetailsByRequester(*gorm.DB, int, int, utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	Create(*gorm.DB, approvalrequestpayloads.CreateApprovalRequest) (bool, *exceptions.BaseErrorResponse)
	CreateDetails(*gorm.DB, []approvalrequestpayloads.CreateApprovalRequestDetails) (bool, *exceptions.BaseErrorResponse)
	Update(*gorm.DB, approvalrequestpayloads.UpdateApprovalRequest) (bool, *exceptions.BaseErrorResponse)
	UpdateDetails(*gorm.DB, approvalrequestpayloads.UpdateApprovalRequestDetails) (bool, *exceptions.BaseErrorResponse)
	UpdateStatusByApprover(*gorm.DB, approvalrequestpayloads.UpdateStatusRequestDetails) (bool, *exceptions.BaseErrorResponse)
}
