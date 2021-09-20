package main

import (
	"os"

	"github.com/123shang60/image-load/pkg/router"
	"github.com/sirupsen/logrus"
)

func main() {
	serverAddr := os.Getenv("serverAddr")
	if serverAddr == "" {
		serverAddr = "0.0.0.0:8080"
	}

	r := router.NewServerRouter()
	err := r.Run(serverAddr)
	if err != nil {
		logrus.Panic("server 启动失败！", err)
	}
}
