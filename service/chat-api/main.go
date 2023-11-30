package main

import (
	"tychat/service/chat-api/app"
	"tychat/util"
)

//type Args struct {
//	A, B int
//}

func main() {
	//pprof
	util.PprofListen(6061)
	//run api server
	app.ApiServer()
	//sig notify
	util.Signal()

	//conn, err := rpc.Dial("tcp", "localhost:8080")
	//if err != nil {
	//	log.Fatal("Dial error:", err)
	//}
	//defer conn.Close()
	//var reply int
	//err = conn.Call("MathService.Add", Args{A: 1, B: 2}, &reply)
	//if err != nil {
	//	log.Fatal("Call MathService.Add error:", err)
	//}
	//fmt.Println(reply)
}
