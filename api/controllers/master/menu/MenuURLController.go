package menucontrollers

import (
	"net/http"
	"strconv"
	"user-services/api/exceptions"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	payloads "user-services/api/payloads"
	menupayloads "user-services/api/payloads/master/menu"
	"user-services/api/securities"
	menuservices "user-services/api/services/master/menu"
	"user-services/api/utils"
	"user-services/api/utils/validation"

	"github.com/go-chi/chi/v5"
)

type MenuURLController interface {
	GetByCompanyAndUser(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
}

type MenuURLControllerImpl struct {
	MenuAccessService menuservices.MenuAccessService
	MenuListService   menuservices.MenuListService
	MenuUrlService    menuservices.MenuUrlService
}

func NewMenuURLController(
	menuUrlService menuservices.MenuUrlService,
) MenuURLController {
	return &MenuURLControllerImpl{
		MenuUrlService: menuUrlService,
	}
}

func (controller *MenuURLControllerImpl) GetByCompanyAndUser(writer http.ResponseWriter, request *http.Request) {
	companyID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	claims, _ := securities.ExtractAuthToken(request)

	getParent, errors := controller.MenuUrlService.GetByCompanyAndUser(companyID, claims.UserID)
	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}

	payloads.HandleSuccess(writer, getParent, utils.GetDataSuccess, http.StatusOK)
}

func (controller *MenuURLControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	menuUrlRequest := []menupayloads.CreateMenuUrlRequest{}
	err := jsonchecker.ReadFromRequestBody(request, &menuUrlRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
	}
	err = validation.ValidationForm(writer, request, menuUrlRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
	}

	_, err = controller.MenuUrlService.Create(menuUrlRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, true, utils.CreateDataSuccess, http.StatusCreated)
}
