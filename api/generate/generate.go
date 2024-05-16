package migration

import (
	"fmt"
	"user-services/api/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Generate() error {
	config.InitEnvConfigs(false, "")
	g := gen.NewGenerator(gen.Config{
		OutPath: "../api/entities",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := fmt.Sprintf(`%s://%s:%s@%s:%v?database=%s`, config.EnvConfigs.DBDriver, config.EnvConfigs.DBUser, config.EnvConfigs.DBPass, config.EnvConfigs.DBHost, config.EnvConfigs.DBPort, config.EnvConfigs.DBName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "dbo.", // schema name
			SingularTable: false,
		}})

	if err != nil {
		logrus.Info("Cannot connected database ", err)
		return err
	}
	g.UseDB(db) // reuse your gorm db

	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.UserDetail{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func() {}, model.UserDetail{})

	// // Generate the code
	g.GenerateModel("users")
	g.Execute()
	return nil
}

// gentool -dsn "sqlserver://sa:P@ssw0rd@10.1.32.62:1433?database=dms_microservices_dev" -tables "auth_user"
