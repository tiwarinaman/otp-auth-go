package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"otp-auth/configs"
	"time"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis(config configs.RedisConfig) {

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DB,
	})

	// Test the connection
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}

func SetValue(key string, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

func GetValue(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}
