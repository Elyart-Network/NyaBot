package logger

import (
	"log"
	"strconv"
)

func Info(err error) {
	if err == nil {
		return
	}
	go log.Println("[Logger] Info: " + err.Error())
}

func InfoStr(err string) {
	if err == "" {
		go log.Println("[Logger] Info: empty string.")
		return
	}
	go log.Println("[Logger] Info: " + err)
}

func Infof(format string, args ...any) {
	if format == "" {
		go log.Println("[Logger] Info: empty format.")
		return
	}
	go func() {
		for k, arg := range args {
			log.Println("["+strconv.Itoa(k)+"] [Logger] Info ("+format+"): ", arg)
		}
	}()
}
