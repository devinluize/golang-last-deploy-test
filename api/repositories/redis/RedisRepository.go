package redisrepo

import (
	"user-services/api/config"
	"user-services/api/exceptions"
	"user-services/api/payloads"
)

type RedisRepository interface {
	GetSession(*config.Database, int) (string, *exceptions.BaseErrorResponse)
	UpdateCredential(*config.Database, payloads.LoginCredential, int) (bool, *exceptions.BaseErrorResponse)
	DeleteCredential(*config.Database, int) (bool, *exceptions.BaseErrorResponse)
}
