package action

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"tychat/util"
)

func Init() {
	initGoPool()
	go Handle()
	initSocketServer()

}

func initGoPool() (err error) {
	Gpool, err = util.NewGPool(&util.Options{
		MaxWorker:   1000,             // 最大的协程数
		MinWorker:   2,                // 最小的协程数
		JobBuffer:   10000,            // 缓冲队列的大小
		IdleTimeout: 30 * time.Second, // 协程的空闲超时退出时间
	})
	return err
}

func initSocketServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "(^v^) Hello, I'm TuiYun Chat!!!")
	})

	//起websocket服务
	r.GET("/ws", webSocket)

	//运行服务
	port := fmt.Sprintf(":%d", util.GlobalCfg.Msggw.Port)
	if err := r.Run(port); err == nil {
		fmt.Println("Run MsgGw Success")
	}
}

// WebSocket https://blog.csdn.net/qq_34857250/article/details/105122272/
func webSocket(c *gin.Context) {
	//websocket connection
	conn, err := Upgrade.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()
	//judge error
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	//获取token
	tokens := util.GetHeadInfos(c.Request, "token")
	if len(tokens) == 0 {
		fmt.Println("print logs")
		return
	}

	//check token
	token := tokens[0]
	if err := checkToken(token); err != nil {
		return
	}

	//conn channel
	WsClients <- map[string]*websocket.Conn{
		token: conn,
	}
	select {}

	////获取Header信息
	//func() (token string) {
	//	for k, v := range c.Request.Header {
	//		if strings.ToLower(k)
	//	}
	//}()
	//
	////检测header 是否有token
	//if err := checkHeader(func() (token string) {
	//	TK := c.Request.Header
	//	fmt.Println(TK)
	//	for s, strings := range TK {
	//		fmt.Println(s, strings)
	//	}
	//	if _, ok := c.Request.Header["token"]; ok {
	//		//get token
	//		token = cast.ToString(c.Request.Header["token"])
	//		return
	//	}
	//	//return 数据
	//	return
	//}()); err != nil {
	//
	//}

}

func checkToken(token string) (err error) {
	//judge token is nil
	if token == "" {
		err = errors.New("token is nil")
		return
	}
	//从verify中检验token的值
	//从token中获取相关登录的值
	//测试数据
	//if _,ok:=
	return
}

func Handle() {
	for tk := range WsClients {
		for k, v := range tk {
			go Write(k, v)
			go Read(k, v)
		}
	}
}

func Write(token string, wsConn *websocket.Conn) {
	defer wsConn.Close()
	for {

	}
}

func Read(token string, wsConn *websocket.Conn) {
	defer wsConn.Close()
	for {
		//读取数据
		msgType, msg, err := wsConn.ReadMessage()
		if err != nil || msgType == websocket.CloseMessage {
			break
		}
		fmt.Println("recv msg=>", msg)
	}
}
