package health

import (
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

func Gin(ctx *gin.Context) {
	status, errors := Check(ctx)
	config := Config()
	resp, _ := json.Marshal(Response{
		Health: len(errors) == 0,
		Status: status,
		Errors: errors,
		Config: config,
	})
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.String(200, string(resp))
	log.Debug("[Gin] Health Checked!")
}
