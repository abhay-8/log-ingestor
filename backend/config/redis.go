package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abhay-8/log-ingestor/backend/database"
	"github.com/abhay-8/log-ingestor/backend/models"
	"github.com/abhay-8/log-ingestor/models"
	"github.com/redis/go-redis/v9"
)

var ctx = context.TODO()

func GetFromCache(key string) []models.Log {
	data, err := database.RedisClient.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return nil
		}
		Logger.Warnw("Error retrieving data from cache", "Error:", err)
		return nil

	}

	logs := []models.Log{}
	if err := json.Unmarshal([]byte(data), &logs); err != nil {
		Logger.Warnw("Error unmarshalling data from cache", "Error:", err)
		return nil
	}

	return logs
}

func SetInCache(key string, data []models.Log) error {
	data, err := json.Marshal(logs)
	if err != nil {
		Logger.Warnw("Error marshalling data from cache", "Error", err)
	}

	if err := database.RedisClient.Set(ctx, key, data, 0).Err(); err != nil {
		Logger.Warnw("Error setting data in cache", "Error:", err)
	}
	return nil
}

func RemovedFromCache(key string) error {
	err := database.RedisClient.Del(ctx, key).Err()
	if err != nil {
		if err == redis.Nil {
			return nil
		}

		Logger.Warnw("Error removing data from cache", "Error:", err)
		return fmt.Errorf("Error removing data from cache")
	}
	return nil
}
