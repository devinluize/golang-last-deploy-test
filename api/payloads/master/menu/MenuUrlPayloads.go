package menupayloads

type CreateMenuUrlRequest struct {
	MenuUrlPath        string `json:"menu_url_name"`
	MenuUrlDescription string `json:"menu_url_description"`
}

type GetMenuUrlByName struct {
	MenuUrlID          int    `json:"menu_url_id"`
	MenuUrlPath        string `json:"menu_url_name"`
	MenuUrlDescription string `json:"menu_url_description"`
}
