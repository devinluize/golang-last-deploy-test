package approvalrequestservicesimpl

import (
	"fmt"
	"time"
	"user-services/api/exceptions"
	"user-services/api/helper"
	"user-services/api/payloads"
	approvalpayloads "user-services/api/payloads/master/approval"
	approverpayloads "user-services/api/payloads/master/approver"
	usergrouppayloads "user-services/api/payloads/master/user-group"
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
	approvalrepo "user-services/api/repositories/master/approval"
	approverrepo "user-services/api/repositories/master/approver"
	usergrouprepo "user-services/api/repositories/master/user-group"
	approvalrequestrepo "user-services/api/repositories/transaction"
	approvalrequestservices "user-services/api/services/transaction/approval-request"
	"user-services/api/utils"
	"user-services/api/utils/email"

	"gorm.io/gorm"
)

type ApprovalRequestServiceImpl struct {
	ApprovalRequestRepository approvalrequestrepo.ApprovalRequestRepository
	ApprovalLevelRepository   approvalrepo.ApprovalLevelRepository
	ApproverRepository        approverrepo.ApproverRepository
	ApprovalMappingRepository approvalrepo.ApprovalMappingRepository
	UserGroupRepository       usergrouprepo.UserGroupRepository
	DB                        *gorm.DB
}

// GetByUserID implements approvalrequestservices.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) GetByUserType(userType string, userID int, pagination utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	var approval utils.Pagination
	err := &exceptions.BaseErrorResponse{}
	if userType == "requester" {
		approval, err = service.ApprovalRequestRepository.GetByRequester(tx, userID, pagination)
	} else {
		approval, err = service.ApprovalRequestRepository.GetByApprover(tx, userID, pagination)
	}
	if err != nil {
		return approval, err
	}

	return approval, nil
}

// GetDetailsByRequester implements approvalrequestservices.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) GetDetailsByUserType(userType string, userID int, approvalID int, pagination utils.Pagination) (utils.Pagination, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	var approval utils.Pagination
	err := &exceptions.BaseErrorResponse{}
	if userType == "requester" {
		approval, err = service.ApprovalRequestRepository.GetDetailsByRequester(tx, userID, approvalID, pagination)
	} else {
		approval, err = service.ApprovalRequestRepository.GetDetailsByApprover(tx, userID, approvalID, pagination)
	}
	if err != nil {
		return approval, err
	}

	return approval, nil
}

// Create implements approvalrequestservice.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) Create(
	claims *payloads.UserDetail,
	approvalRequest approvalrequestpayloads.CreateApprovalRequest,
) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	approvalMapping, err := service.ApprovalMappingRepository.Get(
		tx,
		approvalpayloads.GetApprovalMappingRequest{
			CompanyID:         approvalRequest.CompanyID,
			TransactionTypeID: approvalRequest.TransactionTypeID,
			BrandID:           approvalRequest.BrandID,
			ProfitCenterID:    approvalRequest.ProfitCenterID,
			CostCenterID:      approvalRequest.CostCenterID,
			SourceAmount:      approvalRequest.SourceAmount,
			ModuleID:          approvalRequest.ModuleID,
			DocumentTypeID:    approvalRequest.DocumentTypeID,
			IsVoid:            approvalRequest.IsVoid,
		})

	if err != nil {
		return false, err
	}

	create, err := service.ApprovalRequestRepository.Create(tx, approvalrequestpayloads.CreateApprovalRequest{
		ApprovalMappingID: approvalMapping.ApprovalMappingID,
		ApprovalAmountID:  approvalMapping.ApprovalAmountID,
		CompanyID:         approvalRequest.CompanyID,
		ModuleID:          approvalRequest.ModuleID,
		DocumentTypeID:    approvalRequest.DocumentTypeID,
		RequestDate:       time.Now(),
		StatusID:          utils.Draft,
		SourceSysNo:       approvalRequest.SourceSysNo,
		SourceAmount:      approvalRequest.SourceAmount,
		SourceDate:        approvalRequest.SourceDate,
		RequestBy:         claims.UserID,
	})

	if err != nil {
		return create, err
	}

	level, err := service.ApproverRepository.CheckApprovalLevelByUserID(
		tx,
		approverpayloads.GetByUserIDRequest{
			ApprovalMappingID: approvalMapping.ApprovalMappingID,
			ApprovalAmountID:  approvalMapping.ApprovalAmountID,
			UserID:            claims.UserID,
		})

	if err != nil {
		return false, err
	}
	if *level == 0 {
		*level = 1
	} else {
		*level += 1
	}
	approvers, err := service.ApproverRepository.GetByLevel(
		tx,
		approverpayloads.GetByLevelRequest{
			ApprovalMappingID: approvalMapping.ApprovalMappingID,
			ApprovalAmountID:  approvalMapping.ApprovalAmountID,
			Level:             *level,
		})

	if err != nil {
		return false, err
	}

	emails := []string{}
	for _, email := range *approvers {
		emails = append(emails, *email.Email)
	}

	_, err = email.SendEmails(
		emails,
		&email.EmailData{
			URL:     "TEST URL",
			Subject: "Approval Request",
			Remark:  "TEST REMARK",
			Date:    time.Now(),
		},
		"Approval.html")

	if err != nil {
		return false, err
	}

	return create, nil
}

// CreateDetails implements approvalrequestservices.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) CreateDetails(
	claims *payloads.UserDetail,
	approvalRequest approvalrequestpayloads.UpdateApprovalRequestDetails,
) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	approverRequestDetails := []approvalrequestpayloads.CreateApprovalRequestDetails{}
	approvalRequestDetails := []approvalrequestpayloads.UpdateApprovalRequestDetails{}
	approvalRequestDetails = append(approvalRequestDetails, approvalrequestpayloads.UpdateApprovalRequestDetails{
		StatusID: approvalRequest.StatusID, //Waiting for Status API
		Remark:   approvalRequest.Remark,
	})
	approvalRequestMaster, err := service.ApprovalRequestRepository.Get(
		tx,
		approvalRequest.ApprovalRequestID,
	)
	if err != nil {
		return false, err
	}
	approvalLevel, err := service.ApprovalLevelRepository.GetByUserID(
		tx,
		approvalRequestMaster.ApprovalAmountID,
		claims.UserID,
		approvalRequest.CompanyID,
	)
	if err != nil {
		return false, err
	}

	// If Approver Action is Revise then update the Header to Revise/Reject
	if approvalRequest.StatusID == utils.Revise {
		_, err := service.ApprovalRequestRepository.Update(
			tx,
			approvalrequestpayloads.UpdateApprovalRequest{
				ApprovalRequestID:      approvalRequest.ApprovalRequestID,
				StatusID:               approvalRequest.StatusID,
				ApprovalRequestDetails: approvalRequestDetails,
			})
		if err != nil {
			return false, err
		}
	} else {
		//Check if count required with approved status on the same level is fullfiled
		if approvalLevel.CountRequired != utils.CountPayloadsWithStatusID(approvalRequestMaster, *approvalLevel.Level) {
			approver, err := service.ApproverRepository.GetByUserID(
				tx,
				approverpayloads.GetByUserIDRequest{
					ApprovalMappingID: approvalRequestMaster.ApprovalMappingID,
					ApprovalAmountID:  approvalRequestMaster.ApprovalAmountID,
					UserID:            claims.UserID,
				})

			if err != nil {
				return false, err
			}
			approverRequestDetails = append(approverRequestDetails, approvalrequestpayloads.CreateApprovalRequestDetails{
				ApprovalRequestID:  approvalRequest.ApprovalRequestID,
				ApprovalApproverID: *approver.ApprovalApproverID,
				ApproverID:         *approver.ApproverID,
				Level:              *approvalLevel.Level,
				UserId:             claims.UserID,
				StatusID:           approvalRequest.StatusID,
				// AllowApprove:      true,
				Remark: approvalRequest.Remark,
			})
			create, err := service.ApprovalRequestRepository.CreateDetails(tx, approverRequestDetails)
			if err != nil {
				return create, err
			}
		} else if approvalLevel.CountRequired == utils.CountPayloadsWithStatusID(approvalRequestMaster, *approvalLevel.Level) {
			approver, err := service.ApproverRepository.GetByUserID(
				tx,
				approverpayloads.GetByUserIDRequest{
					ApprovalMappingID: approvalRequestMaster.ApprovalMappingID,
					ApprovalAmountID:  approvalRequestMaster.ApprovalAmountID,
					UserID:            claims.UserID,
				})

			if err != nil {
				return false, err
			}

			approverRequestDetails = append(approverRequestDetails, approvalrequestpayloads.CreateApprovalRequestDetails{
				ApprovalRequestID: approvalRequest.ApprovalRequestID,
				ApproverID:        *approver.ApproverID,
				Level:             *approvalLevel.Level,
				UserId:            claims.UserID,
				StatusID:          approvalRequest.StatusID,
				// AllowApprove:      true,
				Remark: approvalRequest.Remark,
			})
			create, err := service.ApprovalRequestRepository.CreateDetails(tx, approverRequestDetails)
			if err != nil {
				return create, err
			}
			var approvers *[]approverpayloads.GetApproverForApprovalResponse
			if *approvalLevel.IsHierarchy {
				approvers, err = service.UserGroupRepository.GetLeaderForApproval(
					tx,
					usergrouppayloads.GetLeaderForApprovalRequest{
						CompanyID: 1,
						UserID:    claims.UserID,
					})

				if err != nil {
					return false, err
				}
			} else {
				approvers, err = service.ApproverRepository.GetByLevel(
					tx,
					approverpayloads.GetByLevelRequest{
						ApprovalMappingID: approvalRequestMaster.ApprovalMappingID,
						ApprovalAmountID:  approvalRequestMaster.ApprovalAmountID,
						Level:             *approvalLevel.Level + 1,
					})

				if err != nil {
					return false, err
				}
			}

			emails := []string{}
			for _, email := range *approvers {
				emails = append(emails, *email.Email)
			}
			fmt.Println(emails)
			_, err = email.SendEmails(
				emails,
				&email.EmailData{
					URL:     "",
					Subject: "Approval Request",
					Remark:  "",
					Date:    time.Now(),
				},
				"Approval.html")
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

// UpdateDetails implements approvalrequestservices.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) UpdateDetails(createApprovalRequestDetails approvalrequestpayloads.UpdateApprovalRequestDetails) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	create, err := service.ApprovalRequestRepository.UpdateDetails(tx, createApprovalRequestDetails)

	if err != nil {
		return create, err
	}

	return create, nil
}

// Get implements approvalrequestservice.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) Get(approvalrequestID int) (approvalrequestpayloads.GetApprovalRequest, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.ApprovalRequestRepository.Get(tx, approvalrequestID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// Update implements approvalrequestservice.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) Update(updateApprovalRequestRequest approvalrequestpayloads.UpdateApprovalRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApprovalRequestRepository.Update(tx, updateApprovalRequestRequest)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

// Update implements approvalrequestservice.ApprovalRequestService.
func (service *ApprovalRequestServiceImpl) UpdateStatusByApprover(updateApprovalRequestRequest approvalrequestpayloads.UpdateStatusRequestDetails) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	delete, err := service.ApprovalRequestRepository.UpdateStatusByApprover(tx, updateApprovalRequestRequest)

	if err != nil {
		return delete, err
	}

	return delete, nil
}

func NewApprovalRequestService(
	ApprovalRequestRepository approvalrequestrepo.ApprovalRequestRepository,
	ApproverRepository approverrepo.ApproverRepository,
	ApprovalLevelRepository approvalrepo.ApprovalLevelRepository,
	ApprovalMappingRepository approvalrepo.ApprovalMappingRepository,
	UserGroupRepository usergrouprepo.UserGroupRepository,
	db *gorm.DB,
) approvalrequestservices.ApprovalRequestService {
	return &ApprovalRequestServiceImpl{
		ApprovalRequestRepository: ApprovalRequestRepository,
		ApproverRepository:        ApproverRepository,
		ApprovalLevelRepository:   ApprovalLevelRepository,
		ApprovalMappingRepository: ApprovalMappingRepository,
		UserGroupRepository:       UserGroupRepository,
		DB:                        db,
	}
}
