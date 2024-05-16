package usergroupentities

const TableNameUserGroupMembers = "user_group_members"

// UserGroup mapped from table <UserGroup>
type UserGroupMember struct {
	ID      int               `gorm:"column:id;primaryKey;size:30;" json:"id"`
	UserID  int               `gorm:"column:user_id;uniqueIndex:idx_user_group_member;size:30;" json:"user_id"`
	Leaders []UserGroupLeader `gorm:"many2many:user_group_hierarchy;foreignKey:id;joinForeignKey:member_id;references:id;joinReferences:leader_id"`
}

// custom tablename
func (e *UserGroupMember) TableName() string {
	return TableNameUserGroupMembers
}
