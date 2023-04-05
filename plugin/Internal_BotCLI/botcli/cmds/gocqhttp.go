package cmds

import (
	"github.com/Elyart-Network/NyaBot/extend/fastlib"
	"strconv"
)

func CqSendGroupMessage(cmd []string) string {
	GroupID, err := strconv.ParseInt(cmd[0], 10, 64)
	if err != nil {
		return "[BotCli] Invalid GroupID!"
	}
	if fastlib.CqSendMsg(cmd[1], GroupID, true) != 0 {
		return "[BotCli] Message sent!"
	} else {
		return "[BotCli] Message failed to send!"
	}
}

func CqSendPrivateMessage(cmd []string) string {
	UserID, err := strconv.ParseInt(cmd[0], 10, 64)
	if err != nil {
		return "[BotCli] Invalid UserID!"
	}
	if fastlib.CqSendMsg(cmd[1], UserID, false) != 0 {
		return "[BotCli] Message sent!"
	} else {
		return "[BotCli] Message failed to send!"
	}
}
