package actions

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/data/models"
	log "github.com/sirupsen/logrus"
)

type Plugin struct{}

func (p *Plugin) SaveConfig(UID int64, PName string, PType string, PConfig any, PAdd string) error {
	cfg, _ := json.Marshal(PConfig)
	data := models.Plugin{
		ID:       UID,
		Name:     PName,
		Type:     PType,
		Config:   string(cfg),
		Addition: PAdd,
	}
	err := CONN.DB.Save(&data).Error
	log.Debug("[DBAct](SaveConfig) Saved! @PName:", PName, " @PType:", PType, " @PConfig", PConfig, " @PAdd:", PAdd)
	return err
}

func (p *Plugin) GetConfig(PName string, PType string) (models.Plugin, error) {
	data := models.Plugin{Name: PName, Type: PType}
	err := CONN.DB.First(&data).Error
	log.Debug("[DBAct](GetConfig) Get Config. @PName:", PName, " @Type:", data.Type, " @Config:", data.Config, " @Addition:", data.Addition)
	return data, err
}

func (p *Plugin) DeleteConfig(PName string, PType string) error {
	data := models.Plugin{Name: PName, Type: PType}
	err := CONN.DB.Where("name = ? AND type = ?", PName, PType).Delete(&data).Error
	log.Debug("[DBAct](DeleteConfig) Deleted! @PName:", PName)
	return err
}
