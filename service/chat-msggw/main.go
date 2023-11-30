package main

import (
	_ "net/http/pprof"
	"tychat/service/chat-msggw/app"
	"tychat/util"
)

// chat message gateway
// 消息网关
func main() {
	//pprof
	util.PprofListen(6060)
	//open new message gateway
	app.GateWayServer()
	//sig notify
	util.Signal()
}
