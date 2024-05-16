package main

import (
	"fmt"
	"net/http"
	"os"
	"user-services/api/config"
	migration "user-services/api/generate"
	"user-services/api/route"

	"github.com/go-playground/validator/v10"
)

// func init() {
// 	config.SetupConfiguration()
// }

//	@title			DMS User Service
//	@version		1.0
//	@description	DMS User Service Architecture
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Indomobil
//	@contact.url	https://github.com/IMSIDevOps
//	@contact.email	dev.ops@indomobil.com

//	@license.name	MIT
//	@license.url	https://github.com/IMSIDevOps/-/blob/main/LICENSE

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

//	@host		localhost:3000
//	@BasePath	/v1

func main() {
	args := os.Args
	env := ""
	if len(args) > 1 {
		env = args[1]
	}

	if env == "migrate" {
		err := migration.Migrate()
		if err != nil {
			handleServerError(err)
		}
	} else if env == "generate" {
		err := migration.Generate()
		if err != nil {
			handleServerError(err)
		}
	} else {
		config.InitEnvConfigs(false, env)
		db := config.InitDB()
		// dbRedis := config.InitRedisDB()
		validate := validator.New()
		// route.StartRouting(db, dbRedis, validate)
		route.StartRouting(db, nil, validate)

	}
}

func handleServerError(err error) {
	fmt.Printf("Error starting the server: %s\n", err)

	statusCode := http.StatusInternalServerError
	if isTemporaryError(err) {
		statusCode = http.StatusServiceUnavailable
	}

	os.Exit(statusCode)
}

func isTemporaryError(err error) bool {
	return false
}
