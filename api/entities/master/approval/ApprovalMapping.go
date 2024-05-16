package approvalentities

import approvalrequestentities "user-services/api/entities/transaction/approval-request"

const TableNameApprovalMapping = "approval_mapping"

// ApprovalMapping mapped from table <approval>
type ApprovalMapping struct {
	IsActive          bool                                       `gorm:"column:is_active;not null" json:"is_active"`
	ID                int                                        `gorm:"column:id;primaryKey;autoIncrement;size:30;" json:"id"`
	CompanyID         int                                        `gorm:"column:company_id;uniqueIndex:idx_approval_mapping;size:30;not null" json:"company_id"`
	ApprovalID        int                                        `gorm:"column:approval_id;uniqueIndex:idx_approval_mapping;size:30;" json:"approval_id"`
	ModuleID          int                                        `gorm:"column:module_id;uniqueIndex:idx_approval_mapping;size:30;" json:"module_id"`
	DocumentTypeID    int                                        `gorm:"column:document_type_id;uniqueIndex:idx_approval_mapping;;size:30;" json:"document_type_id"`
	TransactionTypeID int                                        `gorm:"column:transaction_type_id;uniqueIndex:idx_approval_mapping;size:30;" json:"transaction_type_id"`
	BrandID           int                                        `gorm:"column:brand_id;uniqueIndex:idx_approval_mapping;size:30;" json:"brand_id"`
	ProfitCenterID    int                                        `gorm:"column:profit_center_id;uniqueIndex:idx_approval_mapping;size:30;" json:"profit_center_id"`
	CostCenterID      int                                        `gorm:"column:cost_center_id;uniqueIndex:idx_approval_mapping;size:30;" json:"cost_center_id"`
	Approval          *Approval                                  `gorm:"foreignKey:approval_id;references:id"`
	ApprovalAmount    []ApprovalAmount                           `gorm:"foreignKey:approval_mapping_id;references:id"`
	ApprovalRequest   []*approvalrequestentities.ApprovalRequest `gorm:"foreignKey:approval_mapping_id;references:id"`
	ApprovalApprover  []ApprovalApprover                         `gorm:"foreignKey:approval_mapping_id;references:id"`
}

// custom tablename
func (e *ApprovalMapping) TableName() string {
	return TableNameApprovalMapping
}
