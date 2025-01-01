package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"log"
	"strconv"
	"strings"
)

func MakeSendCoreMessage(ctx *zero.Ctx) []byte {
	MessageReport := MessageReceive{
		Bot_id:      "zerobot",
		Bot_self_id: ctx.GetLoginInfo().Get("user_id").Raw,
		Msg_id:      strconv.FormatInt(message.NewMessageIDFromInteger(ctx.Event.MessageID.(int64)).ID(), 10),
		User_type:   CheckMessageType(ctx),
		Group_id:    GroupIsExit(strconv.FormatInt(ctx.Event.GroupID, 10)),
		User_id:     strconv.FormatInt(ctx.Event.UserID, 10),
		User_pm:     CheckUserPermission(ctx),
		Content: []Message{
			{
				Type: "text",
				Data: strings.Replace(ctx.MessageString(), "&amp;", "&", -1),
			},
		},
	}
	marshal, err := json.Marshal(MessageReport)
	if err != nil {
		return nil
	}
	return marshal
}

func ReadAndSendMessage(conn *websocket.Conn) {
	for {
	ReReadWs:
		_, p, err := conn.ReadMessage()
		if err != nil {
			ReConnectCore()
			conn = <-ConnChan
			ConnChan <- conn
			goto ReReadWs
		}
		messageStruct, err := ParseCoreMessage(p)
		if err != nil {
			log.Println(err.Error())
			return
		}
		if messageStruct.Content[0].Type == "log_INFO" {
			log.Println(messageStruct.Content)
		} else {
			zero.RangeBot(func(id int64, Ctx *zero.Ctx) bool {
				parseInt, err := strconv.ParseInt(messageStruct.TargetId, 10, 64)
				if err != nil {
					log.Println(err.Error())
					return false
				}
				switch messageStruct.Content[0].Type {
				case "text":
					if messageStruct.TargetType == "group" {
						Ctx.SendGroupMessage(parseInt, message.Text(messageStruct.Content[0].Data))
					} else {
						Ctx.SendPrivateMessage(parseInt, message.Text(messageStruct.Content[0].Data))
					}
				case "image":
					if messageStruct.TargetType == "group" {
						Ctx.SendGroupMessage(parseInt, message.Image(messageStruct.Content[0].Data.(string)))
					} else {
						Ctx.SendPrivateMessage(parseInt, message.Image(messageStruct.Content[0].Data.(string)))
					}
				case "node":
					if messageStruct.TargetType == "group" {
						Ctx.SendGroupForwardMessage(parseInt, fakeforwardnodemessage(messageStruct.Content[0].Data.([]interface{})))
					} else {
						Ctx.SendPrivateForwardMessage(parseInt, fakeforwardnodemessage(messageStruct.Content[0].Data.([]interface{})))
					}
				}
				return true
			})
		}

	}
}

// 此处代码由 @Jiang-Red 编写，感谢江林大佬
// https://github.com/Jiang-Red
func fakeforwardnodemessage(parsedata []interface{}) message.Message {
	msg := make(message.Message, len(parsedata))
	for i, v := range parsedata {
		m := v.(map[string]interface{})
		switch m["type"] {
		case "text":
			msg[i] = message.CustomNode("GSUID", 2854196310, m["data"])
		case "image":
			msg[i] = message.CustomNode("GSUID", 2854196310, message.Message{message.Image(m["data"].(string))})
		}
	}
	return msg
}
