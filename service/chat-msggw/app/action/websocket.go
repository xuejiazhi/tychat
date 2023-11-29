package action

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

// https://blog.csdn.net/qq_34857250/article/details/105122272/
func WebSocket(c *gin.Context) {
	//websocket connection
	conn, err := Upgrade.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()
	//judge error
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	//检测header 是否有token
	if err := checkHeader(func() (token string) {
		if _, ok := c.Request.Header["token"]; ok {
			//get token
			token = cast.ToString(c.Request.Header["token"])
			return
		}
		//return 数据
		return
	}()); err != nil {

	}

}

func checkHeader(token string) (err error) {
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
