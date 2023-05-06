package drivers

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/redis/go-redis/v9"
	"log"
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

func Redis(dbn int) *RedisClient {
	enable := config.Get("cache.enable").(bool)
	if !enable {
		return nil
	}
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      config.Get("cache.hosts").([]string),
		MasterName: config.Get("cache.master").(string),
		Username:   config.Get("cache.username").(string),
		Password:   config.Get("cache.password").(string),
		DB:         dbn,
	})
	ctx := context.Background()
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panicln("Redis ping error:", err)
	}
	log.Println("Redis ping:", ping)
	return &RedisClient{rdb}
}
