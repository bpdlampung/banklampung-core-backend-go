package redis

import (
	"banklampung-core/logs"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var redisClient *redis.Client
var logger logs.Collections

func InitConnection(redisDB, redisHost, redisPort, redisPassword string, logger logs.Collections) Collections {
	db := 0

	parseRedisDb, err := strconv.ParseInt(redisDB, 10, 32)

	if err == nil {
		db = int(parseRedisDb)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", redisHost, redisPort),
		Password: redisPassword,
		DB:       db,
	})

	if client.Ping(context.Background()).Err() != nil {
		logger.Error("cannot connect redis")
		panic("cannot connect redis")
	}

	logger.Info("Redis Connected")

	redisClient = client
	logger = logger

	return redisClient
}
