package drivers

import (
	"github.com/Elyart-Network/NyaBot/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

func SplitHostPort(url string) (host, port string) {
	if !strings.Contains(url, ":") {
		return url, "5432"
	}
	split := strings.Split(url, ":")
	return split[0], split[1]
}

func Postgres(dsn ExternalDSN) (*DBClient, error) {
	host, port := SplitHostPort(dsn.Host)
	conn := "host=" + host + " user=" + dsn.Username + " password=" + dsn.Password + " dbname=" + dsn.Name + " port=" + port + " sslmode=disable TimeZone=" + config.TZ()
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	return &DBClient{db}, err
}
