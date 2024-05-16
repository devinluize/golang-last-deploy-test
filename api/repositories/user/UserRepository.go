package userrepo

import (
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/payloads"

	"gorm.io/gorm"
)

type UserRepository interface {
	CheckUserExists(*gorm.DB, string) (bool, *exceptions.BaseErrorResponse)
	FindUser(*gorm.DB, string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	ViewUser(*gorm.DB) ([]masterentities.User, *exceptions.BaseErrorResponse)
	GetCurrentUser(*gorm.DB, int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse)
	GetByID(*gorm.DB, int) (masterentities.User, *exceptions.BaseErrorResponse)
	GetByUsername(*gorm.DB, string) (masterentities.User, *exceptions.BaseErrorResponse)
	GetEmails(*gorm.DB, []int) ([]string, *exceptions.BaseErrorResponse)
	GetUsernameByUserID(*gorm.DB, int) (string, *exceptions.BaseErrorResponse)
	GetUserIDByUsername(*gorm.DB, string) (int, *exceptions.BaseErrorResponse)
	GetByEmail(*gorm.DB, string) (bool, *exceptions.BaseErrorResponse)
	GetRoleByCompanyAndUserID(*gorm.DB, int, int) (int, *exceptions.BaseErrorResponse)
	GetUserDetailByUsername(*gorm.DB, string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	Create(*gorm.DB, payloads.CreateRequest, int) (int, *exceptions.BaseErrorResponse)
	Update(*gorm.DB, payloads.CreateRequest, int) (bool, *exceptions.BaseErrorResponse)
	Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
}
