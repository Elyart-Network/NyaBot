package botcli

import (
	"github.com/Elyart-Network/NyaBot/internal/botcli/botcli/cmds"
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/fastcq"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
)

func cqReturn(call interface{}, msg string) {
	data := call.(callback.Full)
	c := fastcq.MessageFunc{}
	c.Message(data, c.Reply(msg))
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
