package helper

import (
	"net/http"
	"user-services/api/exceptions"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) exceptions.BaseErrorResponse {
	err := recover()
	if err != nil {
		tx.Rollback()
		logrus.Info(err)
		return exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
		}
	} else {
		tx.Commit()
		return exceptions.BaseErrorResponse{}
	}
}
