package masterentities

import (
	menuentities "user-services/api/entities/master/menu"
)

const TableNameRoles = "roles"

// role mapped from table <role>
type Role struct {
	IsActive   bool                    `gorm:"column:is_active;not null" json:"is_active"`
	ID         int                     `gorm:"column:id;size:30;primaryKey" json:"id"`
	Code       string                  `gorm:"column:code;unique;not null;size:10;index:idx_role_code,unique" json:"code"`
	Name       string                  `gorm:"column:name;unique;not null;size:50;index:idx_role_name,unique" json:"name"`
	Role       User                    `gorm:"foreignKey:role_id;references:id"`
	MenuByRole menuentities.MenuAccess `gorm:"foreignKey:role_id;references:id"`
}

// custom tablename
func (e *Role) TableName() string {
	return TableNameRoles
}
