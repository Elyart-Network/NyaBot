package actions

import "github.com/Elyart-Network/NyaBot/data/drivers"

var handler Handler

type Handler struct {
	Sqlite  *drivers.SqliteClient
	MongoDB *drivers.MongoClient
	Redis   *drivers.RedisClient
}

func New(h Handler) {
	handler = h
}
