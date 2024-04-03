package caching

import (
	"context"
	"flick_tickets/configs"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	config *configs.Configs
	client *redis.Client
}

func NewRedisDb(config *configs.Configs) *Redis {
	redisDb := &Redis{
		config: config,
	}
	err := redisDb.connect()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	log.Println("connected to Redis successfully")
	return redisDb
}

func (r *Redis) connect() error {
	address := r.config.AddressRedis
	if address == "" {
		return fmt.Errorf("empty Redis address")
	}
	index, _ := strconv.Atoi(r.config.DatabaseredisIndex)
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: r.config.PasswordRedis,
		DB:       index,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	r.client = client
	log.Println("connected to Redis successfully")
	return nil
}

func (r *Redis) CreateCollection() *redis.Client {
	return r.client
}
