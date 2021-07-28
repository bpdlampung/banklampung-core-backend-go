package redis

import (
	"context"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Redis struct {
	client *redis.Client
	logger logs.Collections
}

var redisClient Redis

func InitConnection(redisDB, redisHost, redisPort, redisPassword string, logger logs.Collections) Redis {
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

	redisClient = Redis{
		client: client,
		logger: logger,
	}

	return redisClient
}

func GetClient() Collections {
	return redisClient
}
