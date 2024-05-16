package menupayloads

type CreateMenuListRequest struct {
	MenuTitle string `json:"menu_title"`
	MenuUrlID int    `json:"menu_url_id"`
	ParentID  int    `json:"parent_id"`
	MenuImage string `json:"menu_image"`
}

type GetMenuListByTitleResponse struct {
	MenuListID int    `json:"menu_list_id"`
	MenuTitle  string `json:"menu_title"`
	MenuUrlID  int    `json:"menu_url_id"`
	ParentID   int    `json:"parent_id"`
	MenuImage  string `json:"menu_image"`
}

type DropDownListForMenuListResponse struct {
	MenuListID int    `json:"menu_list_id"`
	MenuTitle  string `json:"menu_title"`
}
