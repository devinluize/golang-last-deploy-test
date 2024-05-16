package migration

import (
	"fmt"
	"user-services/api/config"
	masterentities "user-services/api/entities/master"
	approvalentities "user-services/api/entities/master/approval"
	approverentities "user-services/api/entities/master/approver"
	menuentities "user-services/api/entities/master/menu"
	usergroupentities "user-services/api/entities/master/user-group"
	approvalrequestentities "user-services/api/entities/transaction/approval-request"

	// approvalentities "user-service/api/entities/master/approval"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Migrate() error {
	config.InitEnvConfigs(false, "")
	logEntry := "Auto Migrating..."
	// dsn := fmt.Sprintf(`%s://%s:%s@%s:%v?database=%s&connection+timeout=30`, config.config.config.EnvConfigs.DBDriver, config.config.config.EnvConfigs.DBUser, config.config.config.EnvConfigs.DBPass, config.config.config.EnvConfigs.DBHost, config.config.config.EnvConfigs.DBPort, config.config.config.EnvConfigs.DBName)
	dsn := fmt.Sprintf(`%s://%s:%s@%s:%v?database=%s`, config.EnvConfigs.DBDriver, config.EnvConfigs.DBUser, config.EnvConfigs.DBPass, config.EnvConfigs.DBHost, config.EnvConfigs.DBPort, config.EnvConfigs.DBName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "dbo.", // schema name
			SingularTable: false,
		}})
	db.AutoMigrate(
		&masterentities.Role{},
		&masterentities.User{},
		&approverentities.Approver{},
		&approverentities.ApproverDetails{},
		&approvalentities.Approval{},
		&approvalentities.ApprovalAmount{},
		&approvalentities.ApprovalMapping{},
		&approvalentities.ApprovalLevel{},
		&approvalrequestentities.ApprovalRequest{},
		&approvalrequestentities.ApprovalRequestDetails{},
		&menuentities.MenuUrl{},
		&menuentities.MenuList{},
		&menuentities.MenuAccess{},
		&menuentities.MenuUserAccess{},
		&usergroupentities.UserGroup{},
		&usergroupentities.UserGroupCompany{},
		&usergroupentities.UserGroupLeader{},
		&usergroupentities.UserGroupMember{},
	) //<--- Line 84
	if db != nil && db.Error != nil {
		//We have an error
		fmt.Sprintf("%s %s with error %s", logEntry, "Failed", db.Error)
		logrus.Info(err)
		return err
	}
	fmt.Sprintf("%s %s", logEntry, "Success")
	return nil
}
