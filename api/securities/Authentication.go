package securities

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	redisservices "user-services/api/services/redis"

	//redisservices "user-services/api/services/redis"
	"user-services/api/utils"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GetAuthentication(request *http.Request, service redisservices.RedisService) error {
	token, err := VerifyToken(request)
	if err != nil {
		return errors.New(utils.SessionError)
	}
	_, ok := token.Claims.(jwt.Claims)

	claims, _ := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	client := fmt.Sprintf("%v", claims["client"])

	id, _ := strconv.Atoi(userID)
	session, errs := service.GetSession(id)
	if errs != nil {
		return errors.New(utils.SessionError)
	}
	if client != session {
		return errors.New(utils.SessionError)
	}

	if !ok && !token.Valid {
		return errors.New(utils.SessionError)
	}

	return nil
}

func VerifyToken(request *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(request)
	if tokenString == "" {
		return nil, errors.New("Session Invalid, Please re-login")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("secret_key")), nil
	})
	if err != nil {
		return nil, errors.New("Session Invalid, Please re-login")
	}

	return token, nil
}

func ExtractToken(request *http.Request) string {
	// Get the query string parameters.
	keys := request.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token
	}

	// Get the Authorization header.
	authHeader := request.Header.Get("Authorization")

	// If the Authorization header is not empty, split it into two parts.
	if authHeader != "" {
		bearerToken := strings.Split(authHeader, " ")

		// If the Authorization header is in the format "Bearer token", return the token.
		if len(bearerToken) == 2 {
			return bearerToken[1]
		}
	}

	// If no token is found, return an empty string.
	return ""
}
