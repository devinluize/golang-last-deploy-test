package approvalrequestservices

import (
	"user-services/api/exceptions"
	"user-services/api/payloads"
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
	"user-services/api/utils"
)

type ApprovalRequestService interface {
	Get(int) (approvalrequestpayloads.GetApprovalRequest, *exceptions.BaseErrorResponse)
	GetByUserType(string, int, utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	GetDetailsByUserType(string, int, int, utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse)
	Create(*payloads.UserDetail, approvalrequestpayloads.CreateApprovalRequest) (bool, *exceptions.BaseErrorResponse)
	CreateDetails(*payloads.UserDetail, approvalrequestpayloads.UpdateApprovalRequestDetails) (bool, *exceptions.BaseErrorResponse)
	Update(approvalrequestpayloads.UpdateApprovalRequest) (bool, *exceptions.BaseErrorResponse)
	UpdateDetails(approvalrequestpayloads.UpdateApprovalRequestDetails) (bool, *exceptions.BaseErrorResponse)
	UpdateStatusByApprover(approvalrequestpayloads.UpdateStatusRequestDetails) (bool, *exceptions.BaseErrorResponse)
}
