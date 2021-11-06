package db

import (
	"log"

	"github.com/go-redis/redis"
)

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, _ := client.Ping().Result()
	log.Println(pong)
	return client
}
