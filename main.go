package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	zero "github.com/wdvxdr1123/ZeroBot"
	"log"
	"os"
	"time"
)

var ConnChan = make(chan *websocket.Conn, 1)

var Config BotConfig

func init() {
	Config, _ = GetBotConfig()
}

func main() {

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	conn, ConnErr := ConnectCore(Config.CoreUrl)
	if ConnErr != nil {
		log.Println("Core连接失败", ConnErr.Error())
		os.Exit(404)
	} else {
		log.Println("Core连接成功")
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	//Bot事务处理端
	zero.OnMessage().
		Handle(func(ctx *zero.Ctx) {
		ReSend:
			SendErr := SendWsMessage(MakeSendCoreMessage(ctx), conn)
			if SendErr != nil {
				conn = <-ConnChan
				log.Println(SendErr.Error())
				goto ReSend
				//ctx.SendChain(message.Text("发送数据到Core服务器错误"))
			}
		})
	zero.Run(&zero.Config{
		NickName:      []string{"GSUID"},
		CommandPrefix: Config.CommandPrefix,
		SuperUsers:    Config.SuperUsers,
		Driver: []zero.Driver{
			// 正向 WS
			//driver.NewWebSocketClient(Config.ConnectGoCqUrl, Config.ConnectGoCqAccToken),
			// 反向 WS
			//driver.NewWebSocketServer(16, "ws://127.0.0.1:6701", ""),
			IsPositiveConnections(Config.ConnectGoIsClient, Config),
		},
	})
	ReadAndSendMessage(conn)
}

func ReConnectCore() *websocket.Conn {
	for {
		log.Println("连接断开")
		log.Println("五秒后尝试重连")
		time.Sleep(time.Second * 2)
		conn, err := ConnectCore(Config.CoreUrl)
		if err != nil {
			log.Println("连接失败:", err.Error())
		} else {
			ConnChan <- conn
			fmt.Println("--------重连成功----------")
			return conn
		}
	}
}
