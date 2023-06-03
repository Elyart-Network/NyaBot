package database

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func jsonDe(data []byte) map[string]any {
	var kv map[string]any
	err := json.Unmarshal(data, &kv)
	if err != nil {
		log.Warning("[DBPlugin] Error Unmarshal JSON: ", err)
		return nil
	}
	return kv
}
