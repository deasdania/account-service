package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
)

func InitDbRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_URL"),
		MinIdleConns: 20,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Panic(err)
		fmt.Print(err.Error())
	}
	return redisClient
}
