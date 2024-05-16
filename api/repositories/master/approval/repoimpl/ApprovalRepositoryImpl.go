package approvalrepoimpl

import (
	"errors"
	"net/http"
	"strings"
	approvalentities "user-services/api/entities/master/approval"
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
	approvalrepo "user-services/api/repositories/master/approval"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type ApprovalRepositoryImpl struct {
}

func NewApprovalRepository() approvalrepo.ApprovalRepository {
	return &ApprovalRepositoryImpl{}
}

// CheckDataExists implements approvalrepo.ApprovalRepository.
func (*ApprovalRepositoryImpl) CheckDataExists(tx *gorm.DB, approvalCode string) (bool, *exceptions.BaseErrorResponse) {
	var exists bool
	err := tx.Model(approvalentities.Approval{}).
		Select("count(code)").
		Where(
			"code = ?",
			approvalCode,
		).
		Find(&exists).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if !exists {
		return exists, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("not exists"),
		}
	}
	return exists, nil
}

// Get implements approvalrepo.ApprovalRepository.
func (*ApprovalRepositoryImpl) GetAll(tx *gorm.DB, pagination utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	var approval []approvalentities.Approval
	err := tx.
		Model(&approval).
		Scopes(utils.Paginate(approval, &pagination, tx.Model(&approval))).
		Scan(&approval).
		Error

	if err != nil {
		return pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if len(approval) == 0 {
		return pagination, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(utils.GetDataNotFound),
		}
	}
	pagination.Rows = approval

	return pagination, nil
}

// Get implements approvalrepo.ApprovalRepository.
func (*ApprovalRepositoryImpl) Get(tx *gorm.DB, approvalID int) (approvalentities.Approval, *exceptions.BaseErrorResponse) {
	var approval approvalentities.Approval

	err := tx.
		Model(approval).
		Where(approvalID).
		Scan(&approval).
		Error

	if err != nil {
		return approval, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if approval.ID == 0 {
		return approval, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return approval, nil
}

// CreateApproval implements approvalrepo.ApprovalRepository.
func (*ApprovalRepositoryImpl) Create(tx *gorm.DB, approvalRequest approvalpayloads.CreateApproval) (bool, *exceptions.BaseErrorResponse) {
	approvalEntities := approvalentities.Approval{
		Code: approvalRequest.ApprovalCode,
		Name: approvalRequest.ApprovalName,
	}

	for maps, mapping := range approvalRequest.ApprovalMapping {
		approvalEntities.ApprovalMapping = append(approvalEntities.ApprovalMapping, approvalentities.ApprovalMapping{
			CompanyID:         mapping.CompanyID,
			ModuleID:          mapping.ModuleID,
			DocumentTypeID:    mapping.DocumentTypeID,
			TransactionTypeID: mapping.TransactionTypeID,
			BrandID:           mapping.BrandID,
			ProfitCenterID:    mapping.ProfitCenterID,
			CostCenterID:      mapping.CostCenterID,
		})
		for amounts, amount := range mapping.ApprovalAmount {
			approvalEntities.ApprovalMapping[maps].ApprovalAmount = append(approvalEntities.ApprovalMapping[maps].ApprovalAmount, approvalentities.ApprovalAmount{
				MaxAmount: amount.MaxAmount,
			})
			for _, level := range amount.ApprovalLevel {
				approvalEntities.ApprovalMapping[maps].ApprovalAmount[amounts].ApprovalLevel =
					append(approvalEntities.ApprovalMapping[maps].ApprovalAmount[amounts].ApprovalLevel,
						approvalentities.ApprovalLevel{
							Level:              level.Level,
							ApprovalApproverID: level.ApprovalApproverID,
						},
					)
			}
		}
	}

	err := tx.
		Create(&approvalEntities).
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

// UpdateApproval implements approvalrepo.ApprovalRepository.
func (*ApprovalRepositoryImpl) Update(tx *gorm.DB, approvalID int, approvalRequest approvalpayloads.UpdateApproval) (bool, *exceptions.BaseErrorResponse) {
	approvalEntities := approvalentities.Approval{
		ID:   approvalID,
		Code: approvalRequest.ApprovalCode,
		Name: approvalRequest.ApprovalName,
	}

	for maps, mapping := range approvalRequest.ApprovalMapping {
		approvalEntities.ApprovalMapping[maps] = approvalentities.ApprovalMapping{
			ApprovalID:        approvalRequest.ApprovalID,
			CompanyID:         mapping.CompanyID,
			ModuleID:          mapping.ModuleID,
			DocumentTypeID:    mapping.DocumentTypeID,
			TransactionTypeID: mapping.TransactionTypeID,
			BrandID:           mapping.BrandID,
			ProfitCenterID:    mapping.ProfitCenterID,
			CostCenterID:      mapping.CostCenterID,
		}
		for amounts, amount := range approvalRequest.ApprovalMapping[maps].ApprovalAmount {
			approvalEntities.ApprovalMapping[maps].ApprovalAmount = append(approvalEntities.ApprovalMapping[maps].ApprovalAmount, approvalentities.ApprovalAmount{
				ID:        amount.ApprovalAmountID,
				MaxAmount: amount.MaxAmount,
			})
			for _, level := range amount.ApprovalLevel {
				approvalEntities.ApprovalMapping[maps].ApprovalAmount[amounts].ApprovalLevel =
					append(approvalEntities.ApprovalMapping[maps].ApprovalAmount[amounts].ApprovalLevel,
						approvalentities.ApprovalLevel{
							Level:              level.Level,
							ApprovalApproverID: level.ApprovalApproverID,
						},
					)
			}
		}
	}

	err := tx.
		Updates(&approvalEntities).
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
func (*ApprovalRepositoryImpl) Delete(tx *gorm.DB, approvalID int) (bool, *exceptions.BaseErrorResponse) {
	var approvalEntities approvalentities.Approval
	err := tx.
		Where(approvalID).
		Delete(&approvalEntities).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return true, nil
}
