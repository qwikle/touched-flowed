package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
	"touchedFlowed/features/utils"
)

var RedisInstance *redis.Client

type RedisDatabase struct {
	redis *redis.Client
}

func (r RedisDatabase) Sadd(key string, members ...string) error {
	return r.redis.SAdd(context.Background(), key, members).Err()
}

func (r RedisDatabase) Srem(key string, members ...string) error {
	return r.redis.SRem(context.Background(), key, members).Err()
}

func (r RedisDatabase) Smembers(key string) ([]string, error) {
	return r.redis.SMembers(context.Background(), key).Result()
}

func (r RedisDatabase) Delete(key string) error {
	return r.redis.Del(context.Background(), key).Err()
}

func RedisConnection() {
	if RedisInstance == nil {
		addr := os.Getenv("REDIS_HOST")
		if addr == "" {
			addr = "localhost:6379"
		}
		RedisInstance = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
	}
}

func (r RedisDatabase) Set(key, value string) error {
	return r.redis.Set(context.Background(), key, value, 0).Err()

}

func (r RedisDatabase) SetWithExpiration(key, value string, exp time.Duration) error {
	return r.redis.Set(context.Background(), key, value, exp).Err()
}

func (r RedisDatabase) Get(key string) (string, error) {
	return r.redis.Get(context.Background(), key).Result()
}

func NewRedisDatabase() utils.MemoryDatabase {
	return &RedisDatabase{redis: RedisInstance}
}
