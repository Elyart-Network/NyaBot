package actions

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/data/models"
)

type Plugin struct{}

func (p *Plugin) SaveConfig(PName string, PType string, PConfig any, PAdd string) {
	cfg, _ := json.Marshal(PConfig)
	data := models.PluginConfig{
		Name:     PName,
		Type:     PType,
		Config:   string(cfg),
		Addition: PAdd,
	}
	dbType := config.Get("database.type").(string)
	switch dbType {
	case "sqlite":
		handler.Sqlite.Save(&data)
	}
}

func (p *Plugin) GetConfig(PName string) models.PluginConfig {
	data := models.PluginConfig{Name: PName}
	dbType := config.Get("database.type").(string)
	switch dbType {
	case "sqlite":
		handler.Sqlite.First(&data)
	}
	return data
}

func (p *Plugin) DeleteConfig(PName string) {
	data := models.PluginConfig{Name: PName}
	dbType := config.Get("database.type").(string)
	switch dbType {
	case "sqlite":
		handler.Sqlite.Delete(&data)
	}
}
