package callback

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Encode(ctx interface{}, ws bool) (full Full, err error) {
	if ws {
		wsContext := ctx.([]byte)
		err = json.Unmarshal(wsContext, &full)
	} else {
		err = ctx.(*gin.Context).BindJSON(&full)
	}
	return
}
