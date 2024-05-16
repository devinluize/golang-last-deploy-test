package payloads

import "time"

type ForgotPasswordInput struct {
	Email string `json:"email" validate:"required,email"`
}

type ChangePasswordInput struct {
	OldPassword string `json:"old_password" validate:"required" `
	Password    string `json:"password" validate:"required" `
	NewPassword string `json:"new_password" validate:"required,nefield=OldPassword,eqfield=Password"`
}

type ResetPasswordInput struct {
	Password        string `validate:"required,eqfield=PasswordConfirm" json:"password"`
	PasswordConfirm string `validate:"required" json:"password_confirm"`
}

type UpdateEmailTokenRequest struct {
	Email              string     `validate:"required,email" json:"email"`
	PasswordResetToken *string    `json:"password_reset_token"`
	PasswordResetAt    *time.Time `json:"password_reset_at"`
}

type ResetPasswordRequest struct {
	PasswordResetToken *string    `json:"password_reset_token"`
	PasswordResetAt    *time.Time `json:"password_reset_at"`
}

type UpdatePasswordByTokenRequest struct {
	PasswordResetToken *string `json:"password_reset_token"`
	Password           string  `json:"password"`
}
