package logger

import (
	"log"
	"strconv"
)

func Error(err error) {
	if err == nil {
		return
	}
	go log.Panicln("[Logger] Error: " + err.Error())
}

func ErrorStr(err string) {
	if err == "" {
		go log.Panicln("[Logger] Error: empty string.")
		return
	}
	go log.Panicln("[Logger] Error: " + err)
}

func Errorf(format string, args ...any) {
	if format == "" {
		go log.Panicln("[Logger] Error: empty format.")
		return
	}
	go func() {
		for k, arg := range args {
			log.Panicln("["+strconv.Itoa(k)+"] [Logger] Error ("+format+"): ", arg)
		}
	}()
}

func ErrorFatal(err error) {
	if err == nil {
		return
	}
	go log.Fatalln("[Logger] Error: " + err.Error())
}
