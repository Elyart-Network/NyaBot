package actions

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	log "github.com/sirupsen/logrus"
)

type Logging struct{}

func (l *Logging) Cache(ctx context.Context, collection string, content any) {
	// Connect to Cache and insert message
	cacheNum := config.Get("logging.cache_num").(int)
	externalCache := config.Get("cache.external").(bool)
	if cacheNum > 0 {
		switch externalCache {
		case true:
			if handler.Redis == nil {
				return
			}
			// Connect to Redis and insert message
			r := handler.Redis
			clen := r.LLen(ctx, collection)
			if clen.Val() >= int64(cacheNum) {
				// Remove oldest message
				r.LPopCount(ctx, collection, 1)
			}
			r.RPush(ctx, collection, content)
			err := r.Close()
			if err != nil {
				log.Warningf("[Redis] Failed to disconnect from Redis: %v", err)
			}
		case false:
			// TODO
		}
	}
}

func (l *Logging) Mongo(ctx context.Context, collection string, content any) {
	if handler.MongoDB == nil {
		return
	}
	// Connect to MongoDB and insert message
	enableMDB := config.Get("logging.external").(bool)
	if enableMDB {
		m := handler.MongoDB
		_, err := m.Database("NyaBot").Collection(collection).InsertOne(ctx, content)
		err = m.Disconnect(ctx)
		if err != nil {
			log.Warningf("[MongoDB] Failed to disconnect from MongoDB: %v", err)
		}
	}
}
