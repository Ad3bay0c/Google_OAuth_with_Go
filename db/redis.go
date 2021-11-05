package db

import (
	"github.com/go-redis/redis"
	"log"
)

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, _ := client.Ping().Result()
	log.Println(pong)
	return client
}