package logger

import (
	"log"
	"strconv"
)

var debug bool

func SetDebug(b bool) {
	debug = b
}

func Debug(err error) {
	if !debug || err == nil {
		return
	}
	go log.Println("[Logger] Debug: " + err.Error())
}

func DebugStr(err string) {
	if !debug || err == "" {
		go log.Println("[Logger] Debug: empty string.")
		return
	}
	go log.Println("[Logger] Debug: " + err)
}

func Debugf(format string, args ...any) {
	if !debug || format == "" {
		go log.Println("[Logger] Debug: empty format.")
		return
	}
	go func() {
		for k, arg := range args {
			log.Println("["+strconv.Itoa(k)+"] [Logger] Debug ("+format+"): ", arg)
		}
	}()
}
