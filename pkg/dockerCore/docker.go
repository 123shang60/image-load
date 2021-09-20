package dockerCore

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	cli *client.Client
)

func Init() {
	client2, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Panic("agent 启动失败！docker cli 无法创建！", err)
	}

	cli = client2
}

func LoadImage(image string) error {
	ctx := context.Background()
	file, err := os.OpenFile(image, os.O_RDONLY, 0600)
	if err != nil {
		logrus.Error("读取已下载镜像文件失败！", err)
		return err
	}
	defer file.Close()

	resp, err := cli.ImageLoad(ctx, file, false)
	if err != nil {
		logrus.Error("镜像加载失败！", err)
		return err
	}
	byte, _ := json.Marshal(resp)
	logrus.Info("镜像加载成功！", string(byte))
	return nil
}
