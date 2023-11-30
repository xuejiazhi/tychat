package app

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"tychat/util"
)

type Args struct {
	A, B int
}

type MathService struct{}

func (m *MathService) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func RpcServer() {
	//解析配置文件
	if err := util.PhraseConfig(); err != nil {
		util.HandleError(err)
	}

	mathService := new(MathService)
	rpc.Register(mathService)
	t := util.GlobalCfg
	fmt.Println(t)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", util.GlobalCfg.Dbs.Port))
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	go rpc.Accept(listener)
	select {}
}
