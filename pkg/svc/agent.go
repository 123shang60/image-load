package svc

import (
	"github.com/123shang60/image-load/pkg/dockerCore"
	"github.com/123shang60/image-load/pkg/s3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

func AgentLoad(c *gin.Context) {
	var s3File s3.S3File
	name := os.Getenv("name")
	if name == "" {
		name = "agent-local"
	}
	if err := c.ShouldBind(&s3File); err != nil {
		logrus.Error("请求无法解析！", err)
		c.JSON(200, LoadResult{
			Name: name,
			Code: 500,
			Data: err.Error(),
		})
		return
	}

	file, err := s3.Download(s3File)
	if err != nil {
		logrus.Error("文件获取失败！", err)
		c.JSON(200, LoadResult{
			Name: name,
			Code: 501,
			Data: err.Error(),
		})
		return
	}

	if err = dockerCore.LoadImage(file); err != nil {
		logrus.Error("镜像导入！", err)
		c.JSON(200, LoadResult{
			Name: name,
			Code: 502,
			Data: err.Error(),
		})
		return
	}

	os.Remove(file)

	c.JSON(200, LoadResult{
		Name: name,
		Code: 200,
		Data: "ok!",
	})
}
