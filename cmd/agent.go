package main

import (
	"github.com/123shang60/image-load/pkg/register"
	"github.com/123shang60/image-load/pkg/router"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
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
		logrus.Panic(err)
	}
}
