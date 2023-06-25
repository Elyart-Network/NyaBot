package actions

import (
	"github.com/Elyart-Network/NyaBot/data/drivers"
	"gorm.io/gorm"
)

var CONN Handler

type Handler struct {
	DB    *gorm.DB
	Mongo *drivers.MongoClient
	Redis *drivers.RedisClient
}

func New(h Handler) {
	CONN = h
}
