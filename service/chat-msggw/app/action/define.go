package action

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsClients 定义websocket
var WsClients = make(map[string]*websocket.Conn)

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
