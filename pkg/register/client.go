package register

import (
	"context"
	"encoding/json"
	data "github.com/123shang60/image-load/pkg/register/proto"
	"google.golang.org/grpc"
	"os"
	"time"

	"github.com/123shang60/image-load/pkg/common"
)

var client data.RegisterClient

func InitClient() {
	registAddr := os.Getenv("registAddr")
	if registAddr == "" {
		registAddr = "127.0.0.1:8082"
	}

	conn,err := grpc.Dial(registAddr, grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		common.Logger().Panic("agent grpc 客户端创建失败！",err)
	}
	client = data.NewRegisterClient(conn)
}

func RegistAgent() {
	InitClient()
	registAgent()
	tickTimer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tickTimer.C:
			registAgent()
		}
	}
}

func registAgent() {
	addr := os.Getenv("addr")
	if addr == "" {
		addr = "127.0.0.1"
	}
	port := os.Getenv("port")
	if port == "" {
		port = "8081"
	}
	name := os.Getenv("name")
	if name == "" {
		name = "agent-local"
	}

	nodeInfo := &data.NodeInfo{
		Name: name,
		Addr: addr,
		Port: port,
	}

	resp,err := client.RegistNode(context.Background(), nodeInfo)
	if err != nil {
		common.Logger().Error("grpc 调用失败！",err)
		return
	}

	common.Logger().Debug("grpc 注册结果：", resp)
	if resp.Code != 200 {
		common.Logger().Error("grpc 注册失败！",resp.Err)
		return
	}
	bytes,_ := json.Marshal(resp)
	common.Logger().Info("grpc 注册成功！",string(bytes))
}
