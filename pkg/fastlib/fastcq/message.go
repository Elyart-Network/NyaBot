package fastcq

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/callback"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqcode"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
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
	SMsg    []any
	SFwd    []any
	Errors  []error
}

type MsgMiddleware func(*MessageFunc) error

func (c *MessageFunc) Message(callback callback.Full) *MessageFunc {
	switch callback.MessageType {
	case "group":
		c.IsGroup = true
		c.UserId = callback.UserID
		c.GroupId = callback.GroupID
	case "private":
		c.IsGroup = false
		c.UserId = callback.UserID
	}
	c.Msg = callback.MessageData
	c.MsgId = callback.MessageID
	return c
}

func (c *MessageFunc) Command(cmd string) *MessageFunc {
	c.Cmd = cmd
	return c
}

func (c *MessageFunc) Direct() *MessageFunc {
	c.Raw = true
	return c
}

func (c *MessageFunc) Send() *MessageFunc {
	if (c.Cmd == c.Msg) || c.Raw {
		var Id int64
		switch c.IsGroup {
		case true:
			Id = c.GroupId
		case false:
			Id = c.UserId
		}
		if len(c.SFwd) > 0 {
			_, err := SendForwardMsg(c.SFwd, Id, c.IsGroup)
			if err != nil {
				c.Errors = append(c.Errors, err)
			}
		}
		if len(c.SMsg) > 1 {
			for _, v := range c.SMsg {
				_, err := SendMsg(v.(string), Id, c.IsGroup)
				if err != nil {
					c.Errors = append(c.Errors, err)
				}
			}
		} else if len(c.SMsg) == 1 {
			_, err := SendMsg(c.SMsg[0].(string), Id, c.IsGroup)
			if err != nil {
				c.Errors = append(c.Errors, err)
			}
		}
	}
	return c
}

func (c *MessageFunc) Text(ctx string) *MessageFunc {
	c.SMsg = append(c.SMsg, ctx)
	return c
}

func (c *MessageFunc) Reply(ctx string) *MessageFunc {
	ReplyCodeData := types.ReplyData{ID: strconv.Itoa(int(c.MsgId))}
	ReplyCode := cqcode.Reply(ReplyCodeData)
	c.SMsg = append(c.SMsg, ReplyCode+ctx)
	return c
}

func (c *MessageFunc) IDFwd(id string) *MessageFunc {
	data := GenIdForward(id)
	c.SFwd = append(c.SFwd, data)
	return c
}

func (c *MessageFunc) CustomFwd(name string, id string, content string) *MessageFunc {
	data := GenCustomForward(name, id, content)
	c.SFwd = append(c.SFwd, data)
	return c
}

func (c *MessageFunc) Pic(url string) *MessageFunc {
	data := types.ImageData{
		File: "img.jpg",
		Url:  url,
	}
	c.SMsg = append(c.SMsg, cqcode.Image(data))
	return c
}

func (c *MessageFunc) Delete() *MessageFunc {
	if c.Cmd == c.Msg {
		err := DeleteMsg(c.MsgId)
		if err != nil {
			c.Errors = append(c.Errors, err)
		}
	}
	return c
}
