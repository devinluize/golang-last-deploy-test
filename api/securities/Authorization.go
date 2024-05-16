package securities

import (
	"fmt"
	"net/http"
	"strconv"
	"user-services/api/payloads"

	"github.com/golang-jwt/jwt/v4"
)

func ExtractAuthToken(request *http.Request) (*payloads.UserDetail, error) {
	token, err := VerifyToken(request)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userID := fmt.Sprintf("%v", claims["user_id"])
		username := fmt.Sprintf("%s", claims["username"])
		authorized := fmt.Sprintf("%s", claims["authorized"])
		role := fmt.Sprintf("%v", claims["role"])
		companyID := fmt.Sprintf("%s", claims["company_id"])
		ipAddress := fmt.Sprintf("%s", claims["ip_address"])
		client := fmt.Sprintf("%s", claims["client"])

		roles, _ := strconv.Atoi(role)
		userIDs, _ := strconv.Atoi(userID)

		authDetail := payloads.UserDetail{
			UserID:    int(userIDs),
			Username:  username,
			Authorize: authorized,
			Role:      roles,
			CompanyID: companyID,
			Client:    client,
			IpAddress: ipAddress,
		}

		return &authDetail, nil
	}

	return nil, nil
}
