package drivers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDSN struct {
	MongoUri string
}

type MongoClient struct {
	*mongo.Client
}

func MongoDB(dsn MongoDSN) *MongoClient {
	// Connect to MongoDB
	ctx := context.Background()
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn.MongoUri))
	if err != nil {
		log.Errorf("[MongoDB] Failed to connect to MongoDB: %v", err)
	}
	err = conn.Ping(ctx, nil)
	if err != nil {
		log.Errorf("[MongoDB] Failed to ping MongoDB: %v", err)
	}
	return &MongoClient{conn}
}
