package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tychat/middle"
)

type Args struct {
	A, B int
}

// Login 注册
func Login(c *gin.Context) {
	//获取
	username := c.PostForm("username")
	password := c.PostForm("password")
	//
	var reply int
	fmt.Println(username, password)
	err := middle.RpcClient.Call("MathService.Add", Args{A: 1, B: 2}, &reply)
	if err != nil {
		log.Fatal("Call MathService.Add error:", err)
	}
	fmt.Println(reply)
}
