package approvalpayloads

type GetApproval struct {
	ApprovalID     int    `json:"approval_id"`
	ApprovalCode   string `json:"approval_code"`
	ApprovalName   string `json:"approval_name"`
	ModuleID       int    `json:"module_id"`
	DocumentTypeID int    `json:"document_type_id"`
}

type CreateApproval struct {
	ApprovalCode    string `json:"approval_code" validate:"required"`
	ApprovalName    string `json:"approval_name" validate:"required"`
	ApprovalMapping []CreateApprovalMapping
}

type UpdateApproval struct {
	ApprovalID      int    `json:"approval_id"`
	ApprovalCode    string `json:"approval_code"`
	ApprovalName    string `json:"approval_name"`
	ApprovalMapping []UpdateApprovalMapping
}

type CreateApprovalAmount struct {
	MaxAmount     float64               `json:"max_amount"`
	ApprovalLevel []CreateApprovalLevel `json:"approval_level"`
}

type CreateApprovalLevel struct {
	Level              int  `json:"level"`
	IsHierarchy        bool `json:"is_hierarchy"`
	ApprovalApproverID int  `json:"approval_approver_id"`
}

type UpdateApprovalAmount struct {
	ApprovalAmountID int                   `json:"approval_amount_id"`
	ApprovalLevel    []UpdateApprovalLevel `json:"approval_level"`
	MaxAmount        float64               `json:"max_amount"`
}

type UpdateApprovalLevel struct {
	ApprovalLevelID    int `json:"approval_level_id"`
	Level              int `json:"level"`
	ApprovalApproverID int `json:"approval_approver_id"`
}
