package middlewares

import (
	"net/http"
	"user-services/api/config"
	"user-services/api/exceptions"
	redisrepoimpl "user-services/api/repositories/redis/repoimpl"
	userrepoimpl "user-services/api/repositories/user/repoimpl"
	"user-services/api/securities"
	redisserviceimpl "user-services/api/services/redis/serviceimpl"
)

func SetupCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Continue with the next middleware or handler
		next.ServeHTTP(w, r)
	})
}

func SetupAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		config.InitEnvConfigs(false, "")
		db := config.InitDB()
		dbRedis := config.InitRedisDB()
		authRepository := userrepoimpl.NewAuthRepository()
		userRepository := userrepoimpl.NewUserRepository()
		redisRepository := redisrepoimpl.NewRedisRepository()
		//User
		redisService := redisserviceimpl.NewRedisService(db, dbRedis, authRepository, userRepository, redisRepository)

		err := securities.GetAuthentication(request, redisService)
		if err != nil {
			exceptions.NewAuthorizationException(writer, request, &exceptions.BaseErrorResponse{
				Err: err,
			})
			return
		}
		// Continue with the next middleware or handler
		next.ServeHTTP(writer, request)
	})
}
