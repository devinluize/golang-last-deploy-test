package payloads

type CreateRequest struct {
	Username string `json:"username" validate:"required,max=30,min=5" `
	Email    string `json:"email" validate:"required,email"`
	IsActive bool   `json:"is_active" validate:"required"`
	Password string `json:"password" validate:"required,max=100,min=5"`
}
type RegisterRequest struct {
	UserName   string `gorm:"column:user_name" json:"user_name"`
	UserEmail  string `gorm:"column:user_email" json:"user_email"`
	Password   string `gorm:"column:password" json:"password"`
	UserRoleId int    `gorm:"column:user_role" json:"user_role_id"`
}

type UserDetails struct {
	Role      int    `json:"role"`
	CompanyID string `json:"company_id"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Client   string `json:"client" validate:"required"`
}
type LoginRequestPayloads struct {
	UserName string `gorm:"column:user_name" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
type LoginResponses struct {
	UserName string `json:"username"`
	UserRole int    `json:"user_role"`
	Token    string `json:"token"`
}
type LoginCredential struct {
	Client    string `json:"client"`
	IpAddress string `json:"ip_address"`
	Session   string `json:"session"`
}

type UserDetail struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Authorize string `json:"authorized"`
	CompanyID string `json:"company_id"`
	Role      int    `json:"role"`
	IpAddress string `json:"ip_address"`
	Client    string `json:"client"`
}

type CurrentUserResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
