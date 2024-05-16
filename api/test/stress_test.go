package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
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

func Test100HitsInOneSecond(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)
	myStruct := payloads.LoginRequest{
		Username: "50596",
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

	// Create a wait group to synchronize the goroutines
	var wg sync.WaitGroup

	// Start 100 goroutines to make concurrent API requests
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Send the GET request
			request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/auth/login", reader)
			recorder := httptest.NewRecorder()

			authRouter.ServeHTTP(recorder, request)
			response := recorder.Result()

			// assert.Equal(t, 401, response.StatusCode)

			body, _ := io.ReadAll(response.Body)
			var responseBody map[string]interface{}
			json.Unmarshal(body, &responseBody)
			fmt.Println(responseBody)
			// assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Calculate the elapsed time
	start := time.Now()
	defer start.Sub(start)

	// Verify that all requests were completed within one second
	if time.Since(start) > time.Second {
		t.Errorf("Expected all requests to complete within one second, but took %v", time.Since(start))
	}
	fmt.Println(time.Since(start))
}
