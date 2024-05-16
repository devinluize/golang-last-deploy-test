package utils

import (
	approvalrequestpayloads "user-services/api/payloads/transaction/approval-request"
)

func FindMaxLevel(count approvalrequestpayloads.GetApprovalRequest) int {
	if len(count.ApprovalRequestDetails) == 0 {
		// Handle the case where the slice is empty
		return 1
	}

	maxValue := count.ApprovalRequestDetails[0].Level
	for _, details := range count.ApprovalRequestDetails {
		if details.Level > maxValue {
			maxValue = details.Level + 1
		}
	}

	if maxValue == 0 {
		return 1
	}

	return maxValue
}

func CountPayloadsWithStatusID(payloads approvalrequestpayloads.GetApprovalRequest, level int) *int {
	count := 0
	for _, payload := range payloads.ApprovalRequestDetails {
		if payload.StatusID != 1 && payload.Level == level {
			count++
		}
	}
	count += 1
	return &count
}
