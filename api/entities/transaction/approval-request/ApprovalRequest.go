package approvalrequestentities

import (
	"time"
)

const TableNameApprovalRequest = "approval_request"

type ApprovalRequest struct {
	ID                     int                      `gorm:"column:id;primaryKey;autoIncrement;size:30;" json:"id"`
	ApprovalMappingID      int                      `gorm:"column:approval_mapping_id;size:30;" json:"approval_mapping_id"`
	ApprovalAmountID       int                      `gorm:"column:approval_amount_id;size:30;" json:"approval_amount_id"`
	RequestDate            time.Time                `gorm:"column:request_date" json:"request_date"`
	StatusID               int                      `gorm:"column:status_id;size:30;" json:"status_id"`
	SourceDocNo            string                   `gorm:"column:source_doc_no;size:100;" json:"source_doc_no"`
	RequestBy              int                      `gorm:"column:request_by;size:30;" json:"request_by"`
	ApprovalRequestDetails []ApprovalRequestDetails `gorm:"foreignKey:approval_request_id;references:id"`
}

// custom tablename
func (e *ApprovalRequest) TableName() string {
	return TableNameApprovalRequest
}
