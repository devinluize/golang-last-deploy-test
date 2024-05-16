package usergroupentities

const TableNameUserGroupLeaders = "user_group_leaders"

// UserGroup mapped from table <UserGroup>
type UserGroupLeader struct {
	ID                 int               `gorm:"column:id;primaryKey;size:30;" json:"id"`
	UserID             int               `gorm:"column:user_id;uniqueIndex:idx_user_group_leader;size:30;" json:"user_id"`
	UserGroupCompanyID int               `gorm:"column:user_group_company_id;uniqueIndex:idx_user_group_leader;size:30;not null" json:"user_group_company_id"`
	Members            []UserGroupMember `gorm:"many2many:user_group_hierarchy;foreignKey:id;joinForeignKey:leader_id;references:id;joinReferences:member_id"`
}

// custom tablename
func (e *UserGroupLeader) TableName() string {
	return TableNameUserGroupLeaders
}
