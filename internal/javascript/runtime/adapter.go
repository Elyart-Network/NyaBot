package runtime

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
)

func GoCqAdapter(vm *goja.Runtime, data CallbackData) {
	var message func(callback.Full)
	var request func(callback.Full)
	var notice func(callback.Full)
	var event func(callback.Full)

	// get functions from the javascript script
	err := vm.ExportTo(vm.Get("message"), &message)
	err = vm.ExportTo(vm.Get("request"), &request)
	err = vm.ExportTo(vm.Get("notice"), &notice)
	err = vm.ExportTo(vm.Get("event"), &event)
	if err != nil {
		log.Error("[JSVM] Error exporting functions from javascript script: ", err)
	}

	// send callback data to the javascript script
	switch data.CqCall.PostType {
	case "message":
		message(data.CqCall)
	case "request":
		request(data.CqCall)
	case "notice":
		notice(data.CqCall)
	case "event":
		event(data.CqCall)
	}
}
