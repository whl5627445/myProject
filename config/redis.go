package config

import (
	"github.com/go-redis/redis/v8"
)

func openRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "119.3.155.11:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

var R = openRedis()
