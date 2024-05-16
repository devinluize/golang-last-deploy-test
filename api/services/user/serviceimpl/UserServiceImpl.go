package userserviceimpl

import (
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/helper"
	"user-services/api/payloads"
	userrepo "user-services/api/repositories/user"
	userservices "user-services/api/services/user"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository userrepo.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

// GetCurrentUser implements userservices.UserService.
func (service *UserServiceImpl) GetCurrentUser(userID int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserRepository.GetCurrentUser(tx, userID)

	if err != nil {
		return get, err
	}

	return get, nil
}

func NewUserService(userRepository userrepo.UserRepository, db *gorm.DB, validate *validator.Validate) userservices.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}
func (service *UserServiceImpl) CheckUserExists(username string) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserRepository.CheckUserExists(tx, username)

	if err != nil {
		return false, err
	}

	return get, nil
}

func (service *UserServiceImpl) FindUser(username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserRepository.FindUser(tx, username)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetEmails implements services.UserService.
func (service *UserServiceImpl) GetEmails(users []int) ([]string, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetEmails(tx, users)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetByEmail implements services.UserService.
func (service *UserServiceImpl) GetByEmail(email string) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetByEmail(tx, email)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetUserIDByUsername implements services.UserService.
func (service *UserServiceImpl) GetUserIDByUsername(username string) (int, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetUserIDByUsername(tx, username)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetUsernameByUserID implements services.UserService.
func (service *UserServiceImpl) GetUsernameByUserID(userID int) (string, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetUsernameByUserID(tx, userID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetByID implements services.UserService.
func (service *UserServiceImpl) GetByID(userID int) (masterentities.User, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserRepository.GetByID(tx, userID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetUser implements services.UserService.
func (service *UserServiceImpl) GetUser(username string) (masterentities.User, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetByUsername(tx, username)
	if err != nil {
		return get, err
	}
	return get, nil
}

// GetUserDetailByUsername implements services.UserService.
func (service *UserServiceImpl) GetUserDetailByUsername(username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetUserDetailByUsername(tx, username)

	if err != nil {
		return get, err
	}

	return get, nil
}

// ViewUser implements services.UserService.
func (service *UserServiceImpl) ViewUser() ([]masterentities.User, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.ViewUser(tx)

	if err != nil {
		return get, err
	}

	return get, nil
}

// UpdateUser implements services.UserService.
func (service *UserServiceImpl) UpdateUser(userReq payloads.CreateRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	update, err := service.UserRepository.Update(tx, userReq, userID)

	if err != nil {
		return update, err
	}

	return update, nil
}

// DeleteUser implements services.UserService.
func (service *UserServiceImpl) DeleteUser(userID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	deleteUser, err := service.UserRepository.Delete(tx, userID)

	if err != nil {
		return deleteUser, err
	}

	return deleteUser, nil
}
