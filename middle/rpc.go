package middle

import (
	"fmt"
	"log"
	"net/rpc"
	"tychat/util"
)

func InitRpc() {
	var err error
	t := util.GlobalCfg
	fmt.Println(t)
	RpcClient, err = rpc.Dial("tcp",
		fmt.Sprintf("%s:%d",
			util.GlobalCfg.Dbs.IP,
			util.GlobalCfg.Dbs.Port))
	if err != nil {
		log.Fatal("Dial error:", err)
	}
}

func RpcClose() {
	if RpcClient != nil {
		RpcClient.Close()
	}
}
