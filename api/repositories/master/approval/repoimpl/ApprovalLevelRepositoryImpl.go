package approvalrepoimpl

import (
	"errors"
	"fmt"
	"net/http"
	approvalentities "user-services/api/entities/master/approval"
	approverentities "user-services/api/entities/master/approver"
	"user-services/api/exceptions"
	approvalpayloads "user-services/api/payloads/master/approval"
	approvalrepo "user-services/api/repositories/master/approval"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type ApprovalLevelRepositoryImpl struct {
}

// GetByUserID implements approvalrepo.ApprovalLevelRepository.
func (*ApprovalLevelRepositoryImpl) GetByUserID(
	tx *gorm.DB,
	approvalAmountID int,
	userID int,
	companyID int,
) (
	approvalpayloads.GetApprovalLevelResponse,
	*exceptions.BaseErrorResponse,
) {
	var approver approverentities.Approver
	approvalResponse := approvalpayloads.GetApprovalLevelResponse{}
	err := tx.
		Model(&approver).
		Select(
			"ApprovalApprover__ApprovalLevel.level",
			"ApprovalApprover__ApprovalLevel.count_required",
			"ApprovalApprover__ApprovalLevel.is_hierarchy",
		).
		Joins("ApproverDetails", tx.Select(
			"1",
		).Where("user_id = ?", userID)).
		Joins("ApprovalApprover", tx.Select(
			"1",
		)).
		Joins("ApprovalApprover.ApprovalLevel", tx.Select(
			"1",
		).Where("approval_amount_id = ?", approvalAmountID)).
		Joins("ApprovalApprover.ApprovalMapping", tx.Select(
			"1",
		).Where("ApproverDetails.company_id = ?", companyID)).
		Scan(&approvalResponse).
		Error

	if err != nil {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if approvalResponse.Level == nil {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusForbidden,
			Message:    fmt.Sprintf(utils.PermissionError, " for this approval or someone else already approve/revise this approval"),
			Err:        errors.New(utils.PermissionError),
		}
	}

	return approvalResponse, nil
}

// GetByApprovalApproverID implements approvalrepo.ApprovalLevelRepository.
func (*ApprovalLevelRepositoryImpl) GetByApprovalApproverID(
	tx *gorm.DB,
	approvalAmountID int,
	approvalApproverID int,
	companyID int,
) (
	approvalpayloads.GetApprovalLevelResponse,
	*exceptions.BaseErrorResponse,
) {
	approvalLevelEntity := approvalentities.ApprovalLevel{}
	approvalResponse := approvalpayloads.GetApprovalLevelResponse{}
	err := tx.
		Model(approvalLevelEntity).
		Select(
			"level",
			"count_required",
			"is_hierarchy",
		).
		InnerJoins("ApprovalApprover", tx.Select(
			"1",
		)).
		InnerJoins("ApprovalApprover.ApprovalMapping", tx.Select(
			"1",
		).Where("company_id", companyID)).
		Where(approvalentities.ApprovalLevel{
			ApprovalAmountID:   approvalAmountID,
			ApprovalApproverID: approvalApproverID,
		}).
		Scan(&approvalResponse).
		Error

	if err != nil {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if approvalResponse.Level == nil {
		return approvalResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return approvalResponse, nil
}

func NewApprovalLevelRepository() approvalrepo.ApprovalLevelRepository {
	return &ApprovalLevelRepositoryImpl{}
}
