package userservices

import (
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/payloads"
)

type UserService interface {
	GetCurrentUser(int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse)
	FindUser(string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	CheckUserExists(string) (bool, *exceptions.BaseErrorResponse)
	ViewUser() ([]masterentities.User, *exceptions.BaseErrorResponse)
	GetEmails([]int) ([]string, *exceptions.BaseErrorResponse)
	GetByID(int) (masterentities.User, *exceptions.BaseErrorResponse)
	GetUsernameByUserID(int) (string, *exceptions.BaseErrorResponse)
	GetUserIDByUsername(string) (int, *exceptions.BaseErrorResponse)
	GetUser(string) (masterentities.User, *exceptions.BaseErrorResponse)
	GetUserDetailByUsername(string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	UpdateUser(payloads.CreateRequest, int) (bool, *exceptions.BaseErrorResponse)
	DeleteUser(int) (bool, *exceptions.BaseErrorResponse)
}
