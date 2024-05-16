package approverentities

const TableNameApproverCompany = "approver_company"

// Approver Detail mapped from table <approver detail>
type ApproverCompany struct {
	CompanyID        int `gorm:"column:company_id;index:idx_approver_companies,unique;size:30;" json:"company_id"`
	ApproverDetailID int `gorm:"column:approver_details_id;index:idx_approver_companies,unique;size:30;" json:"approver_detail_id"`
}

// custom tablename
func (e *ApproverCompany) TableName() string {
	return TableNameApproverDetails
}
