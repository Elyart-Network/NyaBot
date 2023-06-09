package drivers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Sqlite(dsn SqliteDSN) (*DBClient, error) {
	db, err := gorm.Open(sqlite.Open(dsn.DB), &gorm.Config{})
	return &DBClient{db}, err
}
