package approvalservicesimpl

import (
	"user-services/api/exceptions"
	"user-services/api/helper"
	approvalpayloads "user-services/api/payloads/master/approval"
	approvalrepo "user-services/api/repositories/master/approval"
	approvalservices "user-services/api/services/master/approval"

	"gorm.io/gorm"
)

type ApprovalMappingServiceImpl struct {
	ApprovalMappingRepository approvalrepo.ApprovalMappingRepository
	DB                        *gorm.DB
}

// Create implements approvalservice.ApprovalMappingService.
func (service *ApprovalMappingServiceImpl) Create(createApprovalMappingRequest approvalpayloads.CreateApprovalMapping) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	create, err := service.ApprovalMappingRepository.Create(tx, createApprovalMappingRequest)

	if err != nil {
		return create, err
	}

	return create, nil
}

// Delete implements approvalservice.ApprovalMappingService.
func (service *ApprovalMappingServiceImpl) Delete(approvalID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApprovalMappingRepository.Delete(tx, approvalID)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

// Get implements approvalservice.ApprovalMappingService.
func (service *ApprovalMappingServiceImpl) Get(approvalMappingRequest approvalpayloads.GetApprovalMappingRequest) (approvalpayloads.GetApprovalMappingResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApprovalMappingRepository.Get(tx, approvalMappingRequest)

	if err != nil {
		return get, err
	}

	return get, nil
}

// Update implements approvalservice.ApprovalMappingService.
func (service *ApprovalMappingServiceImpl) Update(updateApprovalMappingRequest approvalpayloads.UpdateApprovalMapping) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApprovalMappingRepository.Update(tx, updateApprovalMappingRequest)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

func NewApprovalMappingService(
	ApprovalMappingRepository approvalrepo.ApprovalMappingRepository,
	db *gorm.DB,
) approvalservices.ApprovalMappingService {
	return &ApprovalMappingServiceImpl{
		ApprovalMappingRepository: ApprovalMappingRepository,
		DB:                        db,
	}
}
