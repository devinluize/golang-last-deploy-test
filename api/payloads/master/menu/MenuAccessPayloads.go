package menupayloads

type CreateMenuAccessParamRequest struct {
	RoleID     int   `json:"role_id"`
	MenuListID []int `json:"menu_list_id"`
}

type CreateMenuAccessRequest struct {
	MenuListID int `json:"menu_list_id"`
	RoleID     int `json:"role_id"`
}

type CheckCreateWithMenuListAndRoleResponse struct {
	MenuListID int `json:"menu_list_id"`
	RoleID     int `json:"role_id"`
}

type GetMenuByRoleIDResponse struct {
	MenuID      int                        `json:"menu_id"`
	MenuTitle   string                     `json:"title"`
	MenuUrlPath string                     `json:"url"`
	ParentID    int                        `json:"parent_id"`
	Children    []*GetMenuByRoleIDResponse `json:"children"`
}

type DeleteMenuAccessByIDRequest struct {
	MenuAccessID int `json:"menu_access_id"`
}

type GetUrlPathResponse struct {
	MenuUrlPath string `json:"menu_url_name"`
}
