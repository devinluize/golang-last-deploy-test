package approvalservicesimpl

import (
	approvalentities "user-services/api/entities/master/approval"
	"user-services/api/exceptions"
	"user-services/api/helper"
	approvalpayloads "user-services/api/payloads/master/approval"
	approvalrepo "user-services/api/repositories/master/approval"
	approvalservices "user-services/api/services/master/approval"
	"user-services/api/utils"

	"gorm.io/gorm"
)

type ApprovalServiceImpl struct {
	ApprovalRepository approvalrepo.ApprovalRepository
	DB                 *gorm.DB
}

func (service *ApprovalServiceImpl) GetAll(page utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApprovalRepository.GetAll(tx, page)

	if err != nil {
		return get, err
	}

	return get, nil
}

// Create implements approvalservice.ApprovalService.
func (service *ApprovalServiceImpl) Create(createApprovalRequest approvalpayloads.CreateApproval) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	create, err := service.ApprovalRepository.Create(tx, createApprovalRequest)

	if err != nil {
		return create, err
	}

	return create, nil
}

// Delete implements approvalservice.ApprovalService.
func (service *ApprovalServiceImpl) Delete(approvalID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApprovalRepository.Delete(tx, approvalID)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

// Get implements approvalservice.ApprovalService.
func (service *ApprovalServiceImpl) Get(approvalID int) (approvalentities.Approval, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApprovalRepository.Get(tx, approvalID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// Update implements approvalservice.ApprovalService.
func (service *ApprovalServiceImpl) Update(approvalID int, updateApprovalRequest approvalpayloads.UpdateApproval) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApprovalRepository.Update(tx, approvalID, updateApprovalRequest)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

func NewApprovalService(
	ApprovalRepository approvalrepo.ApprovalRepository,
	db *gorm.DB,
) approvalservices.ApprovalService {
	return &ApprovalServiceImpl{
		ApprovalRepository: ApprovalRepository,
		DB:                 db,
	}
}
