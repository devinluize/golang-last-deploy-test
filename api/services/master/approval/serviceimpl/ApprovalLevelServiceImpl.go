package approvalservicesimpl

import (
	"user-services/api/exceptions"
	"user-services/api/helper"
	approvalpayloads "user-services/api/payloads/master/approval"
	approvalrepo "user-services/api/repositories/master/approval"
	approvalservices "user-services/api/services/master/approval"

	"gorm.io/gorm"
)

type ApprovalLevelServiceImpl struct {
	ApprovalLevelRepository approvalrepo.ApprovalLevelRepository
	DB                      *gorm.DB
}

// GetApprovalLevelByUserID implements approvalservices.ApprovalLevelService.
func (service *ApprovalLevelServiceImpl) GetByUserID(
	approvalAmountID int,
	userID int,
	companyID int,
) (
	approvalpayloads.GetApprovalLevelResponse,
	*exceptions.BaseErrorResponse,
) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApprovalLevelRepository.GetByUserID(
		tx,
		approvalAmountID,
		userID,
		companyID,
	)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetApprovalLevelByApprovalApproverID implements approvalservices.ApprovalService.
func (service *ApprovalLevelServiceImpl) GetByApprovalApproverID(
	approvalAmountID int,
	approvalApproverID int,
	companyID int,
) (
	approvalpayloads.GetApprovalLevelResponse, *exceptions.BaseErrorResponse,
) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApprovalLevelRepository.GetByApprovalApproverID(
		tx,
		approvalAmountID,
		approvalApproverID,
		companyID,
	)

	if err != nil {
		return get, err
	}

	return get, nil
}

func NewApprovalLevelService(
	ApprovalLevelRepository approvalrepo.ApprovalLevelRepository,
	db *gorm.DB,
) approvalservices.ApprovalLevelService {
	return &ApprovalLevelServiceImpl{
		ApprovalLevelRepository: ApprovalLevelRepository,
		DB:                      db,
	}
}
