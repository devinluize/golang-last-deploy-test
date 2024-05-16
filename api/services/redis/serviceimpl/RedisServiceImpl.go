package redisserviceimpl

import (
	"net/http"
	"user-services/api/config"
	"user-services/api/exceptions"
	"user-services/api/helper"
	"user-services/api/payloads"
	redisrepo "user-services/api/repositories/redis"
	userrepo "user-services/api/repositories/user"
	"user-services/api/securities"
	redisservices "user-services/api/services/redis"

	"gorm.io/gorm"
)

type RedisServiceImpl struct {
	DB              *gorm.DB
	DBRedis         *config.Database
	AuthRepository  userrepo.AuthRepository
	UserRepository  userrepo.UserRepository
	RedisRepository redisrepo.RedisRepository
}

func NewRedisService(
	db *gorm.DB,
	dbRedis *config.Database,
	authRepo userrepo.AuthRepository,
	userRepo userrepo.UserRepository,
	redisRepo redisrepo.RedisRepository,
) redisservices.RedisService {
	return &RedisServiceImpl{
		DB:              db,
		DBRedis:         dbRedis,
		AuthRepository:  authRepo,
		UserRepository:  userRepo,
		RedisRepository: redisRepo,
	}
}

// GetSession from Redis (Check if user login on another device)
func (service *RedisServiceImpl) GetSession(userID int) (string, *exceptions.BaseErrorResponse) {
	txRedis := service.DBRedis

	get, err := service.RedisRepository.GetSession(txRedis, userID)
	if err != nil {
		return get, err
	}

	return get, nil
}

func (service *RedisServiceImpl) Login(
	loginRequest payloads.LoginRequest,
	remoteAddr string,
) (
	*payloads.ResponseAuth,
	bool,
	*exceptions.BaseErrorResponse,
) {
	txRedis := service.DBRedis
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	response := payloads.ResponseAuth{}
	getUser, err := service.UserRepository.GetByUsername(
		tx,
		loginRequest.Username,
	)

	if err != nil {
		return &response, false, err
	}

	var loginCredential payloads.LoginCredential
	loginCredential.Client = loginRequest.Client
	loginCredential.IpAddress = remoteAddr
	hashPwd := getUser.Password
	pwd := loginRequest.Password

	_, errors := securities.VerifyPassword(hashPwd, pwd)
	if errors != nil {
		return &response, false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        errors,
		}
	}

	token, err := securities.GenerateToken(getUser.Username, getUser.ID, getUser.RoleID, getUser.CompanyID, loginRequest.Client)

	if err != nil {
		return &response, false, err
	}

	response = payloads.ResponseAuth{
		Token:   token,
		Role:    getUser.RoleID,
		Company: getUser.CompanyID,
		UserID:  getUser.ID,
	}

	session, err := service.RedisRepository.GetSession(
		txRedis,
		getUser.ID,
	)

	if err != nil {
		return &response, false, err
	}

	if getUser.OtpEnabled && loginCredential.Client != session {
		return &response, true, nil
	} else {
		_, err = service.AuthRepository.UpdateCredential(
			tx,
			loginCredential,
			int(getUser.ID),
		)
		if err != nil {
			return &response, false, err
		}

		_, err := service.RedisRepository.UpdateCredential(txRedis, loginCredential, getUser.ID)
		if err != nil {
			return &response, false, err
		}
	}

	return &response, false, nil
}

// DeleteCredential implements services.RedisService.
func (service *RedisServiceImpl) DeleteCredential(userID int) (bool, *exceptions.BaseErrorResponse) {
	txRedis := service.DBRedis
	delete, err := service.RedisRepository.DeleteCredential(txRedis, userID)

	if err != nil {
		return delete, err
	}

	return delete, nil
}
