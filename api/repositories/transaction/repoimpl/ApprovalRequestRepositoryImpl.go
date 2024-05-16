package approvalrequestrepoimpl

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"user-services/api/config"
	approvalentities "user-services/api/entities/master/approval"
	approvalrequestentities "user-services/api/entities/transaction/approval-request"
	"user-services/api/exceptions"
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
	approvalrequestrepo "user-services/api/repositories/transaction"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type ApprovalRequestRepositoryImpl struct {
}

// GetByApprover implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) GetByApprover(tx *gorm.DB, userID int, pagination utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	approvalMapping := []approvalentities.ApprovalMapping{}
	approvalResponse := []approvalrequestpayloads.GetApprovalRequestForRequesterResponse{}
	query := tx.
		Model(&approvalMapping).
		Select(
			"approval.id approval_id",
			"approval_mapping.id approval_mapping_id",
			"approval.name approval_name",
			"count(ApprovalRequest.id) count",
		).
		Joins("ApprovalRequest", tx.Select(
			"1",
		)).
		Joins("ApprovalRequest.ApprovalRequestDetails", tx.Select(
			"1",
		).
			Where("user_id = ?", userID)).
		Joins("Approval", tx.Select(
			"1",
		)).
		Group("approval.id,approval_mapping.id, approval.name")

	err := query.
		Scopes(utils.Paginate(&approvalMapping, &pagination, query)).
		Order("approval.name").
		Scan(&approvalResponse).
		Error

	if err != nil {
		return pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	pagination.Rows = approvalResponse
	return pagination, nil
}

// GetDetailsByApprover implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) GetDetailsByApprover(tx *gorm.DB, userID int, approvalID int, pagination utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	approvalMapping := []approvalentities.ApprovalMapping{}
	approvalResponse := []approvalrequestpayloads.GetApprovalRequestDetailsResponse{}

	apiResponse, err := utils.CallExternalAPI(fmt.Sprintf("%s%s", config.EnvConfigs.GeneralAPI, "company?page=0&limit=10"), http.MethodPost, nil, "")
	if err != nil {
		fmt.Printf("Error calling API: %v\n", err)
		return pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	fmt.Println(apiResponse)

	err = tx.
		Model(&approvalMapping).
		Select(
			"ApprovalRequest.ID approval_request_id",
			"ApprovalMapping.CompanyID company_id",
			"ApprovalRequest.source_date document_date",
			"ApprovalRequest.source_doc_no source_doc_no",
		).
		Joins("Approval",
			tx.Select("1").
				Where("id = ?", approvalID),
		).
		Joins("ApprovalRequest").
		Scan(&approvalResponse).
		Error

	if err != nil {
		return pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	pagination.Rows = approvalResponse
	return pagination, nil
}

// GetByUserID implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) GetByRequester(tx *gorm.DB, userID int, pagination utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	approvalMapping := []approvalentities.ApprovalMapping{}
	approvalResponse := []approvalrequestpayloads.GetApprovalRequestForRequesterResponse{}
	query := tx.
		Model(&approvalMapping).
		Select(
			"approval.id approval_id",
			"approval_mapping.id approval_mapping_id",
			"approval.name approval_name",
			"count(ApprovalRequest.id) count",
		).
		Joins("ApprovalRequest", tx.Select(
			"1",
		).
			Where("request_by = ?", userID)).
		Joins("Approval", tx.Select(
			"1",
		)).
		Group("approval.id,approval_mapping.id, approval.name")

	err := query.
		Scopes(utils.Paginate(&approvalMapping, &pagination, query)).
		Order("approval_mapping.id").
		Scan(&approvalResponse).
		Error

	if err != nil {

		return pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	pagination.Rows = approvalResponse
	return pagination, nil
}

// GetDetailsByRequester implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) GetDetailsByRequester(tx *gorm.DB, userID int, approvalID int, pagination utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	approvalMapping := []approvalentities.ApprovalMapping{}
	approvalResponse := []approvalrequestpayloads.GetApprovalRequestDetailsResponse{}
	err := tx.
		Model(&approvalMapping).
		Select(
			"ApprovalRequest.ID approval_request_id",
			"ApprovalMapping.CompanyID company_id",
			"ApprovalRequest.SourceDocNo source_doc_no",
		).
		Joins("Approval",
			tx.Select("1").
				Where("id = ?", approvalID),
		).
		Joins("ApprovalRequest",
			tx.Select("1").
				Where("request_by = ?", userID),
		).
		Scan(&approvalResponse).
		Error

	if err != nil {
		return pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	pagination.Rows = approvalResponse
	return pagination, nil
}

// CreateDetails implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) CreateDetails(tx *gorm.DB, approvalRequestDetails []approvalrequestpayloads.CreateApprovalRequestDetails) (bool, *exceptions.BaseErrorResponse) {
	approvalRequestDetailEntities := []approvalrequestentities.ApprovalRequestDetails{}

	for _, approvalRequestDetail := range approvalRequestDetails {
		approvalRequestDetailEntities = append(approvalRequestDetailEntities, approvalrequestentities.ApprovalRequestDetails{
			ApprovalRequestID:  approvalRequestDetail.ApprovalRequestID,
			ApprovalApproverID: approvalRequestDetail.ApprovalApproverID,
			ApproverID:         approvalRequestDetail.ApproverID,
			Level:              approvalRequestDetail.Level,
			UserID:             approvalRequestDetail.UserId,
			StatusID:           approvalRequestDetail.StatusID,
			Remark:             approvalRequestDetail.Remark,
			// AllowApprove:      approvalRequestDetail.AllowApprove,
		},
		)
	}

	err := tx.
		Create(&approvalRequestDetailEntities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

// UpdateDetails implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) UpdateDetails(tx *gorm.DB, approvalRequestDetails approvalrequestpayloads.UpdateApprovalRequestDetails) (bool, *exceptions.BaseErrorResponse) {
	approvalrequestentities := approvalrequestentities.ApprovalRequestDetails{
		StatusID: approvalRequestDetails.StatusID,
		Remark:   approvalRequestDetails.Remark,
	}

	err := tx.
		Updates(&approvalrequestentities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

// Get implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) Get(tx *gorm.DB, approvalRequestID int) (approvalrequestpayloads.GetApprovalRequest, *exceptions.BaseErrorResponse) {
	var approvalRequest approvalrequestentities.ApprovalRequest
	var approvalResponse approvalrequestpayloads.GetApprovalRequest
	var approvalRequestDetails []approvalrequestpayloads.CreateApprovalRequestDetails

	err := tx.
		Select(
			"id",
			"approval_mapping_id",
			"approval_amount_id",
			"request_date",
			"source_doc_no",
		).
		Find(&approvalRequest, approvalRequestID).
		Error

	approvalResponse = approvalrequestpayloads.GetApprovalRequest{
		ApprovalMappingID:      approvalRequest.ApprovalMappingID,
		ApprovalAmountID:       approvalRequest.ApprovalAmountID,
		RequestDate:            approvalRequest.RequestDate,
		SourceDocNo:            approvalRequest.SourceDocNo,
		ApprovalRequestDetails: approvalRequestDetails,
	}

	if err != nil {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if approvalResponse.ApprovalMappingID == 0 {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(utils.GetDataNotFound),
		}
	}

	return approvalResponse, nil
}

// UpdateStatusByApprover implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) UpdateStatusByApprover(tx *gorm.DB, approvalRequest approvalrequestpayloads.UpdateStatusRequestDetails) (bool, *exceptions.BaseErrorResponse) {
	approvalrequestentities := approvalrequestentities.ApprovalRequestDetails{
		ID:       approvalRequest.ApprovalRequestDetailsID,
		StatusID: approvalRequest.StatusID,
	}

	err := tx.
		Updates(&approvalrequestentities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

// CreateApproval implements approvalrepo.ApprovalRepository.
func (*ApprovalRequestRepositoryImpl) Create(tx *gorm.DB, approvalRequestReq approvalrequestpayloads.CreateApprovalRequest) (bool, *exceptions.BaseErrorResponse) {
	approvalRequest := approvalrequestentities.ApprovalRequest{
		ApprovalMappingID: approvalRequestReq.ApprovalMappingID,
		ApprovalAmountID:  approvalRequestReq.ApprovalAmountID,
		StatusID:          utils.Draft,
		RequestDate:       time.Now(),
		SourceDocNo:       approvalRequestReq.SourceDocNo,
		RequestBy:         approvalRequestReq.RequestBy,
	}

	for details := range approvalRequest.ApprovalRequestDetails {
		approvalRequest.ApprovalRequestDetails = append(approvalRequest.ApprovalRequestDetails, approvalrequestentities.ApprovalRequestDetails{
			ApprovalApproverID: approvalRequest.ApprovalRequestDetails[details].ApprovalApproverID,
			ApproverID:         approvalRequest.ApprovalRequestDetails[details].ApproverID,
			StatusID:           approvalRequest.ApprovalRequestDetails[details].StatusID,
			Level:              approvalRequest.ApprovalRequestDetails[details].Level,
			UserID:             approvalRequest.ApprovalRequestDetails[details].UserID,
			Remark:             approvalRequest.ApprovalRequestDetails[details].Remark,
		})
	}

	err := tx.
		Create(&approvalRequest).
		Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return false, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusConflict,
				Err:        err,
			}
		} else {
			return false, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return true, nil
}

// Update implements approvalrequestrepo.ApprovalRequestRepository.
func (*ApprovalRequestRepositoryImpl) Update(tx *gorm.DB, approvalRequestReq approvalrequestpayloads.UpdateApprovalRequest) (bool, *exceptions.BaseErrorResponse) {
	approvalRequest := approvalrequestentities.ApprovalRequest{
		ID: approvalRequestReq.ApprovalRequestID,
	}

	for details := range approvalRequest.ApprovalRequestDetails {
		approvalRequest.ApprovalRequestDetails = append(approvalRequest.ApprovalRequestDetails, approvalrequestentities.ApprovalRequestDetails{
			ID:     approvalRequest.ApprovalRequestDetails[details].ID,
			Remark: approvalRequest.ApprovalRequestDetails[details].Remark,
		})
	}

	err := tx.
		Updates(&approvalRequest).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

func NewApprovalRequestRepository() approvalrequestrepo.ApprovalRequestRepository {
	return &ApprovalRequestRepositoryImpl{}
}
