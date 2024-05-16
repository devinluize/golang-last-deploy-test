package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"user-services/api/config"
	"user-services/api/exceptions"
)

func RouterMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//cookie, err := request.Cookie("token")
		//request.Header.Get("")
		//request.Header.Set("bearer", "THIS IS BEARER TOKEN PLEASE")
		//str := request.Header.Get("Authorization")

		//panic(dsada)
		// Split the string by space
		//parts := strings.Split(str, " ")

		// Check if there are at least two parts
		//if len(parts) >= 2 {
		//	tokenString := parts[1] // Index 1 corresponds to the second part
		//	//fmt.Println(secondString)
		//
		//} else {
		//	fmt.Println("Not enough parts in the string")
		//}
		writer.Header().Add("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		//if err != nil {
		//	if errors.Is(err, http.ErrNoCookie) {
		//		response = helper.ToAuthResponses(writer, "You are Unauthorized please login", http.StatusUnauthorized)
		//		err := encoder.Encode(response)
		//		helper.PanifIfError(err)
		//		return
		//	}
		//	response = helper.ToAuthResponses(writer, "Internal server error", http.StatusInternalServerError)
		//	err := encoder.Encode(response)
		//	helper.PanifIfError(err)
		//	return
		//}

		//tokenString := cookie.Value
		tokenString := request.Header.Get("Authorization")

		part := strings.Split(tokenString, " ")
		if len(part) >= 2 {
			tokenString = part[1]
		} else {
			exceptions.NewAuthorizationException(writer, request, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token not found from header",
				Data:       nil,
				Err:        gorm.ErrRecordNotFound,
			})
			return
		}

		claims := config.JWTClaim{}

		myToken, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		//helper.PanifIfError(err)
		if err != nil {
			exceptions.NewAuthorizationException(writer, request, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token is Invalid Or Expired",
				Data:       nil,
				Err:        err,
			})
			return
		}
		if claims.UserRole != 1 {
			exceptions.NewAuthorizationException(writer, request, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "You are not authorized to this endpoint",
				Data:       nil,
				Err:        err,
			})
			return
		}
		if !myToken.Valid {
			exceptions.NewAuthorizationException(writer, request, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token is not Valid",
				Data:       nil,
				Err:        err,
			})
			return
		}
		handler.ServeHTTP(writer, request)
	})
}
