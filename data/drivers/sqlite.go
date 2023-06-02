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
		log.Errorf("[Sqlite] Failed to open database: %v", err)
	}
	return &SqliteClient{db}
}

func (s *SqliteClient) Init(tables ...any) {
	err := s.AutoMigrate(tables)
	if err != nil {
		log.Errorf("[Sqlite] Error migrating tables: %v", err)
	}
}
