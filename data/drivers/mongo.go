package drivers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB(dsn MongoDSN) (*MongoClient, error) {
	// Connect to MongoDB
	ctx := context.Background()
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn.MongoUri))
	err = conn.Ping(ctx, nil)
	return &MongoClient{conn}, err
}
