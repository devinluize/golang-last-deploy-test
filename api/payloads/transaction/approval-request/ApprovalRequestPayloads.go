package approvalrequestpayloads

import "time"

type GetApprovalRequest struct {
	ApprovalMappingID      int                            `json:"approval_mapping_id"`
	ApprovalAmountID       int                            `json:"approval_amount_id"`
	ApprovalApproverID     int                            `json:"approval_approver"`
	RequestDate            time.Time                      `json:"request_date"`
	SourceDocNo            string                         `json:"source_doc_no"`
	ApprovalRequestDetails []CreateApprovalRequestDetails `json:"approval_request_details"`
}

type GetApprovalRequestForRequesterResponse struct {
	ApprovalID        int    `json:"approval_id"`
	ApprovalMappingID int    `json:"approval_mapping_id"`
	ApprovalName      string `json:"approval_name"`
	Count             int    `json:"count"`
}

type GetApprovalRequestDetailsResponse struct {
	ApprovalRequestID int       `json:"approval_request_id"`
	CompanyID         int       `json:"company_id"`
	CompanyName       string    `json:"company_name"`
	RequestDate       time.Time `json:"request_date"`
	SourceDocNo       string    `json:"source_doc_no"`
	SourceAmount      float64   `json:"source_amount"`
	SourceDate        time.Time `json:"source_date"`
}

type GetApprovalRequestDetailsForRequesterResponse struct {
	SourceDate  time.Time `json:"source_date"`
	Amount      string    `json:"amount"`
	RequestDate time.Time `json:"request_date"`
	RequestBy   int       `json:"request_by"`
	Remark      string    `json:"remark"`
}

type CreateApprovalRequest struct {
	ApprovalMappingID      int                            `json:"approval_mapping_id"`
	ApprovalAmountID       int                            `json:"approval_amount_id"`
	CompanyID              int                            `json:"company_id"`
	ModuleID               int                            `json:"module_id"`
	DocumentTypeID         int                            `json:"document_type_id"`
	RequestDate            time.Time                      `json:"request_date"`
	SourceDocNo            string                         `json:"source_doc_no"`
	SourceSysNo            int                            `json:"source_sys_no"`
	SourceAmount           float64                        `json:"source_amount"`
	SourceDate             time.Time                      `json:"source_date"`
	TransactionTypeID      int                            `json:"transaction_type_id"`
	BrandID                int                            `json:"brand_id"`
	ProfitCenterID         int                            `json:"profit_center_id"`
	CostCenterID           int                            `json:"cost_center_id"`
	StatusID               int                            `json:"status_id"`
	IsVoid                 bool                           `json:"is_void"`
	RequestBy              int                            `json:"request_by"`
	ApprovalRequestDetails []CreateApprovalRequestDetails `json:"approval_request_details"`
}

type CreateApprovalRequestDetails struct {
	ApprovalRequestDetailsID int     `json:"approval_request_details"`
	ApprovalRequestID        int     `json:"approval_request_id"`
	ApprovalApproverID       int     `json:"approval_approver_id"`
	ApproverID               int     `json:"approver_id"`
	UserId                   int     `json:"user_id"`
	Level                    int     `json:"level"`
	StatusID                 int     `json:"status_id"`
	Remark                   *string `json:"remark"`
	// AllowApprove             bool    `json:"allow_approve"`
}

type UpdateApprovalRequest struct {
	ApprovalID             int                            `json:"approval_id"`
	StatusID               int                            `json:"status_id"`
	ApprovalRequestID      int                            `json:"approval_request_id"`
	ApproverDetailID       int                            `json:"approver_detail_id"`
	ApprovalRequestDetails []UpdateApprovalRequestDetails `json:"approval_request_details"`
}

type UpdateApprovalRequestDetails struct {
	ApprovalRequestID int `json:"approval_request_id"`
	// AllowApprove      bool    `json:"allow_approve"`
	StatusID  int     `json:"status_id"`
	Remark    *string `json:"remark"`
	CompanyID int     `json:"company_id"`
}

type UpdateStatusRequestDetails struct {
	ApprovalRequestDetailsID int `json:"approval_request_details_id"`
	StatusID                 int `json:"status_id"`
}
