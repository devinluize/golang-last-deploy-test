package approverservicesimpl

import (
	approverentities "user-services/api/entities/master/approver"
	"user-services/api/exceptions"
	"user-services/api/helper"
	approverpayloads "user-services/api/payloads/master/approver"
	approverrepo "user-services/api/repositories/master/approver"
	approverservice "user-services/api/services/master/approver"

	"gorm.io/gorm"
)

type ApproverServiceImpl struct {
	ApproverRepository approverrepo.ApproverRepository
	DB                 *gorm.DB
}

// GetAll implements approverservices.ApproverService.
func (service *ApproverServiceImpl) GetAll() ([]approverentities.Approver, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApproverRepository.GetAll(tx)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetByUserID implements approverservices.ApproverService.
func (service *ApproverServiceImpl) GetByUserID(approverRequest approverpayloads.GetByUserIDRequest) (approverpayloads.GetApproverForApprovalByUserIDResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApproverRepository.GetByUserID(tx, approverRequest)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetByLevel implements approverservices.ApproverService.
func (service *ApproverServiceImpl) GetByLevel(approverRequest approverpayloads.GetByLevelRequest) (*[]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApproverRepository.GetByLevel(tx, approverRequest)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetApprovers implements approvalservices.ApprovalService.
func (service *ApproverServiceImpl) GetApprovers(approvalID int) ([]approverpayloads.GetApproverForApprovalResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApproverRepository.GetApprovers(tx, approvalID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// Create implements approverservice.ApproverService.
func (service *ApproverServiceImpl) Create(createApproverRequest approverpayloads.CreateApprover) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	create, err := service.ApproverRepository.Create(tx, createApproverRequest)

	if err != nil {
		return create, err
	}

	return create, nil
}

// Delete implements approverservice.ApproverService.
func (service *ApproverServiceImpl) Delete(approverID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApproverRepository.Delete(tx, approverID)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

// Get implements approverservice.ApproverService.
func (service *ApproverServiceImpl) Get(approverID int) (approverpayloads.CreateApprover, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApproverRepository.Get(tx, approverID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// Update implements approverservice.ApproverService.
func (service *ApproverServiceImpl) Update(approverID int, updateApproverRequest approverpayloads.UpdateApprover) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApproverRepository.Update(tx, approverID, updateApproverRequest)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

func NewApproverService(
	ApproverRepository approverrepo.ApproverRepository,
	db *gorm.DB,
) approverservice.ApproverService {
	return &ApproverServiceImpl{
		ApproverRepository: ApproverRepository,
		DB:                 db,
	}
}
