package drivers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql(dsn ExternalDSN) (*DBClient, error) {
	conn := dsn.Username + ":" + dsn.Password + "@tcp(" + dsn.Host + ")/" + dsn.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	return &DBClient{db}, err
}
