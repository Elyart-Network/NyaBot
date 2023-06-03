package drivers

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDSN struct {
	DB string
}

type SqliteClient struct {
	*gorm.DB
}

func Sqlite(dsn SqliteDSN) *SqliteClient {
	db, err := gorm.Open(sqlite.Open(dsn.DB), &gorm.Config{})
	if err != nil {
		log.Error("[Sqlite] Failed to open database: ", err)
	}
	return &SqliteClient{db}
}
