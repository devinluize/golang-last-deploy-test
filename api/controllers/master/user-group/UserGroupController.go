package usergroupcontrollers

import (
	"net/http"
	"strconv"
	"user-services/api/exceptions"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	"user-services/api/payloads"
	usergrouppayloads "user-services/api/payloads/master/user-group"
	usergroupservices "user-services/api/services/master/user-group"
	"user-services/api/utils"
	"user-services/api/utils/validation"

	"github.com/go-chi/chi/v5"
)

type UserGroupController interface {
	Get(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

type UserGroupControllerImpl struct {
	UserGroupService usergroupservices.UserGroupService
}

// CreateUserGroup implements UserGroupController.
func (controller *UserGroupControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	userGroupRequest := usergrouppayloads.CreateUserGroupRequest{}
	err := jsonchecker.ReadFromRequestBody(request, &userGroupRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
	}
	err = validation.ValidationForm(writer, request, userGroupRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
	}

	_, err = controller.UserGroupService.Create(userGroupRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, true, utils.CreateDataSuccess, http.StatusCreated)
}

// GetUserGroup implements UserGroupController.
func (*UserGroupControllerImpl) Get(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// UpdateUserGroup implements UserGroupController.
func (controller *UserGroupControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	userGroupID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	userGroupRequest := usergrouppayloads.UpdateUserGroupRequest{}

	errors := jsonchecker.ReadFromRequestBody(request, &userGroupRequest)
	if errors != nil {
		exceptions.NewEntityException(writer, request, errors)
		return
	}
	errors = validation.ValidationForm(writer, request, userGroupRequest)
	if errors != nil {
		exceptions.NewBadRequestException(writer, request, errors)
		return
	}
	update, errors := controller.UserGroupService.Update(userGroupID, userGroupRequest)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, update, "Update Approval Success", http.StatusOK)
}

// UpdateUserGroup implements UserGroupController.
func (*UserGroupControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

func NewUserGroupController(userGroupService usergroupservices.UserGroupService) UserGroupController {
	return &UserGroupControllerImpl{
		UserGroupService: userGroupService,
	}
}
