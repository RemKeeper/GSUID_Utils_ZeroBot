# GSUID_Utils_ZeroBot
GSUID-Core(早柚核心)的Zerobot适配版本



### 请保证您已经配置好GSUID-Core与Go-CQHTTP



在右边Releases中下载对应的版本启动

首次启动会自动退出，并创建BotConfig.json

您需要修改BotConfig.json

以下是一个示例

```json
{
   "CoreUrl": "ws://127.0.0.1:8765/ws/zerobot",
   "CommandPrefix": "",
   "SuperUsers": [
     114514191,
     191981000
   ],
   "ConnectGoIsClient": true,
   "ConnectGoCqUrl": "ws://127.0.0.1:11451",
   "ConnectWaitn": 16,
   "ConnectGoCqAccToken": ""
 }
```

各字段解释

CoreUrl：连接Core的Url，一般情况下，不建议修改/zerobot，这可能会导致Core平台判断失误

CommandPrefix:Bot命令前缀，不出意外的话，这个值应该不会生效，所以你最好留空

SuperUser:机器人管理员的QQ号，可以按照以上格式填入多个

ConnectGoIsClient:是否以正向模式向Go-CQHTTP发起连接，反向连接请设置为false

ConnectGoCqUrl:连接Go-CQHTTP的Url

ConnectWaitn:作为反向连接时的等待时间，不明白具体含义就别动，开发者也不明白

ConnectGoCqAccToken:连接Go-CQHTTP的AccToken，不懂留空



#### 修改完成后，如果配置正确，您再次启动后即可正常使用