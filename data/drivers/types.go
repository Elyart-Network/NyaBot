package drivers

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type SqliteDSN struct {
	DB string
}

type ExternalDSN struct {
	Host     string
	Name     string
	Username string
	Password string
}

type RedisDSN struct {
	Hosts    []string
	Master   string
	Username string
	Password string
	DB       int
}

type MongoDSN struct {
	MongoUri string
}

type MongoClient struct {
	*mongo.Client
}

type RedisClient struct {
	redis.UniversalClient
}

type DBClient struct {
	*gorm.DB
}
