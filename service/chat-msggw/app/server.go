package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tychat/middle"
	"tychat/service/chat-msggw/app/action"
	"tychat/util"
)

func GateWayServer() (err error) {
	//解析配置文件
	if err = util.PhraseConfig(); err != nil {
		panic(fmt.Sprintf("Phrase Config error,%s", err.Error()))
	}

	//获取redis
	middle.InitRedis()
	//开始监听
	ServerDefault()
	return err
}

func ServerDefault() {
	//初始化gin
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "(^v^) Hello, I'm TuiYun Chat!!!")
	})

	//起websocket服务
	r.GET("/ws", action.WebSocket)

	//运行服务
	port := fmt.Sprintf(":%d", util.GlobalCfg.Msggw.Port)
	if err := r.Run(port); err == nil {
		fmt.Println("Run MsgGw Success")
	}
}
