package utils

import "time"

type MemoryDatabase interface {
	Set(key, value string) error
	SetWithExpiration(key, value string, exp time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Sadd(key string, members ...string) error
	Srem(key string, members ...string) error
	Smembers(key string) ([]string, error)
}
