package main

import (
	"encoding/json"
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"os"
)

func GetBotConfig() (BotConfig, error) {
	file, readConfig := os.ReadFile("./BotConfig.json")
	if readConfig != nil {
		CreateConfigFile()
		return BotConfig{}, readConfig
	}
	var Config BotConfig
	uJsonErr := json.Unmarshal(file, &Config)
	if uJsonErr != nil {
		return BotConfig{}, uJsonErr
	}
	return Config, nil
}

func CreateConfigFile() {
	Config := BotConfig{
		CoreUrl:             "ws://127.0.0.1:8765/zerobot",
		CommandPrefix:       "",
		SuperUsers:          []int64{123456789, 114514191},
		ConnectGoIsClient:   true,
		ConnectGoCqUrl:      "ws://127.0.0.1:11451",
		ConnectWaitn:        16,
		ConnectGoCqAccToken: "",
	}
	MJson, CreFile := json.MarshalIndent(Config, " ", "  ")
	if CreFile != nil {
		return
	}
	err := os.WriteFile("./BotConfig.json", MJson, 0644)
	if err != nil {
		fmt.Println("文件创建失败，请检查权限")
		return
	}
	fmt.Println("文件创建成功，清修改BotConfig.json后重启")
	os.Exit(0)
}

func IsPositiveConnections(bool2 bool, Config BotConfig) zero.Driver {
	if bool2 {
		return driver.NewWebSocketClient(Config.ConnectGoCqUrl, Config.ConnectGoCqAccToken)
	} else {
		return driver.NewWebSocketServer(Config.ConnectWaitn, Config.ConnectGoCqUrl, Config.ConnectGoCqAccToken)
	}
}
