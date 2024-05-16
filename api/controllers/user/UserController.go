package usercontrollers

import (
	"net/http"
	"strconv"
	"user-services/api/exceptions"
	"user-services/api/helper"
	"user-services/api/payloads"
	"user-services/api/securities"
	userservices "user-services/api/services/user"
	"user-services/api/utils"

	"github.com/go-chi/chi/v5"
)

type UserController interface {
	GetCurrentUser(writer http.ResponseWriter, request *http.Request)
	GetUser(writer http.ResponseWriter, request *http.Request)
	GetUserIDByUsername(writer http.ResponseWriter, request *http.Request)
	GetUsernameByUserID(writer http.ResponseWriter, request *http.Request)
	FindUser(writer http.ResponseWriter, request *http.Request)
}

type UserControllerImpl struct {
	UserService userservices.UserService
}

// GetCurrentUser implements UserController.
func (controller *UserControllerImpl) GetCurrentUser(writer http.ResponseWriter, request *http.Request) {
	claims, _ := securities.ExtractAuthToken(request)

	userResponse, err := controller.UserService.GetCurrentUser(claims.UserID)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, utils.GetDataSuccess, http.StatusOK)
}

func NewUserController(userService userservices.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// GetUser implements UserController.
func (controller *UserControllerImpl) GetUser(writer http.ResponseWriter, request *http.Request) {
	claims, _ := securities.ExtractAuthToken(request)

	userResponse, err := controller.UserService.GetUser(claims.Username)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, utils.GetDataSuccess, http.StatusOK)
}

//	@Summary		Find User By ID
//	@Description	REST API User
//	@Accept			json
//	@Produce		json
//	@Tags			User Controller
//	@Security		BearerAuth
//	@Success		200	{object}	payloads.Respons
//	@Router			/user/{user_id} [get]
func (controller *UserControllerImpl) GetUsernameByUserID(writer http.ResponseWriter, request *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(request, "user_id"))
	if err != nil {
		exceptions.NewAppException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	userResponse, errors := controller.UserService.GetUsernameByUserID(userID)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, userResponse, utils.GetDataSuccess, http.StatusOK)
}

//	@Summary		Find User By ID
//	@Description	REST API User
//	@Accept			json
//	@Produce		json
//	@Tags			User Controller
//	@Security		BearerAuth
//	@Success		200	{object}	payloads.Respons
//	@Router			/user/username/{username} [get]
func (controller *UserControllerImpl) GetUserIDByUsername(writer http.ResponseWriter, request *http.Request) {
	username := chi.URLParam(request, "username")
	userResponse, err := controller.UserService.GetUserIDByUsername(username)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, utils.GetDataSuccess, http.StatusOK)
}

//	@Summary		Find User
//	@Description	REST API User
//	@Accept			json
//	@Produce		json
//	@Tags			User Controller
//	@Security		BearerAuth
//	@Success		200	{object}	payloads.Respons
//	@Router			/user [get]

func (controller *UserControllerImpl) FindUser(writer http.ResponseWriter, request *http.Request) {
	claims, _ := securities.ExtractAuthToken(request)

	userResponse, err := controller.UserService.FindUser(claims.Username)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, utils.GetDataSuccess, http.StatusOK)
}
