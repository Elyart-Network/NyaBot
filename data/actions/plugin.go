package actions

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/data/models"
	log "github.com/sirupsen/logrus"
)

type Plugin struct{}

func (p *Plugin) SaveConfig(PName string, PType string, PConfig any, PAdd string) {
	cfg, _ := json.Marshal(PConfig)
	data := models.Plugin{
		Name:     PName,
		Type:     PType,
		Config:   string(cfg),
		Addition: PAdd,
	}
	handler.DB.Save(&data)
	log.Debug("[DBAct](SaveConfig) Saved! @PName:", PName, " @PType:", PType, " @PConfig", PConfig, " @PAdd:", PAdd)
}

func (p *Plugin) GetConfig(PName string) models.Plugin {
	data := models.Plugin{Name: PName}
	handler.DB.First(&data)
	log.Debug("[DBAct](GetConfig) Get Config. @PName:", PName, " @Type:", data.Type, " @Config:", data.Config, " @Addition:", data.Addition)
	return data
}

func (p *Plugin) DeleteConfig(PName string) {
	data := models.Plugin{Name: PName}
	handler.DB.Delete(&data)
	log.Debug("[DBAct](DeleteConfig) Deleted! @PName:", PName)
}
