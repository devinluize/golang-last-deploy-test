package menucontrollers

import (
	"net/http"
	"strconv"
	"user-services/api/exceptions"
	"user-services/api/helper"
	payloads "user-services/api/payloads"
	menupayloads "user-services/api/payloads/master/menu"
	"user-services/api/securities"
	menuservices "user-services/api/services/master/menu"
	"user-services/api/utils"

	"github.com/go-chi/chi/v5"
)

type MenuAccessController interface {
	Get(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	CheckByFilter(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

type MenuAccessControllerImpl struct {
	MenuAccessService menuservices.MenuAccessService
}

func NewMenuAccessController(
	menuAccessService menuservices.MenuAccessService,
) MenuAccessController {
	return &MenuAccessControllerImpl{
		MenuAccessService: menuAccessService,
	}
}

func (controller *MenuAccessControllerImpl) Get(writer http.ResponseWriter, request *http.Request) {
	companyID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	claims, _ := securities.ExtractAuthToken(request)

	getParent, errors := controller.MenuAccessService.GetByCompanyAndUserID(companyID, claims.UserID)

	if errors != nil {
		exceptions.NewNotFoundException(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, getParent, utils.GetDataSuccess, http.StatusOK)
}

func (controller *MenuAccessControllerImpl) CheckByFilter(writer http.ResponseWriter, request *http.Request) {
	menuListId, err := strconv.Atoi(chi.URLParam(request, "menu_list_id"))

	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		})
		return
	}

	claims, _ := securities.ExtractAuthToken(request)

	menuAccessRequest := menupayloads.CreateMenuAccessParamRequest{
		RoleID:     claims.Role,
		MenuListID: []int{int(menuListId)},
	}

	get, errors := controller.MenuAccessService.IsUserHaveAccess(menuAccessRequest)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, get, utils.GetDataSuccess, http.StatusOK)
}

func (controller *MenuAccessControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {

	var menuAccessRequest menupayloads.CreateMenuAccessParamRequest

	create, err := controller.MenuAccessService.Create(menuAccessRequest)

	if err != nil {
		helper.ReturnError(writer, request, err)
	}

	payloads.HandleSuccess(writer, create, utils.CreateDataSuccess, http.StatusCreated)
}

func (controller *MenuAccessControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	menuAccessId, err := strconv.Atoi(chi.URLParam(request, "menu_access_id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
	}
	delete, errors := controller.MenuAccessService.Delete(int(menuAccessId))
	if errors != nil {
		helper.ReturnError(writer, request, errors)
	}
	payloads.HandleSuccess(writer, delete, utils.DeleteDataSuccess, http.StatusOK)
}
