package botcli

import (
	"github.com/Elyart-Network/NyaBot/internal/botcli/botcli/cmds"
	"github.com/Elyart-Network/NyaBot/pkg/fastlib"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
	"strconv"
)

func cqReturn(call interface{}, msg string) {
	data := call.(callback.Full)
	switch data.MessageType {
	case "group":
		atCode := models.AtData{QQ: strconv.FormatInt(data.UserID, 10)}
		fastlib.CqSendMsg(gocqhttp.At(atCode)+msg, data.GroupID, true)
	case "private":
		fastlib.CqSendMsg(msg, data.UserID, false)
	}
}

func cqProcessor(callback interface{}, cmd []string) {
	if len(cmd) < 1 {
		cqReturn(callback, "[BotCli] Invalided Command!")
		return
	}
	switch cmd[0] {
	case "send":
		if len(cmd[1:]) < 1 {
			cqReturn(callback, "[BotCli] Invalided Command!")
			return
		}
		switch cmd[1] {
		case "group":
			if len(cmd[2:]) < 2 {
				cqReturn(callback, "[BotCli] Invalided Command!")
				return
			}
			cqReturn(callback, cmds.CqSendGroupMessage(cmd[2:]))
		case "private":
			if len(cmd[2:]) < 2 {
				cqReturn(callback, "[BotCli] Invalided Command!")
				return
			}
			cqReturn(callback, cmds.CqSendPrivateMessage(cmd[2:]))
		}
	}
}

func discordProcessor(callback interface{}, cmd []string) {}

func telegramProcessor(callback interface{}, cmd []string) {}

func slackProcessor(callback interface{}, cmd []string) {}
