package usergroupentities

const TableNameUserGroups = "user_group"

// UserGroup mapped from table <UserGroup>
type UserGroup struct {
	IsActive bool             `gorm:"column:is_active;not null" json:"is_active"`
	ID       int              `gorm:"column:id;primaryKey;size:30;" json:"id"`
	Code     string           `gorm:"column:code;unique;not null;size:10;index:idx_user_group_code,unique" json:"code"`
	Name     string           `gorm:"column:name;unique;not null;size:50;index:idx_user_group_name,unique" json:"name"`
	Company  UserGroupCompany `gorm:"foreignKey:user_group_id;references:id"`
}

// custom tablename
func (e *UserGroup) TableName() string {
	return TableNameUserGroups
}
