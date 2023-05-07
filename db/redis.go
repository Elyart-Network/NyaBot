package db

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/logger"
	"github.com/redis/go-redis/v9"
)

type RedisDSN struct {
	Hosts    []string
	Master   string
	Username string
	Password string
	DB       int
}

type RedisClient struct {
	db redis.UniversalClient
}

func Redis(dbn int) *RedisClient {
	// Convert []interface{} to []string
	addrSlice := config.Get("cache.hosts").([]interface{})
	var address []string
	for _, v := range addrSlice {
		address = append(address, v.(string))
	}
	// Connect to Redis
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      address,
		MasterName: config.Get("cache.master").(string),
		Username:   config.Get("cache.username").(string),
		Password:   config.Get("cache.password").(string),
		DB:         dbn,
	})
	ctx := context.Background()
	// Ping Redis
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Errorf("Redis", "Failed to ping Redis", err)
	}
	return &RedisClient{db: rdb}
}
