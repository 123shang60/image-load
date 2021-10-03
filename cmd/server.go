package main

import (
	"github.com/123shang60/image-load/pkg/common"
	"github.com/123shang60/image-load/pkg/register"
	data "github.com/123shang60/image-load/pkg/register/proto"
	"github.com/123shang60/image-load/pkg/router"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	httpAddr := os.Getenv("httpAddr")
	if httpAddr == "" {
		httpAddr = "0.0.0.0:8080"
	}
	rpcAddr := os.Getenv("rpcAddr")
	if rpcAddr == "" {
		rpcAddr = "0.0.0.0:8082"
	}

	rpcService := grpc.NewServer()
	data.RegisterRegisterServer(rpcService, &register.Service{})

	rpc,err := net.Listen("tcp",rpcAddr)
	if err != nil {
		common.Logger().Panic("grpc server 启动失败！",err)
	}
	go func() {
		err = rpcService.Serve(rpc)
		if err != nil {
			common.Logger().Panic("grpc server 启动失败！",err)
		}
		common.Logger().Info("grpc server listen:", rpcAddr)
	}()
	r := router.NewServerRouter()
	err = r.Run(httpAddr)
	if err != nil {
		common.Logger().Panic("http server 启动失败！", err)
	}
	common.Logger().Info("http server listen:",httpAddr)
}
