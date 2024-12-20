package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// RedisClient is an alias for the Redis client
type RedisClient = redis.Client

func ConnectCache() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisAddr := redisHost + ":" + redisPort

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	})

	_, err := client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return client
}
