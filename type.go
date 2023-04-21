package main

type BotConfig struct {
	CoreUrl             string  `description:"连接Core的Url"`
	CommandPrefix       string  `default:"/" description:"命令前缀"`
	SuperUsers          []int64 `description:"管理员的QQ号"`
	ConnectGoIsClient   bool    `description:"是否以正向Ws的连接模式连接go-cqhttp"`
	ConnectGoCqUrl      string  `description:"连接go-cqhttp的Url"`
	ConnectWaitn        int     `default:"16"`
	ConnectGoCqAccToken string  `description:"连接go-cqhttp的AccessToken"`
}

type Message struct {
	Type string      `default:"" json:"type"`
	Data interface{} `default:"" json:"data"`
}

type MessageReceive struct {
	Bot_id      string    `default:"zerobot" json:"bot_id"`                                          //Bot适配器类型，如onebot
	Bot_self_id string    `default:"" json:"bot_self_id"`                                            //Bot自身的QQ号
	Msg_id      string    `default:"" json:"msg_id"`                                                 //接受的消息id
	User_type   string    `default:"group" enum:"group,direct,channel,sub_channel" json:"user_type"` //消息类型 对应 群聊 私聊 频道 ？(未知)
	Group_id    string    `default:"" json:"group_id"`                                               //当消息类型为群聊消息时,此处应为群号
	User_id     string    `default:"" json:"user_id"`                                                //发送者QQ号
	User_pm     int       `default:"3" json:"user_pm"`                                               //用户权限等级，1为超级管理员，2为群管理/群主，3为普通用户
	Content     []Message `default:"" json:"content"`
}

type CoreReceiveMessage struct {
	BotId      string `json:"bot_id"`
	BotSelfId  string `json:"bot_self_id"`
	MsgId      string `json:"msg_id"`
	TargetType string `json:"target_type"`
	TargetId   string `json:"target_id"`
	Content    []Message
}
