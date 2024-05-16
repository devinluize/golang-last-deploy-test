package approvalrepoimpl

import (
	"net/http"
	"strings"
	approvalentities "user-services/api/entities/master/approval"
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
	approvalrepo "user-services/api/repositories/master/approval"

	"gorm.io/gorm"
)

type ApprovalMappingRepositoryImpl struct {
}

func NewApprovalMappingRepository() approvalrepo.ApprovalMappingRepository {
	return &ApprovalMappingRepositoryImpl{}
}

// Get implements approvalrepo.ApprovalMappingRepository.
func (*ApprovalMappingRepositoryImpl) Get(tx *gorm.DB, approvalMappingRequest approvalpayloads.GetApprovalMappingRequest) (approvalpayloads.GetApprovalMappingResponse, *exceptions.BaseErrorResponse) {
	var approvalMappingEntity approvalentities.ApprovalMapping
	var approvalResponse approvalpayloads.GetApprovalMappingResponse

	err := tx.
		Model(&approvalMappingEntity).
		Select(
			"approval_mapping.id approval_mapping_id",
			"ApprovalAmount.id approval_amount_id",
		).
		Joins("Approval", tx.Select(
			"1",
		).
			Where("is_void = ?", approvalMappingRequest.IsVoid),
		).
		Joins("ApprovalAmount", tx.Select(
			"1",
		).
			Where("max_amount >= ?", approvalMappingRequest.SourceAmount).
			Order("approval_amount.max_amount asc").
			Limit(1),
		).
		Where(approvalentities.ApprovalMapping{
			BrandID:           approvalMappingRequest.BrandID,
			CompanyID:         approvalMappingRequest.CompanyID,
			TransactionTypeID: approvalMappingRequest.TransactionTypeID,
			ProfitCenterID:    approvalMappingRequest.ProfitCenterID,
			CostCenterID:      approvalMappingRequest.CostCenterID,
			ModuleID:          approvalMappingRequest.ModuleID,
			DocumentTypeID:    approvalMappingRequest.DocumentTypeID,
		}).
		Scan(&approvalResponse).
		Error

	if err != nil {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if approvalResponse.ApprovalAmountID == 0 {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return approvalResponse, nil
}

// CreateApproval implements approvalrepo.ApprovalMappingRepository.
func (*ApprovalMappingRepositoryImpl) Create(tx *gorm.DB, approvalRequest approvalpayloads.CreateApprovalMapping) (bool, *exceptions.BaseErrorResponse) {
	var approvalMappingEntities approvalentities.ApprovalMapping

	err := tx.
		Create(&approvalMappingEntities).
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

// UpdateApproval implements approvalrepo.ApprovalMappingRepository.
func (*ApprovalMappingRepositoryImpl) Update(tx *gorm.DB, approvalRequest approvalpayloads.UpdateApprovalMapping) (bool, *exceptions.BaseErrorResponse) {
	var approvalMappingEntities approvalentities.ApprovalMapping

	err := tx.
		Updates(&approvalMappingEntities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

// Delete implements approvalrepo.ApprovalMappingRepository.
func (*ApprovalMappingRepositoryImpl) Delete(tx *gorm.DB, approvalMappingID int) (bool, *exceptions.BaseErrorResponse) {
	var approvalMappingEntities approvalentities.ApprovalMapping

	err := tx.
		Where(approvalMappingID).
		Delete(&approvalMappingEntities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return true, nil
}
