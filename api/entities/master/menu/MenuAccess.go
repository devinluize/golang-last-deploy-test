package menuentities

const TableNameMenuAccess string = "menu_access"

type MenuAccess struct {
	ID               int            `gorm:"column:id;primaryKey;size:30;" json:"id"`
	RoleID           int            `gorm:"column:role_id;size:30;index:idx_menu_access,unique;" json:"role_id"`
	CompanyID        int            `gorm:"column:company_id;size:30;index:idx_menu_access,unique;" json:"company_id"`
	MenuUserAccess   MenuUserAccess `gorm:"foreignKey:menu_access_id;references:id" json:"-"`
	MenuListByAccess []MenuList     `gorm:"many2many:menu_access_list;foreignKey:id;joinForeignKey:access_id;references:id;joinReferences:list_id"`
}

func (e *MenuAccess) TableName() string {
	return TableNameMenuAccess
}
