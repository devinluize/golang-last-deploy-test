package approvaltest

import (
	"errors"
	"fmt"
	"testing"
	"user-services/api/config"
	masterentities "user-services/api/entities/master"
	approverpayloads "user-services/api/payloads/master/approver"
	"user-services/api/utils"
)

func TestGetApprovalLevelByUserID(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()

	var users masterentities.User
	var approverResponse approverpayloads.GetApproverForApprovalByUserIDResponse

	err := db.
		Model(&users).
		Select(
			"ApproverDetails__Approver.id approver_id",
			"ApproverDetails.id approver_detail_id",
			"ApproverDetails__Approver__ApprovalApprover.id approval_approver_id",
			"level",
			"email",
		).
		Joins("ApproverDetails", db.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver", db.Select(
			"1",
		)).
		Joins("ApproverDetails.Approver.ApprovalApprover",
			db.Select(
				"1",
			)).
		Joins("ApproverDetails.Approver.ApprovalApprover.ApprovalLevel",
			db.Select(
				"1",
			)).
		Where("users.id = ? and approval_amount_id = ? and approval_mapping_id = ?",
			1, 1, 1).
		Scan(&approverResponse).
		Error

	if err != nil {
		fmt.Println(approverResponse, errors.New(utils.GetDataFailed))
	}

	fmt.Println(approverResponse)

}
