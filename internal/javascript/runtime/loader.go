package runtime

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/webhook"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

type CallbackData struct {
	WhCall webhook.Data
	CqCall callback.Full
}

type Plugin struct {
	Name     string
	Enable   bool
	FileName string
}

var (
	pluginDir = "plugin" //TODO
	plugins   []Plugin
)

func init() {
	_, err := os.Stat(pluginDir)
	if os.IsNotExist(err) {
		log.Infoln("[JavaScript] Plugin folder not found, creating one...")
		err := os.Mkdir(pluginDir, os.ModePerm)
		if err != nil {
			log.Error("[JavaScript] Error creating plugin folder: ", err)
			return
		}
	} else if err != nil {
		log.Error("[JavaScript] Error checking plugin folder: ", err)
		return
	}
	_, err = os.Stat(pluginDir + "/plugin.ini")
	if os.IsNotExist(err) {
		file, err := os.Create(pluginDir + "/plugin.ini")
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Error("[JavaScript] Error closing plugin.ini: ", err)
				return
			}
		}(file)
		var defaultCnf = []byte("[example]\nenable = false\nplugin = example.js\n")
		_, err = file.Write(defaultCnf)
		if err != nil {
			log.Error("[JavaScript] Error writing plugin.ini: ", err)
			return
		}
	} else if err != nil {
		log.Error("[JavaScript] Error checking plugin.ini: ", err)
		return
	}
}

func GetPlugins() {
	enable := true //TODO
	if !enable {
		return
	}
	ClearPlugins()
	cfg, err := ini.Load(pluginDir + "/plugin.ini")
	if err != nil {
		log.Error("[JavaScript] Error loading plugin.ini: ", err)
		return
	}
	for _, section := range cfg.Sections() {
		plugins = append(plugins, Plugin{
			Name:     section.Name(),
			Enable:   section.Key("enable").MustBool(false),
			FileName: section.Key("plugin").String(),
		})
	}
}

func ClearPlugins() {
	plugins = nil
}

func LoadPlugin(data CallbackData) {
	GetPlugins()
	go func() {
		for _, plugin := range plugins {
			if strings.HasSuffix(plugin.FileName, ".lua") && plugin.Enable {
				JSVM(pluginDir+"/"+plugin.FileName, data)
				log.Debug("[JSVM] Executed javascript plugin. @Name:", plugin.Name)
			} else if plugin.Name == "DEFAULT" {
				continue
			}
		}
	}()
}

func CqLoader(call callback.Full) {
	var data = CallbackData{
		CqCall: call,
	}
	LoadPlugin(data)
}
