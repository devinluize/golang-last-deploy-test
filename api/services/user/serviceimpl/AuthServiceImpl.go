package userserviceimpl

import (
	"errors"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"time"
	"user-services/api/config"
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/helper"
	"user-services/api/payloads"
	redisrepo "user-services/api/repositories/redis"
	userrepo "user-services/api/repositories/user"
	"user-services/api/securities"
	userservices "user-services/api/services/user"
	"user-services/api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/pquerna/otp/totp"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	DB              *gorm.DB
	DBRedis         *config.Database
	AuthRepository  userrepo.AuthRepository
	UserRepository  userrepo.UserRepository
	RedisRepository redisrepo.RedisRepository
	Validate        *validator.Validate
}

func NewAuthService(
	db *gorm.DB,
	dbRedis *config.Database,
	authRepository userrepo.AuthRepository,
	userRepository userrepo.UserRepository,
	redisRepository redisrepo.RedisRepository,
	validate *validator.Validate,
) userservices.AuthService {
	return &AuthServiceImpl{
		DB:              db,
		DBRedis:         dbRedis,
		AuthRepository:  authRepository,
		UserRepository:  userRepository,
		RedisRepository: redisRepository,
		Validate:        validate,
	}
}

// Login implements services.AuthService.
func (service *AuthServiceImpl) Login(loginReq payloads.LoginRequest) (masterentities.User, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetByUsername(tx, loginReq.Username)

	if get.Username == "" {
		return get, err
	}

	return get, nil
}

// LoginAuths(requestPayloads payloads.LoginRequestPayloads) (masterentities.UserEntities, *exceptions.BaseErrorResponse)
func (service *AuthServiceImpl) LoginAuth(requestPayloads payloads.LoginRequestPayloads) (masterentities.UserEntities, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.AuthRepository.LoginAuth(tx, requestPayloads)
	return get, err
	//if get.Username == "" {
	//	return get, err
	//}
	//
	//return get, nil
}

// Register implements services.AuthService.
func (service *AuthServiceImpl) Register(userReq payloads.RegisterRequest) (string, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.AuthRepository.RegisterUser(userReq, tx)
	return get, err
}

// CheckPasswordResetTime implements services.AuthService.
func (service *AuthServiceImpl) CheckPasswordResetTime(emailReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.AuthRepository.CheckPasswordResetTime(tx, emailReq)

	if err != nil {
		return get, err
	}
	return get, nil
}

// UpdatePassword implements services.AuthService.
func (service *AuthServiceImpl) UpdatePassword(claims *payloads.UserDetail, changePasswordRequest payloads.ChangePasswordInput) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	getUser, err := service.UserRepository.GetByID(
		tx,
		claims.UserID,
	)
	if err != nil {
		return false, err
	}

	hashPwd := getUser.Password
	pwd := changePasswordRequest.OldPassword

	_, errors := securities.VerifyPassword(hashPwd, pwd)
	if errors != nil {
		return false,
			&exceptions.BaseErrorResponse{
				StatusCode: http.StatusBadRequest,
				Err:        errors,
			}
	}

	pass, err := securities.HashPassword(changePasswordRequest.NewPassword)
	if err != nil {
		return false, err
	}

	update, err := service.AuthRepository.UpdatePassword(tx, pass, claims.UserID)

	if err != nil {
		return update, err
	}

	return update, nil
}

// UpdatePasswordByToken implements services.AuthService.
func (service *AuthServiceImpl) UpdatePasswordByToken(passReq payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	update, err := service.AuthRepository.UpdatePasswordByToken(tx, passReq)

	if err != nil {
		return update, err
	}

	return update, nil
}

// ResetPassword implements services.AuthService.
func (service *AuthServiceImpl) ResetPassword(resetToken string, passReq payloads.ResetPasswordInput) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, errors := securities.ComparePassword(passReq.Password, passReq.PasswordConfirm)

	if errors != nil {
		return false, &exceptions.BaseErrorResponse{
			Err: errors,
		}
	}
	hashedPassword, err := securities.HashPassword(passReq.Password)

	if err != nil {
		return false, err
	}
	passwordResetToken := utils.Encode(resetToken)

	_, err = service.AuthRepository.CheckPasswordResetTime(
		tx,
		payloads.UpdateEmailTokenRequest{
			PasswordResetToken: utils.StringPtr(passwordResetToken),
			PasswordResetAt:    utils.TimePtr(time.Now()),
		})
	if err != nil {
		return false, err
	}

	_, err = service.AuthRepository.UpdatePasswordByToken(
		tx,
		payloads.UpdatePasswordByTokenRequest{
			Password:           hashedPassword,
			PasswordResetToken: utils.StringPtr(passwordResetToken),
		})

	if err != nil {
		return false, err
	}
	update, err := service.AuthRepository.ResetPassword(tx, payloads.ResetPasswordRequest{
		PasswordResetToken: utils.StringPtr(passwordResetToken),
	})

	if err != nil {
		return update, err
	}

	return update, nil
}

// UpdatePasswordTokenByEmail implements services.AuthService.
func (service *AuthServiceImpl) UpdatePasswordTokenByEmail(emailReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	update, err := service.AuthRepository.UpdatePasswordTokenByEmail(tx, emailReq)

	if err != nil {
		return update, err
	}

	return update, nil
}

// UpdateUserOTP implements services.AuthService.
func (service *AuthServiceImpl) UpdateUserOTP(otpReq masterentities.OTPInput, remoteAddr string) (*payloads.ResponseAuth, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	txRedis := service.DBRedis
	defer helper.CommitOrRollback(tx)

	var loginCredential payloads.LoginCredential

	loginCredential.IpAddress = remoteAddr
	loginCredential.Client = otpReq.Client
	getUser, err := service.UserRepository.GetByID(
		tx,
		otpReq.UserID,
	)

	if err != nil {
		return nil, err
	}

	token, err := securities.GenerateToken(getUser.Username, getUser.ID, getUser.RoleID, getUser.CompanyID, loginCredential.Client)

	if err != nil {
		return nil, err
	}

	loginCredential.Session = token
	_, err = service.AuthRepository.UpdateCredential(
		tx,
		loginCredential,
		int(getUser.ID),
	)
	if err != nil {
		return nil, err
	}
	_, err = service.RedisRepository.UpdateCredential(txRedis, loginCredential, int(getUser.ID))
	if err != nil {
		return nil, err
	}

	valid := totp.Validate(otpReq.Token, getUser.OtpSecret)
	if !valid {
		return nil, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("token not valid"),
		}
	}

	updateOTP := payloads.OTPRequest{
		OtpVerified: true,
		OtpEnabled:  true,
	}
	response := payloads.ResponseAuth{
		Token:   token,
		Role:    getUser.RoleID,
		Company: getUser.CompanyID,
		UserID:  getUser.ID,
	}
	_, err = service.AuthRepository.UpdateUserOTP(tx, updateOTP, getUser.ID)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GenerateOTP implements services.AuthService.
func (service *AuthServiceImpl) GenerateOTP(userID int) (string, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.EnvConfigs.Issuer,
		AccountName: config.EnvConfigs.AccountName,
		SecretSize:  15,
	})

	if err != nil {
		return "", &exceptions.BaseErrorResponse{
			Err: err,
		}
	}

	updateSecretUrl := payloads.SecretUrlRequest{
		Secret: key.Secret(),
		Url:    key.URL(),
	}

	_, errors := service.AuthRepository.GenerateOTP(tx, updateSecretUrl, userID)

	if errors != nil {
		return "", errors
	}

	otpResponse := payloads.SecretUrlResponse{
		Secret: key.Secret(),
		Url:    key.URL(),
		UserID: userID,
	}

	img, _ := qrcode.NewQRCodeWriter().Encode(otpResponse.Url, gozxing.BarcodeFormat_QR_CODE, 250, 250, nil)
	fileName := fmt.Sprintf("%v-*.png", otpResponse.UserID)

	file, err := os.CreateTemp("", fileName)
	if err != nil {
		return "", &exceptions.BaseErrorResponse{
			Err: err,
		}
	}

	defer os.Remove(file.Name())

	_ = png.Encode(file, img)

	imgFile, err := os.Open(file.Name())
	if err != nil {
		return "", &exceptions.BaseErrorResponse{
			Err: err,
		}
	}
	defer imgFile.Close()

	return file.Name(), nil
}

// UpdateCredential implements services.AuthService.
func (service *AuthServiceImpl) UpdateCredential(loginReq payloads.LoginCredential, userID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	update, err := service.AuthRepository.UpdateCredential(tx, loginReq, userID)
	if err != nil {
		return update, err
	}

	return update, nil
}
