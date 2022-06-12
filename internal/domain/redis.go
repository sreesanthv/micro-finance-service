package domain

import (
	"time"

	"github.com/go-redis/redis"
)

type Redis interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
	Del(keys ...string) *redis.IntCmd
}
