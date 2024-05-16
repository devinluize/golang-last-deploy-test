package menucontrollers

import (
	"net/http"
	"user-services/api/exceptions"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	payloads "user-services/api/payloads"
	menupayloads "user-services/api/payloads/master/menu"
	"user-services/api/securities"
	menuservices "user-services/api/services/master/menu"
	"user-services/api/utils"
	"user-services/api/utils/validation"
)

type MenuListController interface {
	GetDropdown(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
}

type MenuListControllerImpl struct {
	MenuAccessService menuservices.MenuAccessService
	MenuListService   menuservices.MenuListService
	MenuUrlService    menuservices.MenuUrlService
}

func NewMenuListController(
	menuListService menuservices.MenuListService,
) MenuListController {
	return &MenuListControllerImpl{
		MenuListService: menuListService,
	}
}

func (controller *MenuListControllerImpl) GetDropdown(writer http.ResponseWriter, request *http.Request) {

	get, err := controller.MenuListService.GetDropDown()

	if err != nil {
		exceptions.NewNotFoundException(writer, request, err)
		return
	}
	payloads.HandleSuccess(writer, get, utils.GetDataSuccess, http.StatusOK)
}

func (controller *MenuListControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	menuUrlRequest := []menupayloads.CreateMenuListRequest{}
	err := jsonchecker.ReadFromRequestBody(request, &menuUrlRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}

	err = validation.ValidationForm(writer, request, menuUrlRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}
	claims, _ := securities.ExtractAuthToken(request)

	create, err := controller.MenuListService.Create(claims.Role, menuUrlRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, create, utils.CreateDataSuccess, http.StatusCreated)
}
