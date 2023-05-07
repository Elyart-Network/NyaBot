package logger

import (
	"github.com/Elyart-Network/NyaBot/config"
	"log"
	"strconv"
)

func CheckDebug() bool {
	enable := config.Get("server.debug_mode").(bool)
	return enable
}

func Debug(err error) {
	if !CheckDebug() || err == nil {
		return
	}
	go log.Println("[Logger] Debug: " + err.Error())
}

func DebugStr(err string) {
	if !CheckDebug() || err == "" {
		go log.Println("[Logger] Debug: empty string.")
		return
	}
	go log.Println("[Logger] Debug: " + err)
}

func Debugf(format string, args ...any) {
	if !CheckDebug() || format == "" {
		go log.Println("[Logger] Debug: empty format.")
		return
	}
	go func() {
		for k, arg := range args {
			log.Println("["+strconv.Itoa(k)+"] [Logger] Debug ("+format+"): ", arg)
		}
	}()
}
