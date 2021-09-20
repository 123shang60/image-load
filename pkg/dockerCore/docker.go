package dockerCore

import (
	"context"
	"encoding/json"
	"os"

	"github.com/123shang60/image-load/pkg/common"

	"github.com/docker/docker/client"
)

var (
	cli *client.Client
)

func Init() {
	client2, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		common.Logger().Panic("agent 启动失败！docker cli 无法创建！", err)
	}

	cli = client2
}

func LoadImage(image string) error {
	ctx := context.Background()
	file, err := os.OpenFile(image, os.O_RDONLY, 0600)
	if err != nil {
		common.Logger().Error("读取已下载镜像文件失败！", err)
		return err
	}
	defer file.Close()

	resp, err := cli.ImageLoad(ctx, file, false)
	if err != nil {
		common.Logger().Error("镜像加载失败！", err)
		return err
	}
	byte, _ := json.Marshal(resp)
	common.Logger().Info("镜像加载成功！", string(byte))
	return nil
}
