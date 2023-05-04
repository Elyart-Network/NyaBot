package logger

import (
	"log"
	"strconv"
)

func Warning(err error) {
	if err == nil {
		return
	}
	go log.Println("[Logger] Warning: " + err.Error())
}

func WarningStr(err string) {
	if err == "" {
		go log.Println("[Logger] Warning: empty string.")
		return
	}
	go log.Println("[Logger] Warning: " + err)
}

func Warningf(format string, args ...interface{}) {
	if format == "" {
		go log.Println("[Logger] Warning: empty format.")
		return
	}
	go func() {
		for k, arg := range args {
			log.Println("["+strconv.Itoa(k)+"] [Logger] Warning ("+format+"): ", arg)
		}
	}()
}
