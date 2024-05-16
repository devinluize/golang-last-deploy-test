package menuentities

const TableNameMenu string = "menu"

type Menu struct {
	ID       int      `gorm:"column:id;primaryKey;size:30;" json:"id"`
	Title    string   `gorm:"column:title;index:idx_menu_list,unique;size:100;" json:"title"`
	MenuList MenuList `gorm:"foreignKey:menu_id;references:id" json:"menu_list"`
}

func (e *Menu) TableName() string {
	return TableNameMenu
}
