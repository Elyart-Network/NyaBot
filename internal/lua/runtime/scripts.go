package runtime

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func init() {
	_, err := os.Stat("scripts")
	if os.IsNotExist(err) {
		log.Println("[Lua] Scripts folder not found, creating one...")
		err := os.Mkdir("scripts", os.ModePerm)
		if err != nil {
			log.Println("[Lua] Error creating scripts folder: ", err)
			return
		}
	} else if err != nil {
		log.Println("[Lua] Error checking scripts folder: ", err)
		return
	}
	_, err = os.Stat("scripts/lua.ini")
	if os.IsNotExist(err) {
		file, err := os.Create("scripts/lua.ini")
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println("[Lua] Error closing lua.ini: ", err)
				return
			}
		}(file)
		var defaultCnf = []byte("[example]\nenable = false\nscript = example.lua\n")
		_, err = file.Write(defaultCnf)
		if err != nil {
			log.Println("[Lua] Error writing lua.ini: ", err)
			return
		}
	} else if err != nil {
		log.Println("[Lua] Error checking lua.ini: ", err)
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
	clearScripts()
	cfg, err := ini.Load("scripts/lua.ini")
	if err != nil {
		log.Println("[Lua] Error loading lua.ini: ", err)
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
