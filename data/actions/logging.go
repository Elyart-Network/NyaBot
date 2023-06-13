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
				log.Warning("[Redis] Failed to disconnect from Redis: ", err)
			}
			log.Debug("[Redis] Pushed message to cache.")
		case false:
			// TODO
		}
	}
}

func (l *Logging) Insert(ctx context.Context, collection string, content any) {
	// Connect to MongoDB and insert message
	enableMDB := config.Get("logging.external").(bool)
	if enableMDB {
		m := handler.Mongo
		_, err := m.Database("NyaBot").Collection(collection).InsertOne(ctx, content)
		err = m.Disconnect(ctx)
		if err != nil {
			log.Warning("[MongoDB] Failed to disconnect from MongoDB: ", err)
		}
		log.Debug("[MongoDB] Inserted message to doc database.")
	} else {
		// Return when internal_log is not enabled
		internalLog := config.Get("logging.internal_log").(bool)
		if !internalLog {
			return
		}

		// Save Chat log to Database
		contentStr, err := json.Marshal(content)
		if err != nil {
			log.Warning("[Logging] Failed to Marshal content: ", err)
		}
		data := models.Logging{
			Collection: collection,
			Content:    string(contentStr),
		}
		handler.DB.Save(&data)
		log.Debug("[InternalLog] Inserted message to database.")
	}
}
