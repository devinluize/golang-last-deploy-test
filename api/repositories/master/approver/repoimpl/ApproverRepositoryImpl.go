package approverrepoimpl

import (
	"fmt"
	"net/http"
	"strings"
	masterentities "user-services/api/entities/master"
	approverentities "user-services/api/entities/master/approver"
	"user-services/api/exceptions"
	approverpayloads "user-services/api/payloads/master/approver"
	approverrepo "user-services/api/repositories/master/approver"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type ApproverRepositoryImpl struct {
}

func (*ApproverRepositoryImpl) GetAll(tx *gorm.DB) ([]approverentities.Approver, *exceptions.BaseErrorResponse) {
	var approver []approverentities.Approver

	rows, err := tx.
		Model(&approver).
		Scan(&approver).
		Rows()

	if err != nil {

		return approver, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer rows.Close()
	return approver, nil
}

// CheckApprovalLevelByUserID implements approverrepo.ApproverRepository.
func (*ApproverRepositoryImpl) CheckApprovalLevelByUserID(tx *gorm.DB, approverRequest approverpayloads.GetByUserIDRequest) (*int, *exceptions.BaseErrorResponse) {
	var user masterentities.User
	var level *int

	err := tx.
		Select(
			"level",
		).
		Joins("ApproverDetails", tx.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver", tx.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver.ApprovalApprover",
			tx.Select(
				"1",
			)).
		Joins("ApproverDetails.Approver.ApprovalApprover.ApprovalLevel",
			tx.Select(
				"1",
			)).
		Where("users.id = ? and approval_amount_id = ? and approval_mapping_id = ?", approverRequest.UserID, approverRequest.ApprovalAmountID, approverRequest.ApprovalMappingID).
		Model(&user).
		Scan(&level).
		Error

	if err != nil {
		return level, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if level == nil {
		return nil, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("%s %s", "Approval Level ", utils.GetDataNotFound),
			Err:        err,
		}
	}

	return level, nil
}

// GetByUserID implements approverrepo.ApproverRepository.
func (*ApproverRepositoryImpl) GetByUserID(tx *gorm.DB, approverRequest approverpayloads.GetByUserIDRequest) (approverpayloads.GetApproverForApprovalByUserIDResponse, *exceptions.BaseErrorResponse) {
	var users masterentities.User
	var approverResponse approverpayloads.GetApproverForApprovalByUserIDResponse

	err := tx.
		Model(&users).
		Select(
			"ApproverDetails__Approver.id approver_id",
			"ApproverDetails.id approver_detail_id",
			"ApproverDetails__Approver__ApprovalApprover.id approval_approver_id",
			"level",
			"email",
		).
		Joins("ApproverDetails", tx.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver", tx.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver.ApprovalApprover",
			tx.Select(
				"1",
			)).
		Joins("ApproverDetails.Approver.ApprovalApprover.ApprovalLevel",
			tx.Select(
				"1",
			)).
		Where("users.id = ? and approval_amount_id = ? and approval_mapping_id = ?", approverRequest.UserID, approverRequest.ApprovalAmountID, approverRequest.ApprovalMappingID).
		Scan(&approverResponse).
		Error

	if err != nil {
		return approverResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if approverResponse.ApprovalApproverID == nil {
		return approverResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("%s %s", "Approver ", utils.GetDataNotFound),
			Err:        err,
		}
	}

	return approverResponse, nil
}

// GetByLevel implements approverrepo.ApproverRepository.
func (*ApproverRepositoryImpl) GetByLevel(tx *gorm.DB, approverRequest approverpayloads.GetByLevelRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse) {
	var users []masterentities.User
	var approverResponse []approverpayloads.GetApproverForApprovalResponse
	err := tx.
		Model(&users).
		Select(
			"ApproverDetails__Approver__ApprovalApprover.approver_id",
			"ApproverDetails.id approver_detail_id",
			"ApproverDetails__Approver__ApprovalApprover.id approval_approver_id",
			"level",
			"users.id user_id",
			"email",
		).
		Joins("ApproverDetails", tx.Select(
			"id",
		)).
		Joins("ApproverDetails.Approver", tx.Select(
			"id",
		)).
		Joins("ApproverDetails.Approver.ApprovalApprover",
			tx.Select(
				"approval_approver_id",
				"level",
			)).
		Joins("ApproverDetails.Approver.ApprovalApprover.ApprovalLevel",
			tx.Select(
				"approval_approver_id",
				"level",
			)).
		Where("level = ? and approval_amount_id = ? and approval_mapping_id = ?", approverRequest.Level, approverRequest.ApprovalAmountID, approverRequest.ApprovalMappingID).
		Scan(&approverResponse).
		Error

	if err != nil {
		return &approverResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if approverResponse == nil {
		return &approverResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        fmt.Errorf("%s %s", "Approver ", utils.GetDataNotFound),
		}
	}

	return &approverResponse, nil
}

// Get Approvers implements approverrepo.ApproverRepository.
func (*ApproverRepositoryImpl) GetApprovers(
	tx *gorm.DB,
	approvalID int,
) (
	[]approverpayloads.GetApproverForApprovalResponse,
	*exceptions.BaseErrorResponse,
) {
	var approver []approverentities.Approver
	var approverResponse []approverpayloads.GetApproverForApprovalResponse

	err := tx.
		Model(&approver).
		Select(
			"Approver.ID approver_id",
		).
		InnerJoins("ApproverDetails", tx.Select(
			"1",
		)).
		InnerJoins("ApprovalApprover", tx.Select(
			"1",
		).Where("approval_mapping_id = ?", approvalID)).
		InnerJoins("ApprovalApprover.ApprovalLevel", tx.Select(
			"1",
		)).
		Scan(&approverResponse).
		Error

	if err != nil {
		return approverResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return approverResponse, nil
}

// Get implements approverrepo.ApproverRepository.
func (*ApproverRepositoryImpl) Get(tx *gorm.DB, approverID int) (approverpayloads.CreateApprover, *exceptions.BaseErrorResponse) {
	var approver approverentities.Approver
	var approverResponse approverpayloads.CreateApprover

	err := tx.
		Model(approver).
		Where(approverID).
		Scan(&approverResponse).
		Error

	if err != nil {
		return approverResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return approverResponse, nil
}

// Create implements approverrepo.ApproverRepository.
func (*ApproverRepositoryImpl) Create(tx *gorm.DB, approverRequest approverpayloads.CreateApprover) (bool, *exceptions.BaseErrorResponse) {
	approver := approverentities.Approver{
		IsActive: true,
		Code:     approverRequest.ApproverCode,
		Name:     approverRequest.ApproverName,
	}

	for i := range approverRequest.ApproverDetails {
		approver.ApproverDetails = append(approver.ApproverDetails, approverentities.ApproverDetails{
			UserID: approverRequest.ApproverDetails[i].UserID,
		})
	}

	err := tx.
		Create(&approver).
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

// UpdateApprover implements approverrepo.ApproverRepository.
func (*ApproverRepositoryImpl) Update(
	tx *gorm.DB,
	approverID int,
	approverRequest approverpayloads.UpdateApprover,
) (bool, *exceptions.BaseErrorResponse) {
	approver := approverentities.Approver{
		ID:       approverID,
		IsActive: approverRequest.IsActive,
		Code:     approverRequest.ApproverCode,
		Name:     approverRequest.ApproverName,
	}

	for i := range approverRequest.ApproverDetails {
		approver.ApproverDetails = append(approver.ApproverDetails, approverentities.ApproverDetails{
			ID:     approverRequest.ApproverDetails[i].ApproverDetailID,
			UserID: approverRequest.ApproverDetails[i].UserID,
		})
	}

	err := tx.
		Updates(&approver).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil

}

// Delete implements approvalrepo.ApprovalRepository.
func (*ApproverRepositoryImpl) Delete(tx *gorm.DB, approverID int) (bool, *exceptions.BaseErrorResponse) {
	var approverEntities approverentities.Approver
	err := tx.
		Where(approverID).
		Delete(&approverEntities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return true, nil
}

func NewApproverRepository() approverrepo.ApproverRepository {
	return &ApproverRepositoryImpl{}
}
