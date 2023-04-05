package cqcall

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func CallbackEncode(ctx interface{}, ws bool) (full CallbackFull, err error) {
	if ws {
		wsContext := ctx.([]byte)
		err = json.Unmarshal(wsContext, &full)
	} else {
		err = ctx.(*gin.Context).BindJSON(&full)
	}
	return
}
