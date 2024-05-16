package approvalcontrollers

import (
	"net/http"
	"strconv"
	"user-services/api/exceptions"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	"user-services/api/payloads"
	approvalpayloads "user-services/api/payloads/master/approval"
	approvalservices "user-services/api/services/master/approval"
	"user-services/api/utils"
	"user-services/api/utils/validation"

	"github.com/go-chi/chi/v5"
)

type ApprovalController interface {
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

func NewApprovalController(
	approvalService approvalservices.ApprovalService,
	approvalMappingService approvalservices.ApprovalMappingService,
	approvalLevelService approvalservices.ApprovalLevelService,
) ApprovalController {
	return &ApprovalControllerImpl{
		ApprovalService:        approvalService,
		ApprovalMappingService: approvalMappingService,
		ApprovalLevelService:   approvalLevelService,
	}
}

type ApprovalControllerImpl struct {
	ApprovalService        approvalservices.ApprovalService
	ApprovalMappingService approvalservices.ApprovalMappingService
	ApprovalLevelService   approvalservices.ApprovalLevelService
}

func (controller *ApprovalControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	page, _ := strconv.Atoi(query.Get("page"))

	get, err := controller.ApprovalService.GetAll(
		utils.Pagination{
			Limit: limit,
			Page:  page,
		})

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, get, "Get Approval Success", http.StatusOK)
}
func (controller *ApprovalControllerImpl) Get(writer http.ResponseWriter, request *http.Request) {

	approvalID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	get, errors := controller.ApprovalService.Get(approvalID)
	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}

	payloads.HandleSuccess(writer, get, "Get Approval Success", http.StatusOK)
}
func (controller *ApprovalControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	approvalRequest := approvalpayloads.CreateApproval{}
	err := jsonchecker.ReadFromRequestBody(request, &approvalRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}
	err = validation.ValidationForm(writer, request, approvalRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}

	create, err := controller.ApprovalService.Create(approvalRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}
	payloads.HandleSuccess(writer, create, "Create Approval Success", http.StatusCreated)
}

func (controller *ApprovalControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	approvalID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	approvalRequest := approvalpayloads.UpdateApproval{}

	errors := jsonchecker.ReadFromRequestBody(request, &approvalRequest)
	if errors != nil {
		exceptions.NewEntityException(writer, request, errors)
		return
	}
	errors = validation.ValidationForm(writer, request, approvalRequest)
	if errors != nil {
		exceptions.NewBadRequestException(writer, request, errors)
		return
	}
	update, errors := controller.ApprovalService.Update(approvalID, approvalRequest)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, update, "Update Approval Success", http.StatusOK)
}

func (controller *ApprovalControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	approvalID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	delete, errors := controller.ApprovalService.Delete(int(approvalID))
	if errors != nil {
		helper.ReturnError(writer, request, errors)
	}
	payloads.HandleSuccess(writer, delete, utils.DeleteDataSuccess, http.StatusOK)
}
