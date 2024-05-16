package userrepoimpl

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/payloads"
	userrepo "user-services/api/repositories/user"
	"user-services/api/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

// GetCurrentUser implements userrepo.UserRepository.
func (*UserRepositoryImpl) GetCurrentUser(tx *gorm.DB, userID int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse) {
	var user masterentities.User
	var userResponse payloads.CurrentUserResponse
	err := tx.Model(user).
		Select(
			"id user_id",
			"username",
		).
		Where(masterentities.User{
			ID: userID,
		}).
		Scan(&userResponse).
		Error

	if err != nil {
		return userResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return userResponse, nil
}

// GetRoleByCompanyAndUserID implements userrepo.UserRepository.
func (*UserRepositoryImpl) GetRoleByCompanyAndUserID(tx *gorm.DB, companyID int, userID int) (int, *exceptions.BaseErrorResponse) {
	var user masterentities.User
	var role int
	err := tx.Model(user).
		Select(
			"MenuUserAccess__MenuAccess.role_id",
		).
		InnerJoins("MenuUserAccess",
			tx.Select("1"),
		).
		InnerJoins("MenuUserAccess.MenuAccess",
			tx.Select("1").
				Where("MenuUserAccess__MenuAccess.company_id = ?", companyID),
		).
		Where(masterentities.User{
			ID: userID,
		}).
		Scan(&role).
		Error

	if err != nil {
		return role, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if role == 0 {
		return role, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusForbidden,
			Message:    fmt.Sprintf("%s %s", utils.PermissionError, " of any menus"),
			Err:        errors.New(utils.PermissionError),
		}
	}

	return role, nil
}

func NewUserRepository() userrepo.UserRepository {
	return &UserRepositoryImpl{}
}
func (*UserRepositoryImpl) CheckUserExists(tx *gorm.DB, username string) (bool, *exceptions.BaseErrorResponse) {
	var user masterentities.User
	var exists bool
	err := tx.Model(user).
		Select(
			"count(id)",
		).
		Where(masterentities.User{Username: username}).
		Find(&exists).
		Error
	if exists {
		return exists, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}
	if err != nil {
		return exists, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return exists, nil
}

func (*UserRepositoryImpl) FindUser(tx *gorm.DB, username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	var user masterentities.User
	var userDetail payloads.UserDetails
	err := tx.Model(user).
		Select(
			"role",
			"company_id",
		).
		Where(masterentities.User{Username: username}).
		Scan(&userDetail).
		Error

	if err != nil {

		return userDetail, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return userDetail, nil
}

func (*UserRepositoryImpl) ViewUser(tx *gorm.DB) ([]masterentities.User, *exceptions.BaseErrorResponse) {
	var user []masterentities.User
	row, err := tx.Model(user).Scan(&user).Rows()

	if err != nil {

		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	var users []masterentities.User
	for row.Next() {
		var user masterentities.User

		err := row.Scan(
			&user.ID,
			&user.Username,
			&user.Password)

		if err != nil {

			return users, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
		users = append(users, user)
	}

	return users, nil
}

func (*UserRepositoryImpl) GetByUsername(tx *gorm.DB, username string) (masterentities.User, *exceptions.BaseErrorResponse) {
	var user masterentities.User

	row, err := tx.
		Model(user).
		Where(masterentities.User{
			Username: username,
		}).
		Scan(&user).
		Rows()

	if err != nil {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if user.Username == "" {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	defer row.Close()

	return user, nil
}

func (*UserRepositoryImpl) GetByEmail(tx *gorm.DB, email string) (bool, *exceptions.BaseErrorResponse) {
	var user masterentities.User

	row, err := tx.
		Model(user).
		Where(masterentities.User{
			Email: email,
		}).
		Scan(&user).
		Rows()

	if err != nil {

		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if user.Email == "" {

		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// GetEmails implements repositories.UserRepository.
func (*UserRepositoryImpl) GetEmails(tx *gorm.DB, userIDs []int) ([]string, *exceptions.BaseErrorResponse) {
	var email []string

	row, err := tx.
		Model(masterentities.User{}).
		Select("email").
		Where("id in (?)", userIDs).
		Scan(email).
		Rows()

	if err != nil {

		return email, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return email, nil
}
func (*UserRepositoryImpl) GetUserDetailByUsername(tx *gorm.DB, username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	var user payloads.UserDetails

	row, err := tx.
		Model(masterentities.User{}).
		Where(masterentities.User{Username: username}).
		Scan(&user).
		Rows()

	if err != nil {

		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return user, nil
}

func (*UserRepositoryImpl) GetByID(tx *gorm.DB, userID int) (masterentities.User, *exceptions.BaseErrorResponse) {
	var user masterentities.User

	rows, err := tx.Model(user).
		Where(masterentities.User{
			ID: userID,
		}).Scan(&user).
		Rows()

	if err != nil {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if user.Username == "" {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}
	defer rows.Close()

	return user, nil
}

// GetUsernameByUserID implements repositories.UserRepository.
func (*UserRepositoryImpl) GetUserIDByUsername(tx *gorm.DB, username string) (int, *exceptions.BaseErrorResponse) {
	var user masterentities.User

	err := tx.
		Select("id").
		Where(masterentities.User{
			Username: username,
		}).
		Find(&user).
		Error

	if err != nil {
		logrus.Info(err)

		return 0, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if user.ID == 0 {
		return 0, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return user.ID, nil
}

// GetUsernameByUserID implements repositories.UserRepository.
func (*UserRepositoryImpl) GetUsernameByUserID(tx *gorm.DB, userID int) (string, *exceptions.BaseErrorResponse) {
	var user masterentities.User

	err := tx.
		Select("username").
		Where(masterentities.User{
			ID: userID,
		}).
		Find(&user).
		Error

	if err != nil {

		return user.Username, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if user.Username == "" {
		return user.Username, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return user.Username, nil
}

func (*UserRepositoryImpl) Create(tx *gorm.DB, userReq payloads.CreateRequest, roleID int) (int, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		Username:   userReq.Username,
		Password:   userReq.Password,
		IsActive:   true,
		DateJoined: time.Now(),
		Email:      userReq.Email,
		RoleID:     roleID,
	}
	err := tx.
		Create(&user).
		Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return 0, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusConflict,
				Err:        err,
			}
		} else {
			return 0, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return user.ID, nil
}

func (*UserRepositoryImpl) Update(tx *gorm.DB, userReq payloads.CreateRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		Username:   userReq.Username,
		Password:   userReq.Password,
		IsActive:   userReq.IsActive,
		Email:      userReq.Email,
		OtpEnabled: true,
	}

	err := tx.
		Where(userID).
		Updates(&user).
		Error
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}
func (*UserRepositoryImpl) Delete(tx *gorm.DB, userID int) (bool, *exceptions.BaseErrorResponse) {

	err := tx.
		Where(userID).
		Delete(userID).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return true, nil
}
