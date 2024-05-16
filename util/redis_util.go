package util

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheUtil struct {
	RDB *redis.Client
	Ctx context.Context
}

func (u *RedisCacheUtil) Get(key string) string {
	result, err := u.RDB.Get(u.Ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result
}

func (u *RedisCacheUtil) Set(key, value string, ttl time.Duration) {
	err := u.RDB.Set(u.Ctx, key, value, ttl).Err()
	if err != nil {
		panic(err)
	}
}
