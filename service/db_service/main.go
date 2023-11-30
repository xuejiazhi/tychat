package main

import (
	app "tychat/service/db_service/app"
	"tychat/util"
)

func main() {
	//pprof
	util.PprofListen(6061)
	//run api server
	app.RpcServer()
	//sig notify
	util.Signal()
}
