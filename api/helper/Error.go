package helper

import (
	"net/http"
	"user-services/api/exceptions"
)

func ReturnError(writer http.ResponseWriter, request *http.Request, err *exceptions.BaseErrorResponse) {
	if err.StatusCode == http.StatusUnauthorized {
		exceptions.NewAuthorizationException(writer, request, err)
		return
	} else if err.StatusCode == http.StatusBadRequest {
		exceptions.NewBadRequestException(writer, request, err)
		return
	} else if err.StatusCode == http.StatusUnprocessableEntity {
		exceptions.NewEntityException(writer, request, err)
		return
	} else if err.StatusCode == http.StatusNotFound {
		exceptions.NewNotFoundException(writer, request, err)
		return
	} else if err.StatusCode == http.StatusForbidden {
		exceptions.NewRoleException(writer, request, err)
		return
	} else if err.StatusCode == http.StatusConflict {
		exceptions.NewConflictException(writer, request, err)
		return
	} else {
		exceptions.NewAppException(writer, request, err)
		return
	}
}
