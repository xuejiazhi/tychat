package app

import (
	"fmt"
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
	action.Init()
}
