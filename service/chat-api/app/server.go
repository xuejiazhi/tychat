package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tychat/middle"
	"tychat/service/chat-api/app/action"
	"tychat/util"
)

func ApiServer() {
	//解析配置文件
	if err := util.PhraseConfig(); err != nil {
		util.HandleError(err)
	}

	//链接RPC
	middle.InitRpc()

	//获取redis
	middle.InitRedis()

	//开始监听
	ServerDefault()
}

func ServerDefault() {
	//初始化gin
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "(^v^) Hello, I'm TuiYun Chat!!!")
	})

	//起websocket服务
	rg := r.Group("/chat")
	{
		rg.POST("/register", action.Register) //注册
		rg.POST("/login", action.Login)       //登录
		rg.GET("/userinfo", action.UserInfo)  //用户信息

	}

	//运行服务
	port := fmt.Sprintf(":%d", util.GlobalCfg.Chat_Api.Port)
	if err := r.Run(port); err == nil {
		fmt.Println("Run MsgGw Success")
	}
}
