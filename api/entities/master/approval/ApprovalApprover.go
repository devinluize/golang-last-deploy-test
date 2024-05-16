package approvalentities

import approvalrequestentities "user-services/api/entities/transaction/approval-request"

const TableNameApprovalApprover = "approval_approver"

// ApprovalApprover mapped from table <ApprovalApprover>
type ApprovalApprover struct {
	IsActive               bool                                             `gorm:"column:is_active;not null" json:"is_active"`
	ID                     int                                              `gorm:"column:id;primaryKey;autoIncrement;size:30;" json:"id"`
	ApprovalMappingID      int                                              `gorm:"size:30;column:approval_mapping_id;uniqueIndex:idx_approval_approvers" json:"approval_mapping_id"`
	ApproverID             int                                              `gorm:"column:approver_id;uniqueIndex:idx_approval_approvers;size:30;" json:"approver_id"`
	ApprovalMapping        *ApprovalMapping                                 `gorm:"foreignKey:approval_mapping_id;references:id"`
	ApprovalLevel          []ApprovalLevel                                  `gorm:"foreignKey:approval_approver_id;references:id"`
	ApprovalRequestDetails []approvalrequestentities.ApprovalRequestDetails `gorm:"foreignKey:approval_approver_id;references:id"`
}

// custom tablename
func (e *ApprovalApprover) TableName() string {
	return TableNameApprovalApprover
}
