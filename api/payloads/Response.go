package payloads

import (
	"net/http"
	jsonresponse "user-services/api/helper/json/json-response"
)

type ResponseAuth struct {
	Token   string `json:"token"`
	UserID  int    `json:"user_id"`
	Role    int    `json:"role"`
	Company int    `json:"company"`
}

type ResponseLogin struct {
	Token    string `json:"token"`
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type Respons struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func HandleSuccess(writer http.ResponseWriter, data interface{}, message string, status int) {
	res := Respons{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}

	jsonresponse.WriteToResponseBody(writer, res)
}
