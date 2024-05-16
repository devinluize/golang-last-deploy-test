package usergrouppayloads

type CreateUserGroupRequest struct {
	UserGroupCode    string                   `json:"user_group_code"`
	UserGroupName    string                   `json:"user_group_name"`
	UserGroupCompany []CreateUserGroupCompany `json:"user_group_companies"`
}

type UpdateUserGroupRequest struct {
	UserGroupCode    string                   `json:"user_group_code"`
	UserGroupName    string                   `json:"user_group_name"`
	UserGroupCompany []CreateUserGroupCompany `json:"user_group_companies"`
}

type CreateUserGroupCompany struct {
	CompanyID        int                      `json:"company_id"`
	UserGroupLeaders []CreateUserGroupLeaders `json:"user_group_leaders"`
}

type CreateUserGroupLeaders struct {
	UserID           int                     `json:"user_id"`
	UserGroupMembers []CreateUserGroupMember `json:"user_group_members"`
}

type CreateUserGroupMember struct {
	UserID int `json:"user_id"`
}

type GetLeaderForApprovalRequest struct {
	CompanyID int `json:"company_id"`
	UserID    int `json:"user_id"`
}
