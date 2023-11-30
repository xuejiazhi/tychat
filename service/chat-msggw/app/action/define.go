package action

import (
	"github.com/gorilla/websocket"
	"net/http"
	"tychat/util"
)

var Upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	Gpool *util.GoPool
	//SocketPool *SocketConPool
)

// SocketConPool  定义websocket
//type SocketConPool struct {
//	Conn   *websocket.Conn //连接
//	Locked *sync.RWMutex
//	Closed bool //是否关闭
//}

var WsClients = make(chan map[string]*websocket.Conn, 10000)

type LoginSession struct {
	UserID string `json:"userID"`
	Token  string `json:"token"`
}

var TestToken = map[string]LoginSession{
	"111-111-111-111": {
		"13520545568",
		"111-111-111-111",
	},
	"222-222-222-222": {
		"13520545568",
		"222-222-222-222",
	},
}
