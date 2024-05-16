package approvalrequestentities

const TableNameApprovalRequestDetails = "approval_request_details"

type ApprovalRequestDetails struct {
	ID                 int     `gorm:"size:30;column:id;primaryKey;autoIncrement" json:"id"`
	ApprovalRequestID  int     `gorm:"size:30;column:approval_request_id" json:"approval_request_id"`
	ApprovalApproverID int     `gorm:"size:30;column:approval_approver_id;" json:"approval_approver_id"`
	ApproverID         int     `gorm:"size:30;column:approver_id" json:"approver_id"`
	UserID             int     `gorm:"size:30;column:user_id" json:"user_id"`
	Level              int     `gorm:"size:30;column:level" json:"level"`
	StatusID           int     `gorm:"size:30;column:status_id" json:"status_id"`
	Remark             *string `gorm:"size:100;column:remark" json:"remark"`
	// AllowApprove       bool    `gorm:"column:allow_approve" json:"allow_approve"`
}

// custom tablename
func (e *ApprovalRequestDetails) TableName() string {
	return TableNameApprovalRequestDetails
}
