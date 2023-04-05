package botcli

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcode"
	"github.com/Elyart-Network/NyaBot/extend/fastlib"
	"github.com/Elyart-Network/NyaBot/plugin/Internal_BotCLI/botcli/cmds"
	"strconv"
)

func cqReturn(callback interface{}, msg string) {
	data := callback.(cqcall.CallbackFull)
	switch data.MessageType {
	case "group":
		atCode := cqcode.AtData{QQ: strconv.FormatInt(data.UserID, 10)}
		fastlib.CqSendMsg(cqcode.At(atCode)+msg, data.GroupID, true)
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
