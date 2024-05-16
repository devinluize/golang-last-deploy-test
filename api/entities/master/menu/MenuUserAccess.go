package menuentities

const TableNameMenuUserAccess string = "menu_user_access"

type MenuUserAccess struct {
	ID           int `gorm:"column:id;primaryKey;size:30;" json:"id"`
	UserID       int `gorm:"column:user_id;size:30;index:idx_menu_user_access,unique;" json:"user_id"`
	MenuAccessID int `gorm:"column:menu_access_id;size:30;index:idx_menu_user_access,unique;" json:"menu_access_id"`
	MenuAccess   *MenuAccess
}

func (e *MenuUserAccess) TableName() string {
	return TableNameMenuUserAccess
}
