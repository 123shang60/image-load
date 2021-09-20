package main

import (
	"os"

	"github.com/123shang60/image-load/pkg/common"
	"github.com/123shang60/image-load/pkg/router"
)

func main() {
	serverAddr := os.Getenv("serverAddr")
	if serverAddr == "" {
		serverAddr = "0.0.0.0:8080"
	}

	r := router.NewServerRouter()
	err := r.Run(serverAddr)
	if err != nil {
		common.Logger().Panic("server 启动失败！", err)
	}
}
