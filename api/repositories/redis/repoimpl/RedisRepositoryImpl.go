package redisrepoimpl

import (
	"context"
	"fmt"
	"net/http"
	"user-services/api/config"
	"user-services/api/exceptions"
	"user-services/api/payloads"
	redisrepo "user-services/api/repositories/redis"
	"user-services/api/utils"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisRepositoryImpl struct {
}

func NewRedisRepository() redisrepo.RedisRepository {
	return &RedisRepositoryImpl{}
}

func (*RedisRepositoryImpl) GetSession(dbRedis *config.Database, userID int) (string, *exceptions.BaseErrorResponse) {

	getRedis := dbRedis.Client.HGet(context.Background(), fmt.Sprintf("session:%d", userID), "client")
	session, err := getRedis.Result()

	if session != "" || err == redis.Nil {
		return session, nil
	} else {
		return "", &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
}

func (*RedisRepositoryImpl) UpdateCredential(txRedis *config.Database, loginReq payloads.LoginCredential, userID int) (bool, *exceptions.BaseErrorResponse) {

	setRedis := txRedis.Client.HSet(
		context.Background(),
		fmt.Sprintf("session:%d", userID),
		"client", loginReq.Client,
		"ip_address", loginReq.IpAddress,
		"session", loginReq.Session,
		"user_id", userID,
	).Err()

	if setRedis != nil {
		logrus.Info(utils.UpdateDataFailed, "REDIS")
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}

	return true, nil
}

// DeleteCredential implements repositories.AuthRepository.
func (*RedisRepositoryImpl) DeleteCredential(tx *config.Database, userID int) (bool, *exceptions.BaseErrorResponse) {
	setRedis := tx.Client.HDel(
		context.Background(),
		fmt.Sprintf("session:%d", userID),
	).Err()
	if setRedis != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
		}
	}
	return true, nil
}
