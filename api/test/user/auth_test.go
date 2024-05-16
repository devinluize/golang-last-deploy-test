package usertest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-services/api/config"
	authcontrollers "user-services/api/controllers/auth"
	usercontrollers "user-services/api/controllers/user"
	"user-services/api/payloads"
	redisrepoimpl "user-services/api/repositories/redis/repoimpl"
	userrepoimpl "user-services/api/repositories/user/repoimpl"
	"user-services/api/route"
	redisserviceimpl "user-services/api/services/redis/serviceimpl"
	userserviceimpl "user-services/api/services/user/serviceimpl"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/magiconair/properties/assert"
	"gorm.io/gorm"
)

var token string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjbGllbnQiOiIiLCJleHAiOjE3MDIyNzU4NjQsInVzZXJfaWQiOjIzMywidXNlcm5hbWUiOiI1MDU5NiJ9.OU3QHLvgHrsUE1UCZIg0bwRalvDrfhdL62gUNwEF5Sc"

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

func truncateUser(db *gorm.DB) {
	db.Exec("TRUNCATE users")
}

// Test Login User
func TestLoginUser(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)

	myStruct := payloads.LoginRequest{
		Username: "50596",
		Password: "string",
		Client:   "A",
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}

	// Create a strings.NewReader from the JSON data
	reader := strings.NewReader(string(jsonData))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/auth/login", reader)
	recorder := httptest.NewRecorder()

	authRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()

	// assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, 302, int(responseBody["status_code"].(float64)))
	// assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}

// Test Register User
func TestRegisterUser(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)

	myStruct := payloads.CreateRequest{
		Username: "69696969",
		Email:    "string@gmail.com",
		IsActive: true,
		Password: "string",
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}
	// Create a strings.NewReader from the JSON data
	reader := strings.NewReader(string(jsonData))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/auth/register", reader)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()
	authRouter.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
}

func TestVerifyOTP(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)

	myStruct := payloads.CreateRequest{
		Username: "50596",
		Email:    "string",
		IsActive: true,
		Password: "string",
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}
	// Create a strings.NewReader from the JSON data
	reader := strings.NewReader(string(jsonData))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/auth/verify", reader)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()
	authRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()
	// assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// assert.Equal(t, 201, int(responseBody["status_code"].(int)))
	// assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}

func TestGenerateOTP(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)

	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/auth/generate", nil)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()
	authRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()
	// assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// assert.Equal(t, 201, int(responseBody["status_code"].(int)))
	// assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}

func TestChangePassword(t *testing.T) {
	config.InitEnvConfigs(true, "")
	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)
	passwordRequest := payloads.ChangePasswordInput{
		OldPassword: "string",
		Password:    "string",
		NewPassword: "string",
	}

	jsonData, err := json.Marshal(passwordRequest)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}
	reader := strings.NewReader(string(jsonData))
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/auth/password/change", reader)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()
	authRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()
	// assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// assert.Equal(t, 201, int(responseBody["status_code"].(int)))
	// assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}

func TestForgotPassword(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)
	emailRequest := payloads.ForgotPasswordInput{
		Email: "test@gmail.com",
	}

	jsonData, err := json.Marshal(emailRequest)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}
	reader := strings.NewReader(string(jsonData))
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/auth/password/change", reader)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()
	authRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()
	// assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// assert.Equal(t, 201, int(responseBody["status_code"].(int)))
	// assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}

func TestResetPassword(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)
	emailRequest := payloads.ResetPasswordInput{
		Password:        "string",
		PasswordConfirm: "string",
	}

	jsonData, err := json.Marshal(emailRequest)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}
	reader := strings.NewReader(string(jsonData))
	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodPatch, "http://localhost:3000/auth/password/reset/testtoken", reader)
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()
	authRouter.ServeHTTP(recorder, request)
	// fmt.Println(request)
	response := recorder.Result()
	// assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// assert.Equal(t, 201, int(responseBody["status_code"].(int)))
	// assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
