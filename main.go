package main

import (
	"fmt"
	"ginDemo/pkg/logging"
	"github.com/fvbock/endless"
	"syscall"

	"ginDemo/pkg/setting"
	"ginDemo/routers"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		logging.Log.Printf("Server err: %v", err)
	}
}