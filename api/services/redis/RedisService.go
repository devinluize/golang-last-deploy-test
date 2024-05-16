package redisservices

import (
	"user-services/api/exceptions"
	"user-services/api/payloads"
)

type RedisService interface {
	GetSession(int) (string, *exceptions.BaseErrorResponse)
	Login(payloads.LoginRequest, string) (
		*payloads.ResponseAuth,
		bool,
		*exceptions.BaseErrorResponse,
	)
	DeleteCredential(int) (bool, *exceptions.BaseErrorResponse)
}
