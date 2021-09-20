package register

import (
	"encoding/json"
	"os"
	"time"

	"github.com/123shang60/image-load/pkg/common"
)

func RegistAgent() {
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
	registAddr := os.Getenv("registAddr")
	if registAddr == "" {
		registAddr = "http://127.0.0.1:8080"
	}
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

	nodeInof := NodeInfo{
		Name: name,
		Addr: addr,
		Port: port,
	}

	byte, err := json.Marshal(nodeInof)
	if err != nil {
		common.Logger().Error("构造环境信息失败！", err)
		return
	}

	res, err := common.DoJsonHttp(registAddr+"/regist", byte, "POST")
	if err != nil {
		common.Logger().Error("发送注册信息失败！", err)
		return
	}

	common.Logger().Info("接收到注册信息:", string(res))

	var rul RegistResult
	if err := json.Unmarshal(res, &rul); err != nil {
		common.Logger().Error("解析注册结果失败！", err)
	}
	if rul.Code == 200 {
		common.Logger().Debug("注册成功！")
		return
	}
}
