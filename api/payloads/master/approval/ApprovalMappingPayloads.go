package approvalpayloads

type GetApprovalMappingRequest struct {
	CompanyID         int     `json:"company_id"`
	TransactionTypeID int     `json:"transaction_type_id"`
	BrandID           int     `json:"brand_id"`
	ProfitCenterID    int     `json:"profit_center_id"`
	CostCenterID      int     `json:"cost_center_id"`
	SourceAmount      float64 `json:"source_amount"`
	ModuleID          int     `json:"module_id"`
	DocumentTypeID    int     `json:"document_type_id"`
	IsVoid            bool    `json:"is_void"`
}

type GetApprovalMappingResponse struct {
	ApprovalMappingID int `json:"approval_mapping_id"`
	ApprovalAmountID  int `json:"approval_amount_id"`
}

type CreateApprovalMapping struct {
	ModuleID          int `json:"module_id"`
	DocumentTypeID    int `json:"document_type_id"`
	CompanyID         int `json:"company_id"`
	TransactionTypeID int `json:"transaction_type_id"`
	BrandID           int `json:"brand_id"`
	ProfitCenterID    int `json:"profit_center_id"`
	CostCenterID      int `json:"cost_center_id"`
	ApprovalAmount    []CreateApprovalAmount
}

type UpdateApprovalMapping struct {
	ApprovalMappingID int `json:"approval_amount_id"`
	ApprovalID        int `json:"approval_id"`
	ModuleID          int `json:"module_id"`
	DocumentTypeID    int `json:"document_type_id"`
	CompanyID         int `json:"company_id"`
	TransactionTypeID int `json:"transaction_type_id"`
	BrandID           int `json:"brand_id"`
	ProfitCenterID    int `json:"profit_center_id"`
	CostCenterID      int `json:"cost_center_id"`
	ApprovalAmount    []UpdateApprovalAmount
}
