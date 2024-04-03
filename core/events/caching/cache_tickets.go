package caching

import (
	"context"
	"encoding/json"
	"flick_tickets/common/log"
	"flick_tickets/configs"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	config *configs.Configs
}

func NewRedisCache(
	client *redis.Client,
	config *configs.Configs,
) *RedisCache {

	redisCache := &RedisCache{
		config: config,
	}
	err := redisCache.connectRedis()

	if err != nil {
		log.Error(err, "error")
		return nil
	}
	return redisCache
}

func (c *RedisCache) connectRedis() error {
	addr := c.config.AddressRedis
	password := c.config.PasswordRedis
	dbIndex, _ := strconv.Atoi(c.config.DatabaseRedis)
	c.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbIndex,
	})

	_, err := c.client.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return nil
}

func (c *RedisCache) SetObjectById(ctx context.Context, key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = c.client.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set data in Redis: %w", err)
	}

	return nil
}

func (c *RedisCache) GetObjectById(ctx context.Context, key string, objectUseConvert interface{}) (interface{}, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get data from Redis: %w", err)
	}

	err = json.Unmarshal([]byte(val), objectUseConvert)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return objectUseConvert, nil
}
