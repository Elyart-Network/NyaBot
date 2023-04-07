package cmds

import (
	"github.com/Elyart-Network/NyaBot/pkg/fastlib/fastcq"
	"strconv"
)

func CqSendGroupMessage(cmd []string) string {
	GroupID, err := strconv.ParseInt(cmd[0], 10, 64)
	if err != nil {
		return "[BotCli] Invalid GroupID!"
	}
	if msgId, _ := fastcq.SendMsg(cmd[1], GroupID, true); msgId != 0 {
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
	if msgId, _ := fastcq.SendMsg(cmd[1], UserID, false); msgId != 0 {
		return "[BotCli] Message sent!"
	} else {
		return "[BotCli] Message failed to send!"
	}
}
