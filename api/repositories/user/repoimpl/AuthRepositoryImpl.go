package userrepoimpl

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/helper"
	"user-services/api/payloads"
	userrepo "user-services/api/repositories/user"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() userrepo.AuthRepository {
	return &AuthRepositoryImpl{}
}

// CheckPasswordResetTime implements repositories.AuthRepository.
// LoginAuth(*gorm.DB, payloads.LoginRequestPayloads) (masterentities.UserEntities, *exceptions.BaseErrorResponse)
func (*AuthRepositoryImpl) LoginAuth(tx *gorm.DB, req payloads.LoginRequestPayloads) (masterentities.UserEntities, *exceptions.BaseErrorResponse) {
	var user masterentities.UserEntities
	err := tx.Where("user_name =?", req.UserName).First(&user).Error
	if err != nil {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Err:        err,
			Message:    "Username Or Password Wrong",
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Err:        err,
			Message:    "Password salah",
		}

	}
	return user, nil
}
func (*AuthRepositoryImpl) CheckPasswordResetTime(tx *gorm.DB, tokenReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	var exists bool
	err := tx.Model(masterentities.User{}).
		Select("count(company_id)").
		Where(
			"password_reset_token = ? AND password_reset_at > ?",
			tokenReq.PasswordResetToken, tokenReq.PasswordResetAt,
		).
		Find(&exists).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if !exists {
		return exists, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Invalid Token"),
		}
	}
	return exists, nil
}

// UpdateCredential implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdateCredential(tx *gorm.DB, loginReq payloads.LoginCredential, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		IpAddress: loginReq.IpAddress,
		LastLogin: time.Now(),
		ID:        userID,
	}

	err := tx.
		Where(masterentities.User{ID: userID}).
		Updates(&user).Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return true, nil
}
func (repo *AuthRepositoryImpl) RegisterUser(request payloads.RegisterRequest, db *gorm.DB) (string, *exceptions.BaseErrorResponse) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	//tx := db.Begin()
	//defer tx.Rollback()
	request.Password = string(hashPassword)
	user := helper.ToDomainRegister(request)
	err := db.Create(user)
	if err == nil {
		db.Commit()
	}
	if err.Error != nil {

		return "Register To Database Gagal Bad Responses", &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err.Error,
		}
	}
	return "Register Success", nil
}

// UpdatePassword implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdatePassword(tx *gorm.DB, password string, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		Password: password,
	}
	row, err := tx.
		Where(masterentities.User{ID: userID}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// UpdatePasswordByToken implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdatePasswordByToken(tx *gorm.DB, passReq payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		Password: passReq.Password,
	}
	row, err := tx.
		Where(masterentities.User{PasswordResetToken: passReq.PasswordResetToken}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// ResetPassword implements repositories.AuthRepository.
func (*AuthRepositoryImpl) ResetPassword(tx *gorm.DB, updateReq payloads.ResetPasswordRequest) (bool, *exceptions.BaseErrorResponse) {
	var user masterentities.User
	var nullString *string
	var nullTime *time.Time
	row, err := tx.
		Model(&user).
		Where(masterentities.User{PasswordResetToken: updateReq.PasswordResetToken}).
		Updates(map[string]interface{}{
			"password_reset_token": nullString,
			"password_reset_at":    nullTime,
		}).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// UpdatePasswordTokenByEmail implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdatePasswordTokenByEmail(tx *gorm.DB, emailReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		PasswordResetToken: emailReq.PasswordResetToken,
		PasswordResetAt:    emailReq.PasswordResetAt,
	}

	row, err := tx.
		Where(masterentities.User{Email: emailReq.Email}).
		Updates(&user).
		Rows()

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// UpdateUserOTP implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdateUserOTP(tx *gorm.DB, otpReq payloads.OTPRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		OtpVerified: otpReq.OtpVerified,
		OtpEnabled:  otpReq.OtpEnabled,
	}

	row, err := tx.
		Where(masterentities.User{ID: userID}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// GenerateOTP implements repositories.AuthRepository.
func (*AuthRepositoryImpl) GenerateOTP(tx *gorm.DB, userReq payloads.SecretUrlRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := masterentities.User{
		OtpSecret:  userReq.Secret,
		OtpAuthUrl: userReq.Url,
	}

	row, err := tx.
		Where(masterentities.User{ID: userID}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}
