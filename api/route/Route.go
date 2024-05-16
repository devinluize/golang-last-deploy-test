package route

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/gorm"
	"net/http"
	"user-services/api/config"
	authcontrollers "user-services/api/controllers/auth"
	"user-services/api/controllers/binning/controllerImpl"
	approvalcontrollers "user-services/api/controllers/master/approval"
	approvercontrollers "user-services/api/controllers/master/approver"
	menucontrollers "user-services/api/controllers/master/menu"
	approvalrequestcontrollers "user-services/api/controllers/transaction"
	usercontrollers "user-services/api/controllers/user"
	"user-services/api/repositories/binning/repoImpl"
	approvalrepoimpl "user-services/api/repositories/master/approval/repoimpl"
	approverrepoimpl "user-services/api/repositories/master/approver/repoimpl"
	menurepoimpl "user-services/api/repositories/master/menu/repoimpl"
	usergrouprepoimpl "user-services/api/repositories/master/user-group/repoimpl"
	redisrepoimpl "user-services/api/repositories/redis/repoimpl"
	approvalrequestrepoimpl "user-services/api/repositories/transaction/repoimpl"
	userrepoimpl "user-services/api/repositories/user/repoimpl"
	"user-services/api/services/binning/serviceImpl"
	approvalservicesimpl "user-services/api/services/master/approval/serviceimpl"
	approverservicesimpl "user-services/api/services/master/approver/serviceimpl"
	menuservicesimpl "user-services/api/services/master/menu/serviceimpl"
	usergroupservicesimpl "user-services/api/services/master/user-group/serviceimpl"
	redisserviceimpl "user-services/api/services/redis/serviceimpl"
	approvalrequestservicesimpl "user-services/api/services/transaction/approval-request/serviceimpl"
	userserviceimpl "user-services/api/services/user/serviceimpl"
	_ "user-services/docs"
)

func StartRouting(db *gorm.DB, dbRedis *config.Database, validate *validator.Validate) {
	r := chi.NewRouter()

	r.Mount("/v1", versionedRouterV1(db, dbRedis, validate))
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	http.ListenAndServe(fmt.Sprintf(":%v", config.EnvConfigs.Port), r)
}

func versionedRouterV1(db *gorm.DB, dbRedis *config.Database, validate *validator.Validate) chi.Router {
	r := chi.NewRouter()
	//Redis
	authRepository := userrepoimpl.NewAuthRepository()
	userRepository := userrepoimpl.NewUserRepository()
	redisRepository := redisrepoimpl.NewRedisRepository()
	redisService := redisserviceimpl.NewRedisService(db, dbRedis, authRepository, userRepository, redisRepository)
	//binning
	binningRepository := repoImpl.NewBinningRepositoryImpl()
	binningService := serviceImpl.NewBinningServiceImpl(db, binningRepository)
	binningController := controllerImpl.NewBinningControllerImpl(binningService)
	//User
	userService := userserviceimpl.NewUserService(userRepository, db, validate)
	userController := usercontrollers.NewUserController(userService)

	userGroupRepository := usergrouprepoimpl.NewUserGroupRepository()
	userGroupService := usergroupservicesimpl.NewUserGroupService(userGroupRepository, db)

	//Approval
	approvalRepository := approvalrepoimpl.NewApprovalRepository()
	approvalRequestRepository := approvalrequestrepoimpl.NewApprovalRequestRepository()
	approverRepository := approverrepoimpl.NewApproverRepository()
	approvalMappingRepository := approvalrepoimpl.NewApprovalMappingRepository()
	approvalLevelRepository := approvalrepoimpl.NewApprovalLevelRepository()

	approvalService := approvalservicesimpl.NewApprovalService(approvalRepository, db)
	approvalRequestService := approvalrequestservicesimpl.NewApprovalRequestService(approvalRequestRepository, approverRepository, approvalLevelRepository, approvalMappingRepository, userGroupRepository, db)
	approverService := approverservicesimpl.NewApproverService(approverRepository, db)
	approvalMappingService := approvalservicesimpl.NewApprovalMappingService(approvalMappingRepository, db)
	approvalLevelService := approvalservicesimpl.NewApprovalLevelService(approvalLevelRepository, db)

	approvalController := approvalcontrollers.NewApprovalController(approvalService, approvalMappingService, approvalLevelService)
	approvalRequestController := approvalrequestcontrollers.NewApprovalRequestController(approvalService, approvalRequestService, approverService, approvalMappingService, approvalLevelService, userGroupService)
	approverController := approvercontrollers.NewApproverController(approverService)

	//Auth
	authService := userserviceimpl.NewAuthService(db, dbRedis, authRepository, userRepository, redisRepository, validate)
	authController := authcontrollers.NewAuthController(authService, userService, redisService)
	//Menu
	menuAccessRepository := menurepoimpl.NewMenuAccessRepository()
	menuListRepository := menurepoimpl.NewMenuListRepository()
	menuUrlRepository := menurepoimpl.NewMenuURLRepository()
	menuAccessService := menuservicesimpl.NewMenuAccessService(menuAccessRepository, userRepository, db)
	menuListService := menuservicesimpl.NewMenuListService(menuListRepository, db)
	menuUrlService := menuservicesimpl.NewMenuUrlService(menuUrlRepository, userRepository, db)
	menuAccessController := menucontrollers.NewMenuAccessController(menuAccessService)
	menuListController := menucontrollers.NewMenuListController(menuListService)
	menuURLController := menucontrollers.NewMenuURLController(menuUrlService)

	approvalRouter := ApprovalRouter(approvalController)
	approvalRequestRouter := ApprovalRequestRouter(approvalRequestController)
	approverRouter := ApproverRouter(approverController)
	authRouter := AuthRouter(authController)
	menuAccessRouter := MenuAccessRouter(menuAccessController)
	menuListRouter := MenuListRouter(menuListController)
	menuURLRouter := MenuURLRouter(menuURLController)
	userRouter := UserRouter(userController)
	binningRouter := BinningRouter(binningController)
	//Approval
	r.Mount("/approval", approvalRouter)
	r.Mount("/approval-request", approvalRequestRouter)
	r.Mount("/approver", approverRouter)

	//binning
	r.Mount("/api", binningRouter)
	//Menu Management
	r.Mount("/menu-access", menuAccessRouter)
	r.Mount("/menu-list", menuListRouter)
	r.Mount("/menu-url", menuURLRouter)

	//User Management
	r.Mount("/auth", authRouter)
	r.Mount("/user", userRouter)

	//binning

	return r
}
