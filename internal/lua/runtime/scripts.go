package runtime

import (
	"github.com/Elyart-Network/NyaBot/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
)

func init() {
	_, err := os.Stat(luaDir)
	if os.IsNotExist(err) {
		log.Infoln("[Lua] Scripts folder not found, creating one...")
		err := os.Mkdir(luaDir, os.ModePerm)
		if err != nil {
			log.Error("[Lua] Error creating scripts folder: ", err)
			return
		}
	} else if err != nil {
		log.Error("[Lua] Error checking scripts folder: ", err)
		return
	}
	_, err = os.Stat(luaDir + "/lua.ini")
	if os.IsNotExist(err) {
		file, err := os.Create(luaDir + "/lua.ini")
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Error("[Lua] Error closing lua.ini: ", err)
				return
			}
		}(file)
		var defaultCnf = []byte("[example]\nenable = false\nscript = example.lua\n")
		_, err = file.Write(defaultCnf)
		if err != nil {
			log.Error("[Lua] Error writing lua.ini: ", err)
			return
		}
	} else if err != nil {
		log.Error("[Lua] Error checking lua.ini: ", err)
		return
	}
}

type luaScript struct {
	Name     string
	Enable   bool
	FileName string
}

var scripts []luaScript

func GetScripts() {
	enable := config.Get("plugin.lua_enable").(bool)
	if !enable {
		return
	}
	clearScripts()
	cfg, err := ini.Load(luaDir + "/lua.ini")
	if err != nil {
		log.Error("[Lua] Error loading lua.ini: ", err)
		return
	}
	for _, section := range cfg.Sections() {
		scripts = append(scripts, luaScript{
			Name:     section.Name(),
			Enable:   section.Key("enable").MustBool(false),
			FileName: section.Key("script").String(),
		})
	}
}

func clearScripts() {
	scripts = nil
}
