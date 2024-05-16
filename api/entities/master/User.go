package masterentities

import (
	"time"
	approverentities "user-services/api/entities/master/approver"
	menuentities "user-services/api/entities/master/menu"
	usergroupentities "user-services/api/entities/master/user-group"
	approvalrequestentities "user-services/api/entities/transaction/approval-request"
)

const TableNameUser = "users"

// User mapped from table <user>
type User struct {
	IsActive               bool                                             `gorm:"column:is_active;not null" json:"is_active"`
	ID                     int                                              `gorm:"column:id;primaryKey;autoIncrement;size:30" json:"id"`
	Password               string                                           `gorm:"column:password;size:100;not null" json:"password"`
	Username               string                                           `gorm:"column:username;size:30;unique;not null" json:"username"`
	Email                  string                                           `gorm:"column:email;size:255;unique;not null" json:"email"`
	RoleID                 int                                              `gorm:"column:role_id;size:30;not null" json:"role_id"`
	LastLogin              time.Time                                        `gorm:"column:last_login" json:"last_login"`
	DateJoined             time.Time                                        `gorm:"column:date_joined;not null" json:"date_joined"`
	PasswordResetToken     *string                                          `gorm:"column:password_reset_token" json:"password_reset_token"`
	PasswordResetAt        *time.Time                                       `gorm:"column:password_reset_at" json:"password_reset_at"`
	IpAddress              string                                           `gorm:"column:ip_address;null;size:20" json:"ip_address"`
	OtpEnabled             bool                                             `gorm:"column:otp_enabled;default:false;"`
	OtpVerified            bool                                             `gorm:"column:otp_verified;default:false;"`
	OtpSecret              string                                           `gorm:"column:otp_secret;default:false;"`
	OtpAuthUrl             string                                           `gorm:"column:otp_auth_url;default:false;"`
	CompanyID              int                                              `gorm:"column:company_id;not null;size:30" json:"company_id"`
	ApproverDetails        []approverentities.ApproverDetails               `gorm:"foreignKey:user_id;references:id"`
	Leaders                usergroupentities.UserGroupLeader                `gorm:"foreignKey:user_id;references:id"`
	Members                usergroupentities.UserGroupMember                `gorm:"foreignKey:user_id;references:id"`
	MenuUserAccess         menuentities.MenuUserAccess                      `gorm:"foreignKey:user_id;references:id"`
	ApprovalRequestDetails []approvalrequestentities.ApprovalRequestDetails `gorm:"foreignKey:user_id;references:id"`
	RequestBy              approvalrequestentities.ApprovalRequest          `gorm:"foreignKey:request_by;references:id"`
	// VerificationCode   string                           `gorm:"column:verification_token;not null" json:"verification_token"`

}
type UserEntities struct {
	UserName   string `gorm:"column:user_name" json:"user_name"`
	UserEmail  string `gorm:"column:user_email" json:"user_email"`
	Password   string `gorm:"column:password" json:"password"`
	UserRoleId int    `gorm:"column:user_role" json:"user_role_id"`
}

func (UserEntities) TableName() string {
	return "api.GoTestUserServices"
}

type OTPInput struct {
	UserID int    `json:"user_id" validate:"required"`
	Token  string `json:"token" validate:"required"`
	Client string `json:"client" validate:"required"`
}

// custom tablename
func (e *User) TableName() string {
	return TableNameUser
}
