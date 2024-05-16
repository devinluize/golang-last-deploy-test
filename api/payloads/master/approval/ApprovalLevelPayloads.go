package approvalpayloads

type GetApprovalLevelResponse struct {
	Level         *int  `json:"level"`
	CountRequired *int  `json:"count_required"`
	IsHierarchy   *bool `json:"is_hierarchy"`
}
