package approvaltest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"user-services/api/config"
	approvalrequestcontrollers "user-services/api/controllers/transaction"
	masterentities "user-services/api/entities/master"
	approvalentities "user-services/api/entities/master/approval"
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
	approverpayloads "user-services/api/payloads/master/approver"
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
	approvalrepoimpl "user-services/api/repositories/master/approval/repoimpl"
	approverrepoimpl "user-services/api/repositories/master/approver/repoimpl"
	usergrouprepoimpl "user-services/api/repositories/master/user-group/repoimpl"
	approvalrequestrepoimpl "user-services/api/repositories/transaction/repoimpl"
	"user-services/api/route"
	approvalservicesimpl "user-services/api/services/master/approval/serviceimpl"
	approverservicesimpl "user-services/api/services/master/approver/serviceimpl"
	usergroupservicesimpl "user-services/api/services/master/user-group/serviceimpl"
	approvalrequestservicesimpl "user-services/api/services/transaction/approval-request/serviceimpl"
	"user-services/api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/zeebo/assert"
	"gorm.io/gorm"
)

var token string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOiJBIiwiY29tcGFueV9pZCI6MTUxLCJleHAiOjE3MTE2MDY1NzksInJvbGUiOjEsInVzZXJfaWQiOjIzMywidXNlcm5hbWUiOiI1MDU5NiJ9.SVEnVz7xwCPkvp4JOMorxgYUsp24lUu5FxR9IAjCUqE"

func SetupRouter(db *gorm.DB) chi.Router {
	config.InitEnvConfigs(true, "")

	userGroupRepository := usergrouprepoimpl.NewUserGroupRepository()
	userGroupService := usergroupservicesimpl.NewUserGroupService(userGroupRepository, db)

	approvalRepository := approvalrepoimpl.NewApprovalRepository()
	approvalService := approvalservicesimpl.NewApprovalService(approvalRepository, db)

	approverRepository := approverrepoimpl.NewApproverRepository()
	approverService := approverservicesimpl.NewApproverService(approverRepository, db)

	approvalMappingRepository := approvalrepoimpl.NewApprovalMappingRepository()
	approvalMappingService := approvalservicesimpl.NewApprovalMappingService(approvalMappingRepository, db)

	approvalLevelRepository := approvalrepoimpl.NewApprovalLevelRepository()
	approvalLevelService := approvalservicesimpl.NewApprovalLevelService(approvalLevelRepository, db)

	approvalRequestRepository := approvalrequestrepoimpl.NewApprovalRequestRepository()
	approvalRequestService := approvalrequestservicesimpl.NewApprovalRequestService(approvalRequestRepository, approverRepository, approvalLevelRepository, approvalMappingRepository, userGroupRepository, db)

	approvalRequestController := approvalrequestcontrollers.NewApprovalRequestController(approvalService, approvalRequestService, approverService, approvalMappingService, approvalLevelService, userGroupService)

	approvalRequestRouter := route.ApprovalRequestRouter(approvalRequestController)

	return approvalRequestRouter
}

func TestCheckApprovalLevelByUserID(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	var users masterentities.User
	var level int

	err := db.
		Model(&users).
		Select(
			"level",
		).
		Joins("ApproverDetails", db.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver", db.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver.ApprovalApprover",
			db.Select(
				"1",
			)).
		Joins("ApproverDetails.Approver.ApprovalApprover.ApprovalLevel",
			db.Select(
				"1",
			)).
		Where("users.id = ? and approval_amount_id = ? and approval_mapping_id = ?", 2, 1, 1).
		Scan(&level).
		Error

	if err != nil {
		fmt.Println(level)
	}

	fmt.Println(level)
}

// Test CheckCreateWithMenuListAndRole User
func TestGetApproval(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	var approvalMappingEntity approvalentities.ApprovalMapping
	var approvalResponse approvalpayloads.GetApprovalMappingResponse

	err := db.
		Preload("ApprovalAmount", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id",
				"approval_mapping_id",
			).
				Where("max_amount >= ?", 10000).
				Order("approval_amount.max_amount asc").
				Limit(1)
		}).
		Select(
			"id",
			"approval_id",
		).
		Where(approvalentities.ApprovalMapping{
			BrandID:           1,
			CompanyID:         1,
			TransactionTypeID: 1,
			ProfitCenterID:    1,
			CostCenterID:      1,
		}).
		Find(&approvalMappingEntity).
		Error

	approvalResponse = approvalpayloads.GetApprovalMappingResponse{
		ApprovalMappingID: approvalMappingEntity.ID,
		ApprovalAmountID:  approvalMappingEntity.ApprovalAmount[0].ID,
	}

	if err != nil {
		fmt.Println(approvalResponse, err)
	}

	fmt.Println(approvalResponse)
}

// Test CheckCreateWithMenuListAndRole User
func TestGetAllApproval(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	var approval []approvalentities.Approval
	pagination := utils.Pagination{
		Limit: 10,
		Page:  0,
	}

	err := db.
		Model(&approval).
		Scopes(utils.Paginate(approval, &pagination, db.Model(&approval))).
		Scan(&approval).
		Error

	if err != nil {
		fmt.Println(pagination, err)
	}

	fmt.Println(pagination)
}

// Test CheckCreateWithMenuListAndRole User
func TestGetEmailsFromApprover(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	var users []masterentities.User
	var approverResponse []approverpayloads.GetApproverForApprovalResponse

	err := db.
		Model(&users).
		Select(
			"ApproverDetails__Approver__ApprovalApprover.approver_id",
			"ApproverDetails.id approver_detail_id",
			"ApproverDetails__Approver__ApprovalApprover.id approval_approver_id",
			"level",
			"users.id user_id",
			"email",
		).
		Joins("ApproverDetails", db.Select(
			"id",
		)).
		Joins("ApproverDetails.Approver", db.Select(
			"id",
		)).
		Joins("ApproverDetails.Approver.ApprovalApprover",
			db.Select(
				"approval_approver_id",
				"level",
			)).
		Joins("ApproverDetails.Approver.ApprovalApprover.ApprovalLevel",
			db.Select(
				"approval_approver_id",
				"level",
			)).
		Where("level = ? and approval_amount_id = ? and approval_mapping_id = ?", 1, 1, 1).
		Scan(&approverResponse).
		Error

	if err != nil {
		fmt.Println(approverResponse)
	}

	fmt.Println(approverResponse)
}

// Test CheckCreateWithMenuListAndRole User
func TestGetDetailsByApprover(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	approvalMapping := []approvalentities.ApprovalMapping{}
	approvalResponse := []approvalrequestpayloads.GetApprovalRequestDetailsResponse{}
	pagination := utils.Pagination{}

	err := db.
		Model(&approvalMapping).
		Select(
			"ApprovalRequest.ID approval_request_id",
			"approval_mapping.company_id company_id",
			"ApprovalRequest.source_doc_no source_doc_no",
			"ApprovalRequest.request_date request_date",
		).
		Joins("Approval",
			db.Select("1").
				Where("Approval.id = ?", 1),
		).
		Joins("ApprovalRequest",
			db.Select("1"),
		).
		Scan(&approvalResponse).
		Error

	if err != nil {
		fmt.Println(pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
		)
	}

	companies := []int{}
	for _, company := range approvalResponse {
		companies = append(companies, company.CompanyID)
	}
	companies = utils.DistinctIntegers(companies)

	companiesString := []string{}
	for _, companyID := range companies {
		companiesString = append(companiesString, strconv.Itoa(companyID))
	}

	apiResponse, err := utils.CallExternalAPI(
		fmt.Sprintf("%s%s%s",
			config.EnvConfigs.GeneralAPI,
			"company-id/", strings.Join(companiesString, ",")),
		http.MethodGet,
		nil,
		"",
	)
	if err != nil {
		fmt.Printf("Error calling API: %v\n", err)
	}

	fmt.Println(apiResponse)
	pagination.Rows = approvalResponse
	fmt.Println(pagination, nil)
}

func TestGetApprovalurl(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	approvalRequestRouter := SetupRouter(db)
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/v1/approval/1", nil)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()

	approvalRequestRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
}
