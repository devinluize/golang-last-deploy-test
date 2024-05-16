package securities

import (
	"net/http"
	"time"
	"user-services/api/exceptions"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const AuthTokenValidTime = time.Hour * 24

func GenerateToken(username string, userID int, role int, companyID int, client string) (string, *exceptions.BaseErrorResponse) {
	timer := time.Now().Add(AuthTokenValidTime).Unix()
	secret := viper.GetString("secret_key")
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = timer
	claims["client"] = client
	claims["user_id"] = userID
	claims["role"] = role
	claims["company_id"] = companyID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		logrus.Info(err)
		return tokenString, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return tokenString, nil
}
