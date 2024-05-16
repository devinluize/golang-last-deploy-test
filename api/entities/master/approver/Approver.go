package approverentities

import (
	approvalentities "user-services/api/entities/master/approval"
	approvalrequestentities "user-services/api/entities/transaction/approval-request"
)

const TableNameApprover = "approver"

// Approver mapped from table <approver>
type Approver struct {
	IsActive         bool                                             `gorm:"column:is_active;not null" json:"is_active"`
	ID               int                                              `gorm:"column:id;primaryKey;autoIncrement;size:30;" json:"id"`
	Code             string                                           `gorm:"column:code;uniqueIndex:idx_approver_code;size:30;" json:"code"`
	Name             string                                           `gorm:"column:name;uniqueIndex:idx_approver_name;size:100;" json:"name"`
	ApproverDetails  []ApproverDetails                                `gorm:"foreignKey:approver_id;references:id"`
	ApprovalApprover []approvalentities.ApprovalApprover              `gorm:"foreignKey:approver_id;references:id"`
	ApprovalRequest  []approvalrequestentities.ApprovalRequestDetails `gorm:"foreignKey:approver_id;references:id"`
}

// custom tablename
func (e *Approver) TableName() string {
	return TableNameApprover
}
