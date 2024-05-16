package approvercontrollers

import (
	"net/http"
	"strconv"
	"user-services/api/exceptions"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	"user-services/api/payloads"
	approverpayloads "user-services/api/payloads/master/approver"
	approverservices "user-services/api/services/master/approver"
	"user-services/api/utils"
	"user-services/api/utils/validation"

	"github.com/go-chi/chi/v5"
)

type ApproverController interface {
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

func NewApproverController(
	approverService approverservices.ApproverService,
) ApproverController {
	return &ApproverControllerImpl{
		ApproverService: approverService,
	}
}

type ApproverControllerImpl struct {
	ApproverService approverservices.ApproverService
}

func (controller *ApproverControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request) {
	get, err := controller.ApproverService.GetAll()
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}
	payloads.HandleSuccess(writer, get, "Get Approval Success", http.StatusCreated)
}

func (controller *ApproverControllerImpl) Get(writer http.ResponseWriter, request *http.Request) {
	approverID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	get, errors := controller.ApproverService.Get(approverID)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, get, "Get Approval Success", http.StatusCreated)
}

func (controller *ApproverControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	approverRequest := approverpayloads.CreateApprover{}

	err := jsonchecker.ReadFromRequestBody(request, &approverRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}
	err = validation.ValidationForm(writer, request, approverRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}
	create, err := controller.ApproverService.Create(approverRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}
	payloads.HandleSuccess(writer, create, "Create Approval Success", http.StatusCreated)
}

func (controller *ApproverControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	approverID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	approverRequest := approverpayloads.UpdateApprover{}
	errors := jsonchecker.ReadFromRequestBody(request, &approverRequest)
	if errors != nil {
		exceptions.NewEntityException(writer, request, errors)
		return
	}
	errors = validation.ValidationForm(writer, request, approverRequest)
	if errors != nil {
		exceptions.NewBadRequestException(writer, request, errors)
		return
	}
	update, errors := controller.ApproverService.Update(approverID, approverRequest)
	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, update, "Update Approval Success", http.StatusCreated)
}

func (controller *ApproverControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	menuAccessId, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
	}
	delete, errors := controller.ApproverService.Delete(int(menuAccessId))
	if errors != nil {
		helper.ReturnError(writer, request, errors)
	}
	payloads.HandleSuccess(writer, delete, utils.DeleteDataSuccess, http.StatusOK)
}
