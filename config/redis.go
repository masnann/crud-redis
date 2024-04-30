package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func RedisInit(redisAddr, redisPassword string) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to ping Redis server: %v", err)
	}

	return nil
}

func RedisConnect() *redis.Client {
	return rdb
}
