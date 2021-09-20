package main

import (
	"github.com/123shang60/image-load/pkg/common"
	"os"

	"github.com/123shang60/image-load/pkg/dockerCore"
	"github.com/123shang60/image-load/pkg/register"
	"github.com/123shang60/image-load/pkg/router"
)

func main() {
	// 初始化 docker cli
	dockerCore.Init()
	// 定时注册启动
	go register.RegistAgent()
	addr := os.Getenv("addr")
	if addr == "" {
		addr = "127.0.0.1"
	}
	port := os.Getenv("port")
	if port == "" {
		port = "8081"
	}
	r := router.NewAgentRouter()
	err := r.Run(addr + ":" + port)
	if err != nil {
		common.Logger().Panic(err)
	}
}
