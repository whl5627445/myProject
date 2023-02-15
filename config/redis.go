package config

import (
	"github.com/go-redis/redis/v8"
)

var addr = "redis:6379"

func openRedis() *redis.Client {
	if DEBUG != "" {
		addr = "192.168.121.199:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		//Addr:     "119.3.155.11:6379",
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

var R = openRedis()
