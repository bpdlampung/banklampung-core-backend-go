package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

// Collections is redis's collection of function
type Collections interface {
	Set(key string, value interface{}, expired time.Duration) error
	Get(key string, value interface{}) error
	Del(key string)
	Scan(prefix string) ([]string, error)
}

func (r Redis) Set(key string, value interface{}, expired time.Duration) error {
	marshaledValue, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return r.client.Set(context.Background(), key, marshaledValue, expired).Err()
}

func (r Redis) Get(key string, value interface{}) error {
	val, err := r.client.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	}

	res := []byte(val)
	return json.Unmarshal(res, value)
}

func (r Redis) Del(key string) {
	r.client.Del(context.Background(), key)
}

func (r Redis) Scan(prefix string) ([]string, error) {
	var keys []string

	iterator := r.client.Scan(context.Background(), 0, prefix, 0).Iterator()

	for iterator.Next(context.Background()) {
		keys = append(keys, iterator.Val())
	}

	if err := iterator.Err(); err != nil {
		return nil, err
	}

	if keys == nil || len(keys) == 0 {
		return nil, nil
	}

	return keys, nil
}
