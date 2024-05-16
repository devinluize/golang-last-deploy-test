package approvalentities

const TableNameApproval = "approval"

// Approval mapped from table <Approval>
type Approval struct {
	IsActive        bool              `gorm:"column:is_active;not null" json:"is_active"`
	ID              int               `gorm:"size:30;column:id;primaryKey;autoIncrement" json:"id"`
	Code            string            `gorm:"size:30;column:code;unique" json:"code"`
	Name            string            `gorm:"size:100;column:name;unique" json:"name"`
	IsVoid          bool              `gorm:"column:is_void;" json:"is_void"`
	ApprovalMapping []ApprovalMapping `gorm:"foreignKey:approval_id;references:id"`
}

// custom tablename
func (e *Approval) TableName() string {
	return TableNameApproval
}
