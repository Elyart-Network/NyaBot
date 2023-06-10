package data

import (
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/data/actions"
	"github.com/Elyart-Network/NyaBot/data/drivers"
	"github.com/Elyart-Network/NyaBot/data/models"
	log "github.com/sirupsen/logrus"
)

func init() {
	handler := actions.Handler{}
	var (
		db  *drivers.DBClient
		err error
	)

	// Init database
	dbType := config.Get("database.type").(string)
	exDSN := drivers.ExternalDSN{
		Host:     config.Get("database.host").(string),
		Name:     config.Get("database.name").(string),
		Username: config.Get("database.username").(string),
		Password: config.Get("database.password").(string),
	}
	switch dbType {
	default:
		liteDSN := drivers.SqliteDSN{DB: "database.db"}
		db, err = drivers.Sqlite(liteDSN)
	case "mysql":
		db, err = drivers.Mysql(exDSN)
	case "postgres":
		db, err = drivers.Postgres(exDSN)
	}
	if err != nil {
		log.Fatal("["+dbType+"] Error connecting to database: ", err)
	}
	handler.DB = db.DB

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
			Master:   config.Get("cache.master").(string),
			Username: config.Get("cache.username").(string),
			Password: config.Get("cache.password").(string),
			DB:       0,
		}
		conn, err := drivers.Redis(dsn)
		if err != nil {
			log.Fatal("[Redis] Error connecting to Redis: ", err)
		}
		handler.Redis = conn
	}

	// Init MongoDB
	externalLogging := config.Get("logging.external").(bool)
	if externalLogging {
		dsn := drivers.MongoDSN{MongoUri: config.Get("logging.mongo_uri").(string)}
		conn, err := drivers.MongoDB(dsn)
		if err != nil {
			log.Fatal("[MongoDB] Error connecting to MongoDB: ", err)
		}
		handler.Mongo = conn
	}

	// Initialize
	err = handler.DB.AutoMigrate(&models.Plugin{}, &models.Logging{})
	if err != nil {
		log.Error("[Database] Error migrating database: ", err)
	}
	actions.New(handler)
	log.Debug("[Database] Data sources initialized!")
}
