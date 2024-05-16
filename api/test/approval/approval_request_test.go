package approvaltest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"user-services/api/config"
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
	"user-services/api/utils"

	"github.com/zeebo/assert"
)

func TestCreateApprovalRequest(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	approvalRequestRouter := SetupRouter(db)
	myStruct := approvalrequestpayloads.CreateApprovalRequest{
		ApprovalMappingID: 1,
		ApprovalAmountID:  1,
		CompanyID:         1,
		ModuleID:          1,
		DocumentTypeID:    1,
		RequestDate:       time.Now(),
		SourceSysNo:       1,
		SourceAmount:      1,
		TransactionTypeID: 1,
		BrandID:           1,
		ProfitCenterID:    1,
		CostCenterID:      1,
		StatusID:          1,
		IsVoid:            false,
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}

	reader := strings.NewReader(string(jsonData))
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/v1/approval-request", reader)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()

	approvalRequestRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 201, int(responseBody["status_code"].(float64)))
}

func TestCreateApprovalRequestDetails(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	approvalRequestRouter := SetupRouter(db)

	myStruct := approvalrequestpayloads.UpdateApprovalRequestDetails{
		ApprovalRequestID: 25,
		// AllowApprove:      true,
		StatusID:  1,
		Remark:    utils.StringPtr("test"),
		CompanyID: 1,
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}

	reader := strings.NewReader(string(jsonData))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/v1/approval-request/detail", reader)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()

	approvalRequestRouter.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 201, int(responseBody["status_code"].(float64)))
}

func TestGetApprovalRequestByApprover(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	approvalRequestRouter := SetupRouter(db)
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/v1/approval-request/approver", nil)
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

func TestGetApprovalRequestByRequester(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	approvalRequestRouter := SetupRouter(db)
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/v1/approval-request/requester", nil)
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

func TestGetApprovalRequestDetailByApprover(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	approvalRequestRouter := SetupRouter(db)
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/v1/approval-request/approver", nil)
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
