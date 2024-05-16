package usergroupentities

const TableNameUserGroupCompany = "user_group_company"

// UserGroupCompany mapped from table <UserGroupCompany>
type UserGroupCompany struct {
	ID          int             `gorm:"column:id;primaryKey;size:30;" json:"id"`
	UserGroupID int             `gorm:"column:user_group_id;size:30;not null" json:"user_group_id"`
	CompanyID   int             `gorm:"column:company_id;size:30;not null" json:"company_id"`
	Leaders     UserGroupLeader `gorm:"foreignKey:user_group_company_id;references:id"`
}

// custom tablename
func (e *UserGroupCompany) TableName() string {
	return TableNameUserGroupCompany
}
