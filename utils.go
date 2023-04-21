package main

import (
	"encoding/json"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func CheckMessageType(ctx *zero.Ctx) string {
	if zero.OnlyGroup(ctx) {
		return "group"
	} else {
		return "direct"
	}
}

func CheckUserPermission(ctx *zero.Ctx) int {
	switch {
	case zero.SuperUserPermission(ctx):
		return 1
	case zero.OwnerPermission(ctx):
		return 2
	case zero.AdminPermission(ctx):
		return 3
	default:
		return 6
	}
}

func ParseCoreMessage(ReceiveMessage []byte) (CoreReceiveMessage, error) {
	var RecMessageStrutc CoreReceiveMessage
	err := json.Unmarshal(ReceiveMessage, &RecMessageStrutc)
	if err != nil {
		return CoreReceiveMessage{}, err
	}
	return RecMessageStrutc, nil
}
func GroupIsExit(GroupId string) string {
	if GroupId == "0" {
		return ""
	} else {
		return GroupId
	}
}
