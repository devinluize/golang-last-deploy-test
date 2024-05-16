package config

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	goredis "github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Database struct {
	Client *goredis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func InitDB() *gorm.DB {
	val := url.Values{}
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf(`%s://%s:%s@%s:%v?database=%s`, EnvConfigs.DBDriver, EnvConfigs.DBUser, EnvConfigs.DBPass, EnvConfigs.DBHost, EnvConfigs.DBPort, EnvConfigs.DBName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "dbo.", // schema name
			SingularTable: false,
		}})
	if err != nil {
		log.Fatal("Cannot connected database ", err)
		return nil
	}

	sqlDB, _ := db.DB()

	err = sqlDB.Ping()

	if err != nil {
		log.Fatal("Request Timeout ", err)
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	log.Info("Connected Database " + EnvConfigs.DBDriver + "- Connected on : " + EnvConfigs.ClientOrigin)

	return db
}

func InitRedisDB() *Database {
	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%v", EnvConfigs.RedisHost, EnvConfigs.RedisPort),
		Username: EnvConfigs.RedisUsername,
		Password: EnvConfigs.RedisPassword,
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		log.Fatal("Redis Error", err)
		return nil
	}
	log.Info("Connected Redis Database ")

	return &Database{
		Client: client,
	}
}
