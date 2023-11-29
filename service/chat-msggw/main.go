package main

import (
	"tychat/service/chat-msggw/app"
	"tychat/util"
)

// chat message gateway
// 消息网关
func main() {
	//open new message gateway
	app.NewMsgGateWayServer()
	//sig notify
	util.Signal()
}
