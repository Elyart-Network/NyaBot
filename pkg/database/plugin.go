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
	orig := pluginComponent.GetConfig(PName, PType)

	// Save orig data to a map when key exists.
	if (orig != models.Plugin{}) {
		err := json.Unmarshal([]byte(orig.Config), &cfg)
		if err != nil {
			log.Error("[DBPlugin] Failed to Unmarshal JSON: ", err)
			return
		}
	}

	// Save CKey with CValue
	cfg[CKey] = CValue
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Error("[DBPlugin] Failed to Marshal JSON: ", err)
		return
	}
	pluginComponent.SaveConfig(PName, PType, data, "")
}

func GetConfig(PName string, PType string) map[string]any {
	// Prepare config map
	cfg := make(map[string]any)
	pluginComponent := actions.Plugin{}
	orig := pluginComponent.GetConfig(PName, PType)

	// Return when plugin is not exist
	if (orig == models.Plugin{}) {
		log.Debug("[DBPlugin] Getting a not existed plugin's config.")
		return nil
	}

	// Unmarshal config JSON
	err := json.Unmarshal([]byte(orig.Config), &cfg)
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
	orig := pluginComponent.GetConfig(PName, PType)

	// Return when plugin is not exist
	if (orig == models.Plugin{}) {
		log.Debug("[DBPlugin] Getting a not existed plugin's config.")
		return
	}

	// Unmarshal config JSON
	err := json.Unmarshal([]byte(orig.Config), &cfg)
	if err != nil {
		log.Error("[DBPlugin] Failed to Marshal JSON: ", err)
		return
	}

	// Remove CKey from map
	delete(cfg, CKey)

	// Update config
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Error("[DBPlugin] Failed to Marshal JSON: ", err)
		return
	}
	pluginComponent.SaveConfig(PName, PType, data, "")
}

func DelPlugin(PName string, PType string) {
	// Delete plugin from table
	pluginComponent := actions.Plugin{}
	pluginComponent.DeleteConfig(PName, PType)
}
