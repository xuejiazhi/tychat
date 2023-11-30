package action

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tychat/util"
)

// WebSocket https://blog.csdn.net/qq_34857250/article/details/105122272/
func WebSocket(c *gin.Context) {
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
	checkToken(tokens[0])

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
