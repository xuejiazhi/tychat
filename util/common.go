package util

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
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

func PprofListen(port int) {
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
}

func GetHeadInfos(req *http.Request, key string) []string {
	for k, v := range req.Header {
		if strings.ToLower(k) == strings.ToLower(key) {
			return v
		}
	}
	return nil
}

func HandleError(err error) {
	fmt.Printf("Handle Error %v", err.Error())
	panic(err.Error())
}
