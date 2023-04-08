package fastcq

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
	"log"
	"strconv"
)

type MessageFunc struct {
	Cmd     string
	Raw     bool
	UserId  int64
	GroupId int64
	Msg     string
	MsgId   int32
	IsGroup bool
}

type MsgMiddleware func(*MessageFunc) error

func callbackDecode(callback callback.Full, message *MessageFunc) {
	switch callback.MessageType {
	case "group":
		message.IsGroup = true
		message.UserId = callback.UserID
		message.GroupId = callback.GroupID
	case "private":
		message.IsGroup = false
		message.UserId = callback.UserID
	}
	message.Msg = callback.MessageData
	message.MsgId = callback.MessageID
}

func (c *MessageFunc) Message(callback callback.Full, middlewares ...MsgMiddleware) {
	callbackDecode(callback, c)
	for _, middleware := range middlewares {
		err := middleware(c)
		if err != nil {
			log.Println("Middleware Error: ", err)
		}
	}
	return
}

func (c *MessageFunc) Command(cmd string) MsgMiddleware {
	return func(c *MessageFunc) error {
		c.Cmd = cmd
		return nil
	}
}

func (c *MessageFunc) Direct() MsgMiddleware {
	return func(c *MessageFunc) error {
		c.Raw = true
		return nil
	}
}

func (c *MessageFunc) Send(ctx string) MsgMiddleware {
	return func(c *MessageFunc) error {
		if (c.Cmd == c.Msg) || c.Raw {
			switch c.IsGroup {
			case true:
				_, err := SendMsg(ctx, c.GroupId, true)
				if err != nil {
					return err
				}
			case false:
				_, err := SendMsg(ctx, c.UserId, false)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func (c *MessageFunc) Reply(ctx string) MsgMiddleware {
	return func(c *MessageFunc) error {
		if (c.Cmd == c.Msg) || c.Raw {
			ReplyCodeData := types.ReplyData{ID: strconv.Itoa(int(c.MsgId))}
			ReplyCode := gocqhttp.Reply(ReplyCodeData)
			switch c.IsGroup {
			case true:
				_, err := SendMsg(ReplyCode+ctx, c.GroupId, true)
				if err != nil {
					return err
				}
			case false:
				_, err := SendMsg(ReplyCode+ctx, c.UserId, false)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func (c *MessageFunc) Forward(messages interface{}) MsgMiddleware {
	return func(c *MessageFunc) error {
		if (c.Cmd == c.Msg) || c.Raw {
			switch c.IsGroup {
			case true:
				_, err := SendForwardMsg(messages, c.GroupId, true)
				if err != nil {
					return err
				}
			case false:
				_, err := SendForwardMsg(messages, c.UserId, false)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func (c *MessageFunc) Delete() MsgMiddleware {
	return func(c *MessageFunc) error {
		if c.Cmd == c.Msg {
			err := DeleteMsg(c.MsgId)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
