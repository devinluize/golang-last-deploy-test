package menuentities

const TableNameMenuList string = "menu_list"

type MenuList struct {
	ID               int          `gorm:"column:id;primaryKey;size:30;" json:"id"`
	MenuID           int          `gorm:"column:menu_id;not null;size:30;" json:"menu_id"`
	MenuUrlID        int          `gorm:"column:menu_url_id;not null;uniqueIndex:idx_menu_url;size:30;" json:"menu_url_id"`
	ParentID         int          `gorm:"column:parent_id;size:30;" json:"parent_id"`
	Image            string       `gorm:"column:image;size:100" json:"image"`
	OrderNo          int          `gorm:"column:order_no;size:30;" json:"order_no"`
	MenuUrl          *MenuUrl     `gorm:"foreignKey:menu_url_id;references:id"`
	MenuMaster       *Menu        `gorm:"foreignKey:menu_id;references:id"`
	MenuAccessByList []MenuAccess `gorm:"many2many:menu_access_list;foreignKey:id;joinForeignKey:list_id;references:id;joinReferences:access_id"`
}

func (e *MenuList) TableName() string {
	return TableNameMenuList
}
