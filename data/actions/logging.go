package actions

import (
	"context"
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/data/models"
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

func (l *Logging) Insert(ctx context.Context, collection string, content any) {
	// Connect to MongoDB and insert message
	enableMDB := config.Get("logging.external").(bool)
	if enableMDB {
		m := handler.MongoDB
		_, err := m.Database("NyaBot").Collection(collection).InsertOne(ctx, content)
		err = m.Disconnect(ctx)
		if err != nil {
			log.Warningf("[MongoDB] Failed to disconnect from MongoDB: %v", err)
		}
	} else {
		dbType := config.Get("database.type").(string)
		contentStr, _ := json.Marshal(content)
		data := models.Logging{
			Collection: collection,
			Content:    string(contentStr),
		}
		switch dbType {
		case "sqlite":
			handler.Sqlite.Save(&data)
		}
	}
}
