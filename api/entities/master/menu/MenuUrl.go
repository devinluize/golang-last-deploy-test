package menuentities

const TableNameMenuUrl string = "menu_url"

type MenuUrl struct {
	ID          int      `gorm:"column:id;primaryKey;size:30;" json:"id"`
	Path        string   `gorm:"column:path;size:100;" json:"path"`
	Description string   `gorm:"column:description;size:100;default:null;" json:"description"`
	MenuList    MenuList `gorm:"foreignKey:menu_url_id;references:id" json:"menu_list"`
}

func (e *MenuUrl) TableName() string {
	return TableNameMenuUrl
}
