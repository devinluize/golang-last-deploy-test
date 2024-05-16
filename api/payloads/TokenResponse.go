package payloads

import (
	"net/http"
	jsonresponse "user-services/api/helper/json/json-response"
)

func ResponseToken(writer http.ResponseWriter, data interface{}, message string, statusCode int) error {
	res := Respons{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	err := jsonresponse.WriteToResponseBody(writer, res)
	if err != nil {
		return err
	}
	return nil
}
