package actions

import (
	"github.com/Elyart-Network/NyaBot/data/drivers"
	"gorm.io/gorm"
)

var handler Handler

type Handler struct {
	DB    *gorm.DB
	Mongo *drivers.MongoClient
	Redis *drivers.RedisClient
}

func New(h Handler) {
	handler = h
}
