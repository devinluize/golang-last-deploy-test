package approvalrequestcontrollers

import (
	"net/http"
	"strconv"
	"user-services/api/exceptions"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	"user-services/api/payloads"
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
	"user-services/api/securities"
	approvalservices "user-services/api/services/master/approval"
	approverservices "user-services/api/services/master/approver"
	usergroupservices "user-services/api/services/master/user-group"
	approvalrequestservices "user-services/api/services/transaction/approval-request"
	"user-services/api/utils"
	"user-services/api/utils/validation"

	"github.com/go-chi/chi/v5"
)

type ApprovalRequestController interface {
	GetApprovalRequest(writer http.ResponseWriter, request *http.Request)
	GetApprovalRequestDetails(writer http.ResponseWriter, request *http.Request)
	CreateApprovalRequest(writer http.ResponseWriter, request *http.Request)
	CreateApprovalRequestDetail(writer http.ResponseWriter, request *http.Request)
}

type ApprovalRequestControllerImpl struct {
	ApprovalService        approvalservices.ApprovalService
	ApprovalMappingService approvalservices.ApprovalMappingService
	ApprovalLevelService   approvalservices.ApprovalLevelService
	ApprovalRequestService approvalrequestservices.ApprovalRequestService
	ApproverService        approverservices.ApproverService
	UserGroupService       usergroupservices.UserGroupService
}

// GetApprovalRequest implements ApprovalRequestController.
func (controller *ApprovalRequestControllerImpl) GetApprovalRequest(writer http.ResponseWriter, request *http.Request) {
	userType := chi.URLParam(request, "user_type")
	claims, _ := securities.ExtractAuthToken(request)
	query := request.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	page, _ := strconv.Atoi(query.Get("page"))
	approval, err := controller.ApprovalRequestService.GetByUserType(
		userType,
		claims.UserID,
		utils.Pagination{
			Limit: limit,
			Page:  page,
		})

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}
	payloads.HandleSuccess(writer, approval, "Get Data Approval Request Success", http.StatusOK)
}

// GetApprovalRequestDetailsForApprover implements ApprovalRequestController.
func (controller *ApprovalRequestControllerImpl) GetApprovalRequestDetails(writer http.ResponseWriter, request *http.Request) {
	claims, _ := securities.ExtractAuthToken(request)
	userType := chi.URLParam(request, "user_type")
	query := request.URL.Query()
	approvalID, _ := strconv.Atoi(query.Get("approval_id"))
	limit, _ := strconv.Atoi(query.Get("limit"))
	page, _ := strconv.Atoi(query.Get("page"))
	approval, err := controller.ApprovalRequestService.GetDetailsByUserType(
		userType,
		claims.UserID,
		approvalID,
		utils.Pagination{
			Limit: limit,
			Page:  page,
		})

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, approval, "Get Data Approval Request Details Success", http.StatusOK)
}

func (controller *ApprovalRequestControllerImpl) CreateApprovalRequest(writer http.ResponseWriter, request *http.Request) {
	approvalRequest := approvalrequestpayloads.CreateApprovalRequest{}
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
	claims, _ := securities.ExtractAuthToken(request)
	_, err = controller.ApprovalRequestService.Create(claims, approvalRequest)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, true, "Create Approval Request Success", http.StatusCreated)
}

func (controller *ApprovalRequestControllerImpl) CreateApprovalRequestDetail(writer http.ResponseWriter, request *http.Request) {
	approvalRequest := approvalrequestpayloads.UpdateApprovalRequestDetails{}
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
	claims, _ := securities.ExtractAuthToken(request)

	createDetails, err := controller.ApprovalRequestService.CreateDetails(claims, approvalRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, createDetails, "Create Approval Request Detail Success", http.StatusCreated)
}

func NewApprovalRequestController(
	approvalService approvalservices.ApprovalService,
	approvalRequestService approvalrequestservices.ApprovalRequestService,
	approverService approverservices.ApproverService,
	approvalMappingService approvalservices.ApprovalMappingService,
	approvalLevelService approvalservices.ApprovalLevelService,
	userGroupService usergroupservices.UserGroupService,
) ApprovalRequestController {
	return &ApprovalRequestControllerImpl{
		ApprovalService:        approvalService,
		ApprovalRequestService: approvalRequestService,
		ApproverService:        approverService,
		ApprovalMappingService: approvalMappingService,
		ApprovalLevelService:   approvalLevelService,
		UserGroupService:       userGroupService,
	}
}
