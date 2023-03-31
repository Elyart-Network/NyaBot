package cqcall

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Callback(ctx *gin.Context) CallbackFull {
	full := CallbackFull{}
	err := ctx.BindJSON(&full)
	if err != nil {
		log.Println(err)
	}
	return full
}
