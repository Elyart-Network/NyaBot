package drivers

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type RedisDSN struct {
	Hosts    []string
	Master   string
	Username string
	Password string
	DB       int
}

type RedisClient struct {
	redis.UniversalClient
}

func Redis(dsn RedisDSN) *RedisClient {
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
	if err != nil {
		log.Error("[Redis] Failed to ping Redis: ", err)
	}
	return &RedisClient{rdb}
}
