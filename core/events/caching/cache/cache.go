package cache

import (
	"context"
	"encoding/json"
	"flick_tickets/core/events/caching"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RepositoryCache interface {
	SetObjectById(ctx context.Context, key string, data interface{}) error
	GetObjectById(ctx context.Context, key string) (string, error)
	KeyExists(ctx context.Context, key string) (bool, error)
}

type Cache struct {
	client *redis.Client
}

func NewCache(ch *caching.Redis) RepositoryCache {
	return &Cache{
		client: ch.CreateCollection(),
	}
}

func (c *Cache) SetObjectById(ctx context.Context, key string, data interface{}) error {
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

func (c *Cache) GetObjectById(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", err
		}
		return "", fmt.Errorf("failed to get data from Redis: %w", err)
	}

	return val, nil
}

func (c *Cache) KeyExists(ctx context.Context, key string) (bool, error) {
	exists, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if key exists: %w", err)
	}
	return exists == 1, nil
}
