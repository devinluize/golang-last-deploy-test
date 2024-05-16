package menutest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-services/api/config"
	authcontrollers "user-services/api/controllers/auth"
	usercontrollers "user-services/api/controllers/user"
	menuentities "user-services/api/entities/master/menu"
	menupayloads "user-services/api/payloads/master/menu"
	redisrepoimpl "user-services/api/repositories/redis/repoimpl"
	userrepoimpl "user-services/api/repositories/user/repoimpl"
	"user-services/api/route"
	redisserviceimpl "user-services/api/services/redis/serviceimpl"
	userserviceimpl "user-services/api/services/user/serviceimpl"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/zeebo/assert"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB, dbRedis *config.Database) (chi.Router, chi.Router) {
	config.InitEnvConfigs(true, "")
	validate := validator.New()

	authRepository := userrepoimpl.NewAuthRepository()
	userRepository := userrepoimpl.NewUserRepository()
	redisRepository := redisrepoimpl.NewRedisRepository()

	authService := userserviceimpl.NewAuthService(db, dbRedis, authRepository, userRepository, redisRepository, validate)
	userService := userserviceimpl.NewUserService(userRepository, db, validate)
	redisService := redisserviceimpl.NewRedisService(db, dbRedis, authRepository, userRepository, redisRepository)

	authController := authcontrollers.NewAuthController(authService, userService, redisService)
	userController := usercontrollers.NewUserController(userService)

	authRouter := route.AuthRouter(authController)
	userRouter := route.UserRouter(userController)

	return authRouter, userRouter
}

// Test GetMenuByRoleID User
func TestGetMenuByRoleID(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)

	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/menu/access", nil)
	recorder := httptest.NewRecorder()

	authRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()

	// assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
}

// Test CheckCreateWithMenuListAndRole User
func TestGetMenuURLByCompanyAndUser(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	var menuAccess menuentities.MenuAccess
	var response []string
	err := db.
		Preload("MenuListByAccess.MenuUrl", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id",
				"path",
			)
		}).
		Select("id").
		Find(&menuAccess, "role_id = ?", 1).
		Error

	if err != nil {
		fmt.Println(response)
	}

	for i := range menuAccess.MenuListByAccess {
		response = append(response,
			menuAccess.MenuListByAccess[i].MenuUrl.Path,
		)
	}

	fmt.Println(response)
}

// Test CheckCreateWithMenuListAndRole User
func TestCheckCreateWithMenuListAndRole(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	var response []menupayloads.CheckCreateWithMenuListAndRoleResponse
	var menuAccess menuentities.MenuAccess
	err := db.
		Joins("MenuListByAccess", 1).
		Find(&menuAccess, "role_id = ?", 1).
		Error

	if err != nil {
		fmt.Println(err)
	}
	for _, menuList := range menuAccess.MenuListByAccess {
		response = append(response, menupayloads.CheckCreateWithMenuListAndRoleResponse{
			MenuListID: menuList.ID,
			RoleID:     menuAccess.RoleID,
		})
	}

	fmt.Println(response)
}

// Test GetMenuByRoleID User
func TestGetMenuAccessChildrenRepo(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	roleID := 1
	var response []menupayloads.CheckCreateWithMenuListAndRoleResponse
	var menuAccess menuentities.MenuAccess
	rows, err := db.
		Preload("MenuListByAccess.MenuUrl", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id menu_url_id",
				"path url",
			)
		}).
		Preload("MenuListByAccess", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id menu_list_id",
				"menu_url_id",
				"menu_id",
				"parent_id",
				"image menu_image",
			)
		}).
		Preload("MenuListByAccess.MenuMaster", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id",
				"title",
			)
		}).
		Find(&menuAccess, "role_id = ?", roleID).
		Scan(&response).
		Rows()

	if err != nil {
		fmt.Println(err)
	}
	for _, menuList := range menuAccess.MenuListByAccess {
		response = append(response, menupayloads.CheckCreateWithMenuListAndRoleResponse{
			MenuListID: menuList.ID,
			RoleID:     menuAccess.RoleID,
		})
	}
	defer rows.Close()
	fmt.Println(response)
}

// Test GetMenuByRoleID User
func TestIsMenuListDuplicate(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	var menuAccess menuentities.MenuAccess
	menuAccessRequest := menupayloads.CreateMenuAccessParamRequest{
		RoleID:     1,
		MenuListID: []int{1, 2},
	}

	err := db.
		Preload("MenuListByAccess", menuAccessRequest.MenuListID).
		Find(&menuAccess, "role_id = ?", menuAccessRequest.RoleID).
		Error

	if err != nil {
		fmt.Println(err)
		fmt.Println(menuAccess)
	}

	if menuAccess.ID != 0 {
		fmt.Println(err)
		fmt.Println(menuAccess)
	}

	fmt.Println(menuAccess)
}
