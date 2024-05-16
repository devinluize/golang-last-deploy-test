package approvalentities

import approvalrequestentities "user-services/api/entities/transaction/approval-request"

const TableNameApprovalAmount = "approval_amount"

// ApprovalAmount mapped from table <ApprovalAmount>
type ApprovalAmount struct {
	IsActive          bool                                      `gorm:"column:is_active;not null" json:"is_active"`
	ID                int                                       `gorm:"size:30;column:id;primaryKey;autoIncrement" json:"id"`
	ApprovalMappingID int                                       `gorm:"size:30;column:approval_mapping_id;uniqueIndex:idx_approval_amount" json:"approval_mapping_id"`
	MaxAmount         float64                                   `gorm:"column:max_amount;uniqueIndex:idx_approval_amount" json:"max_amount"`
	ApprovalLevel     []ApprovalLevel                           `gorm:"foreignKey:approval_amount_id;references:id"`
	ApprovalRequest   []approvalrequestentities.ApprovalRequest `gorm:"foreignKey:approval_amount_id;references:id"`
}

// custom tablename
func (e *ApprovalAmount) TableName() string {
	return TableNameApprovalAmount
}
