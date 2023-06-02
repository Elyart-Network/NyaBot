package data

import (
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/data/actions"
	"github.com/Elyart-Network/NyaBot/data/drivers"
	"github.com/Elyart-Network/NyaBot/data/models"
)

func init() {
	handler := actions.Handler{}

	// Init database
	dbType := config.Get("database.type").(string)
	switch dbType {
	case "sqlite":
		dsn := drivers.SqliteDSN{DB: "database.db"}
		handler.Sqlite = drivers.Sqlite(dsn)
		handler.Sqlite.Init(&models.Plugin{}, &models.Logging{})
	}

	// Init redis
	externalCache := config.Get("cache.external").(bool)
	if externalCache {
		// Convert []interface{} to []string
		addrSlice := config.Get("cache.hosts").([]interface{})
		var rdbAddress []string
		for _, v := range addrSlice {
			rdbAddress = append(rdbAddress, v.(string))
		}
		dsn := drivers.RedisDSN{
			Hosts:    rdbAddress,
			Master:   config.Get("redis.master").(string),
			Username: config.Get("redis.username").(string),
			Password: config.Get("redis.password").(string),
			DB:       0,
		}
		handler.Redis = drivers.Redis(dsn)
	}

	// Init MongoDB
	externalLogging := config.Get("logging.external").(bool)
	if externalLogging {
		dsn := drivers.MongoDSN{MongoUri: config.Get("logging.mongo_uri").(string)}
		handler.MongoDB = drivers.MongoDB(dsn)
	}

	// Initialize
	actions.New(handler)
}
