package pubsub

import (
	"context"
	"fmt"
	goRedis "github.com/bpdlampung/banklampung-core-backend-go/databases/redis"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"github.com/go-redis/redis/v8"
)

type redisPubSub struct {
	rc     *redis.Client
	logger logs.Collections
}

type redisSubscription struct {
	p *redis.PubSub
}

func (r *redisPubSub) Publish(ctx context.Context, topic string, data []byte) error {
	_, err := r.rc.Publish(ctx, topic, data).Result()
	r.logger.Info(fmt.Sprintf("Publish with topic: %s - %s", topic, string(data)))

	return err
}

func (r *redisPubSub) Subscribe(ctx context.Context, name string, topic string, cb MsgHandler) (Subscription, error) {
	p := r.rc.Subscribe(ctx, topic)
	msgChan := p.Channel()
	go func() {
		for redMsg := range msgChan {
			msg := &Msg{
				Topic: redMsg.Channel,
				Data:  []byte(redMsg.Payload),
			}
			_ = cb(msg)
		}
	}()
	return &redisSubscription{p: p}, nil
}

func (r *redisSubscription) Unsubscribe() error {
	err := r.p.Close()
	return err
}

func RedisPubSub(rc goRedis.Redis) PubSub {
	return &redisPubSub{
		rc:     rc.GetRedisClient(),
		logger: rc.GetRedisLogger(),
	}
}
