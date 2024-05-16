package approverpayloads

type GetApproverResponse struct {
	IsActive        bool                    `json:"is_active"`
	ApproverID      int                     `json:"approver_id"`
	ApproverCode    string                  `json:"approver_code"`
	ApproverName    string                  `json:"approver_name"`
	ApproverDetails []UpdateApproverDetails `json:"approver_details"`
}

type GetByLevelRequest struct {
	ApprovalMappingID int `json:"approval_mapping_id"`
	ApprovalAmountID  int `json:"approval_amount_id"`
	Level             int `json:"level"`
}

type GetByUserIDRequest struct {
	ApprovalMappingID int `json:"approval_mapping_id"`
	ApprovalAmountID  int `json:"approval_amount_id"`
	UserID            int `json:"user_id"`
}
type GetApproverForApprovalResponse struct {
	ApproverID         *int    `json:"approver_id"`
	ApproverDetailID   *int    `json:"approver_detail_id"`
	ApprovalApproverID *int    `json:"approval_approver_id"`
	Level              *int    `json:"level"`
	UserID             *int    `json:"user_id"`
	Email              *string `json:"email"`
}

type GetApproverForApprovalByUserIDResponse struct {
	ApproverID         *int    `json:"approver_id"`
	ApproverDetailID   *int    `json:"approver_detail_id"`
	ApprovalApproverID *int    `json:"approval_approver_id"`
	Level              *int    `json:"level"`
	Email              *string `json:"email"`
}

type CreateApprover struct {
	CompanyID       int                     `json:"company_id"`
	ApproverCode    string                  `json:"approver_code"`
	ApproverName    string                  `json:"approver_name"`
	ApproverDetails []CreateApproverDetails `json:"approver_details"`
}

type UpdateApprover struct {
	IsActive        bool                    `json:"is_active"`
	CompanyID       int                     `json:"company_id"`
	ApproverCode    string                  `json:"approver_code"`
	ApproverName    string                  `json:"approver_name"`
	ApproverDetails []UpdateApproverDetails `json:"approver_details"`
}

type CreateApproverDetails struct {
	UserID int `json:"user_id"`
}

type UpdateApproverDetails struct {
	ApproverDetailID int `json:"approver_id"`
	CompanyID        int `json:"company_id"`
	UserID           int `json:"user_id"`
}
