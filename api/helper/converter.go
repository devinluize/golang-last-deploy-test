package helper

import (
	masterentities "user-services/api/entities/master"
	"user-services/api/payloads"
)

func ToDomainRegister(userInput payloads.RegisterRequest) masterentities.UserEntities {
	return masterentities.UserEntities{
		UserName:   userInput.UserName,
		UserEmail:  userInput.UserEmail,
		Password:   userInput.Password,
		UserRoleId: userInput.UserRoleId,
	}
}
