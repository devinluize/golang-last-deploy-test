package usertest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-services/api/config"
	masterentities "user-services/api/entities/master"

	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

func TestGetEmails(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	var emails []string
	users := []masterentities.User{}

	row, err := db.
		Select("email").
		Find(&users, []int{1, 2, 3, 4}).
		Rows()
	for _, email := range users {
		emails = append(emails, email.Email)
	}

	if err != nil {
		logrus.Info(err)

	}
	defer row.Close()

	fmt.Println(emails)
}

func TestGetCurrentUser(t *testing.T) {
	config.InitEnvConfigs(true, "")

	db := config.InitDB()
	dbRedis := config.InitRedisDB()
	authRouter, _ := setupRouter(db, dbRedis)

	// Create a strings.NewReader from the JSON data
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/auth/login", nil)
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
	// assert.Equal(t, 200, int(responseBody["status_code"].(float64)))
	assert.Equal(t, 302, int(responseBody["status_code"].(float64)))
	// assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
