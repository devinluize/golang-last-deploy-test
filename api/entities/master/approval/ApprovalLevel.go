package approvalentities

const TableNameApprovalLevel = "approval_level"

// ApprovalLevel mapped from table <ApprovalLevel>
type ApprovalLevel struct {
	IsActive           bool `gorm:"column:is_active;not null" json:"is_active"`
	ID                 int  `gorm:"column:id;primaryKey;autoIncrement;size:30;" json:"id"`
	ApprovalAmountID   int  `gorm:"column:approval_amount_id;uniqueIndex:idx_approval_level;size:30;" json:"approval_amount_id"`
	Level              int  `gorm:"column:level;uniqueIndex:idx_approval_level;size:30;" json:"level"`
	CountRequired      int  `gorm:"column:count_required;uniqueIndex:idx_approval_level;size:30;" json:"count_required"`
	IsHierarchy        bool `gorm:"column:is_hierarchy;uniqueIndex:idx_approval_level;size:30;" json:"is_hierarchy"`
	ApprovalApproverID int  `gorm:"column:approval_approver_id;uniqueIndex:idx_approval_level;size:30;" json:"approval_approver_id"`
	ApprovalApprover   *ApprovalApprover
}

// custom tablename
func (e *ApprovalLevel) TableName() string {
	return TableNameApprovalLevel
}
