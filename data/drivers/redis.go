package drivers

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func Redis(dsn RedisDSN) (*RedisClient, error) {
	// Connect to Redis
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      dsn.Hosts,
		MasterName: dsn.Master,
		Username:   dsn.Username,
		Password:   dsn.Password,
		DB:         dsn.DB,
	})
	ctx := context.Background()
	// Ping Redis
	_, err := rdb.Ping(ctx).Result()
	return &RedisClient{rdb}, err
}
