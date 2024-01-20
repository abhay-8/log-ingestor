package config

import (
	"context"
	"fmt"

	"github.com/abhay-8/log-ingestor/backend/database"
	"github.com/go-redis/redis/v8"
)

var ctx = context.TODO()

func GetFromCache(key string) (string, error) {
	data, err := database.RedisClient.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key does not exist in cache")
		}
		Logger.Warnw("Error retrieving data from cache", "Error:", err)
		return "", fmt.Errorf("error retrieving data from cache")

	}

	return data, nil
}
