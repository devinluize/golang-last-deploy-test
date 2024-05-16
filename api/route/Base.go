package route

import (
	authcontrollers "user-services/api/controllers/auth"
	"user-services/api/controllers/binning"
	approvalcontrollers "user-services/api/controllers/master/approval"
	approvercontrollers "user-services/api/controllers/master/approver"
	menucontrollers "user-services/api/controllers/master/menu"
	approvalrequestcontrollers "user-services/api/controllers/transaction"
	usercontrollers "user-services/api/controllers/user"
	"user-services/api/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func UserRouter(
	userController usercontrollers.UserController,
) chi.Router {
	router := chi.NewRouter()

	// Apply the CORS middleware to all routes
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middlewares.SetupAuthMiddleware)
	router.Use(middleware.Recoverer)

	router.Get("/info", userController.GetCurrentUser)
	router.Get("/", userController.FindUser)
	router.Get("/by-username/{username}", userController.GetUserIDByUsername)
	router.Get("/{user_id}", userController.GetUsernameByUserID)

	return router
}
func ApprovalRouter(
	approvalController approvalcontrollers.ApprovalController,
) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middlewares.SetupAuthMiddleware)
	router.Use(middleware.Recoverer)

	router.Get("/", approvalController.GetAll)
	router.Get("/{id}", approvalController.Get)
	router.Post("/", approvalController.Create)
	router.Put("/{id}", approvalController.Update)

	return router
}
func ApprovalRequestRouter(
	approvalRequestController approvalrequestcontrollers.ApprovalRequestController,
) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middlewares.SetupAuthMiddleware)
	router.Use(middleware.Recoverer)

	router.Get("/by-user-type/{user_type}", approvalRequestController.GetApprovalRequest)
	router.Get("/by-user-type/{user_type}/{approval_id}", approvalRequestController.GetApprovalRequestDetails)
	router.Post("/", approvalRequestController.CreateApprovalRequest)
	router.Post("/detail", approvalRequestController.CreateApprovalRequestDetail)

	return router
}
func BinningRouter(
	approvalRequestController binning.BinningController,
) chi.Router {
	router := chi.NewRouter()
	//router.Use(middlewares.SetupCorsMiddleware)
	//router.Use(middlewares.SetupAuthMiddleware)
	//router.Use(middleware.Recoverer)

	//router.With(middlewares.RouterMiddleware).Post("/binning/getAll", approvalRequestController.GetAll)
	router.With(middlewares.RouterMiddleware).Post("/binning/getAll", approvalRequestController.GetAll)

	//router.Get("/by-user-type/{user_type}/{approval_id}", approvalRequestController.GetApprovalRequestDetails)
	//router.Post("/", approvalRequestController.CreateApprovalRequest)
	//router.Post("/detail", approvalRequestController.CreateApprovalRequestDetail)

	return router
}
func ApproverRouter(
	approverController approvercontrollers.ApproverController,
) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middlewares.SetupAuthMiddleware)
	router.Use(middleware.Recoverer)

	router.Get("/", approverController.GetAll)
	router.Get("/{id}", approverController.Get)
	router.Post("/", approverController.Create)
	router.Put("/{id}", approverController.Update)
	router.Delete("/{id}", approverController.Delete)

	return router
}

func AuthRouter(
	authController authcontrollers.AuthController,
) chi.Router {
	router := chi.NewRouter()
	router.Post("/register/{role}", authController.Register)
	router.Post("/loginAuth", authController.AuthLogin)

	//router.Use(middlewares.SetupCorsMiddleware)
	//router.Use(middleware.Recoverer)
	//router.Group(func(r chi.Router) {
	//	r.Post("/verify", authController.VerifyOTP)
	//	r.Post("/login", authController.Login)
	//	r.Post("/password/forgot", authController.ForgotPassword)
	//	r.Patch("/password/reset/{reset_token}", authController.ResetPassword)
	//})
	//router.Group(func(r chi.Router) {
	//	r.Use(middlewares.SetupAuthMiddleware)
	//	r.Post("/register", authController.Register)
	//	r.Post("/generate", authController.GenerateOTP)
	//	r.Post("/logout", authController.Logout)
	//	r.Put("/password/change", authController.ChangePassword)
	//})

	return router
}

func MenuAccessRouter(
	menuAccessController menucontrollers.MenuAccessController,
) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middlewares.SetupAuthMiddleware)
	router.Use(middleware.Recoverer)

	router.Get("/{id}", menuAccessController.Get)
	router.Get("/by-list/{menu_list_id}", menuAccessController.CheckByFilter)
	router.Post("/", menuAccessController.Create)
	router.Delete("/{menu_access_id}", menuAccessController.Delete)
	return router
}

func MenuURLRouter(
	menuURLController menucontrollers.MenuURLController,
) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middlewares.SetupAuthMiddleware)
	router.Use(middleware.Recoverer)

	router.Get("/by-company/{id}", menuURLController.GetByCompanyAndUser)
	router.Post("/", menuURLController.Create)
	return router
}

func MenuListRouter(
	menuListController menucontrollers.MenuListController,
) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middlewares.SetupAuthMiddleware)
	router.Use(middleware.Recoverer)

	router.Get("/dropdown", menuListController.GetDropdown)
	router.Post("/", menuListController.Create)
	return router
}
