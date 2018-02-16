package dao

import (
	"github.com/go-redis/redis"
	"sync"
)

var client *redis.Client
var once sync.Once

func InitRedis() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	})
	return client
}
