package health

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Health bool              `json:"health"`
	Status map[string]string `json:"status,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
	Config any               `json:"config,omitempty"`
}

func Raw(ctx context.Context, enc bool) Response {
	status, errors := Check(ctx)
	config := Config(enc)
	log.Debug("[Gin] Health Checked!")
	return Response{
		Health: len(errors) == 0,
		Status: status,
		Errors: errors,
		Config: config,
	}
}

func Gin(ctx *gin.Context) {
	resp, _ := json.Marshal(Raw(ctx, true))
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.String(200, string(resp))
}
