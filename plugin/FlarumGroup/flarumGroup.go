package FlarumGroup

import (
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqapi/comMessage"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqcall"
	"github.com/Elyart-Network/NyaBot/core/gocqhttp/cqutil"
	"github.com/Elyart-Network/NyaBot/extend/plugin"
)

var PluginInfo = plugin.InfoStruct{
	Name:        "flarumGroup",
	Version:     "v1.0.0",
	Author:      "WhiteElytra",
	Description: "A Plugin for Flarum QQ Group",
	License:     "Apache-2.0",
	Homepage:    "http://example.com",
	Repository:  "http://example.com",
	Type:        "QQ",
}

func genSmData(name string, content string) cqutil.ForwardCustomNode {
	var data = cqutil.ForwardCustomData{
		Name:    name,
		Uin:     "113764782",
		Content: content,
	}
	var forward = cqutil.ForwardCustomNode{
		Type: "node",
		Data: data,
	}
	return forward
}

func Entry(callback cqcall.CallbackFull) {
	if (callback.GroupID != 432511748) && (callback.GroupID != 188723593) {
		return
	}
	if callback.Message == ".debug" {
		var messages = []cqutil.ForwardCustomNode{
			genSmData("超级练习册", "F12 看看控制台有无输出"),
			genSmData("超级练习册", "因为有可能是客户端毛病"),
			genSmData("超级练习册", "https://docs.flarum.org/zh/troubleshoot/"),
			genSmData("超级练习册", "你也可以打开 config.php 把 debug 改为 true"),
			genSmData("超级练习册", "尽量一次性提供足够的信息让群友们看"),
			genSmData("超级练习册", "我其实很讨厌说这套说辞。已经挂在了群公告里，但是我每天最少都会说一次"),
		}
		var data = comMessage.SendGroupForwardMsgData{
			GroupID:  callback.GroupID,
			Messages: messages,
		}
		_, err := comMessage.SendGroupForwardMsg(data)
		if err != nil {
			return
		}
	}
	if callback.Message == ".bt" {
		var data = comMessage.SendGroupMsgData{
			GroupID:    callback.GroupID,
			Message:    "这个群不是宝塔售后群，宝塔问题请加宝塔群问",
			AutoEscape: false,
		}
		_, err := comMessage.SendGroupMsg(data)
		if err != nil {
			return
		}
	}
	if callback.Message == ".ks" {
		var data = comMessage.SendGroupMsgData{
			GroupID:    callback.GroupID,
			Message:    "KubeSphere是本群唯一指定面板",
			AutoEscape: false,
		}
		_, err := comMessage.SendGroupMsg(data)
		if err != nil {
			return
		}
	}
	dict := map[string]string{
		"linux":   "Linux",
		"windows": "Windows",
		"mysql":   "MySQL",
		"unix":    "Unix",
		"macos":   "MacOS",
		"openbsd": "OpenBSD",
	}
	for key, value := range dict {
		if callback.Message == "."+key {
			var data = comMessage.SendGroupMsgData{
				GroupID:    callback.GroupID,
				Message:    "这个群不是" + value + "新手教学群",
				AutoEscape: false,
			}
			_, err := comMessage.SendGroupMsg(data)
			if err != nil {
				return
			}
		}
	}
	return
}
