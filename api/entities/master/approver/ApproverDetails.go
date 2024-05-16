package approverentities

const TableNameApproverDetails = "approver_details"

// Approver Detail mapped from table <approver detail>
type ApproverDetails struct {
	IsActive        bool              `gorm:"column:is_active;not null" json:"is_active"`
	ID              int               `gorm:"column:id;primaryKey;autoIncrement;size:30;" json:"id"`
	ApproverID      int               `gorm:"column:approver_id;index:idx_approver_details,unique;size:30;" json:"approver_id"`
	UserID          int               `gorm:"column:user_id;index:idx_approver_details,unique;size:30;" json:"user_id"`
	Approver        *Approver         `gorm:"foreignKey:approver_id;references:id"`
	ApproverCompany []ApproverCompany `gorm:"foreignKey:approver_details_id;references:id"`
}

// custom tablename
func (e *ApproverDetails) TableName() string {
	return TableNameApproverDetails
}
