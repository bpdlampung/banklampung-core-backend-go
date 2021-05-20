package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// Collections is redis's collection of function
type Collections interface {
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Conn(ctx context.Context) *redis.Conn
}
