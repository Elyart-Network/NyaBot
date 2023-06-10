package database

import (
	"encoding/json"
	"github.com/Elyart-Network/NyaBot/data/actions"
	"github.com/Elyart-Network/NyaBot/data/models"
	log "github.com/sirupsen/logrus"
)

func SetConfig(PName string, PType string, CKey string, CValue any) {
	// Prepare config map
	cfg := make(map[string]any)
	pluginComponent := actions.Plugin{}
	orig, err := pluginComponent.GetConfig(PName, PType)
	if err != nil {
		log.Debug("[DBPlugin] Get config event caught err: ", err)
	}

	// Save orig data to a map when key exists.
	if err == nil {
		log.Debug("[DBPlugin] Getting a existed plugin's config.")
		err := json.Unmarshal([]byte(orig.Config), &cfg)
		if err != nil {
			log.Error("[DBPlugin] Failed to Unmarshal JSON: ", err)
			return
		}
	}

	// Save CKey with CValue
	cfg[CKey] = CValue
	err = pluginComponent.SaveConfig(orig.ID, PName, PType, cfg, "")
	if err != nil {
		log.Error("[DBPlugin] Error saving config to DB: ", err)
		return
	}
}

func GetConfig(PName string, PType string) map[string]any {
	// Prepare config map
	cfg := make(map[string]any)
	pluginComponent := actions.Plugin{}
	orig, err := pluginComponent.GetConfig(PName, PType)
	if err != nil {
		log.Error("[DBPlugin] Error getting config from DB: ", err)
		return nil
	}

	// Return when plugin is not exist
	if (orig == models.Plugin{
		Name: PName,
		Type: PType,
	}) {
		log.Debug("[DBPlugin] Getting a not existed plugin's config.")
		return nil
	}

	// Unmarshal config JSON
	err = json.Unmarshal([]byte(orig.Config), &cfg)
	if err != nil {
		log.Error("[DBPlugin] Failed to Marshal JSON: ", err)
		return nil
	}
	return cfg
}

func DelConfig(PName string, PType string, CKey string) {
	// Prepare config map
	cfg := make(map[string]any)
	pluginComponent := actions.Plugin{}
	orig, err := pluginComponent.GetConfig(PName, PType)
	if err != nil {
		log.Error("[DBPlugin] Error getting config from DB: ", err)
		return
	}

	// Return when plugin is not exist
	if (orig == models.Plugin{
		Name: PName,
		Type: PType,
	}) {
		log.Debug("[DBPlugin] Getting a not existed plugin's config.")
		return
	}

	// Unmarshal config JSON
	err = json.Unmarshal([]byte(orig.Config), &cfg)
	if err != nil {
		log.Error("[DBPlugin] Failed to Marshal JSON: ", err)
		return
	}

	// Remove CKey from map
	delete(cfg, CKey)
	if len(cfg) == 0 {
		DelPlugin(PName, PType)
		return
	}

	// Update config
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Error("[DBPlugin] Failed to Marshal JSON: ", err)
		return
	}
	err = pluginComponent.SaveConfig(orig.ID, PName, PType, data, "")
	if err != nil {
		log.Error("[DBPlugin] Error saving config to DB: ", err)
		return
	}
}

func DelPlugin(PName string, PType string) {
	// Delete plugin from table
	pluginComponent := actions.Plugin{}
	err := pluginComponent.DeleteConfig(PName, PType)
	if err != nil {
		log.Error("[DBPlugin] Error deleting plugin from DB: ", err)
		return
	}
}
