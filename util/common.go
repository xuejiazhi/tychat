package util

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	AdminApi = iota
	ChatApi
	ChatMsgGw
	DbService
)

// Signal sig notify
func Signal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigs
}

func PprofListen() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
}
