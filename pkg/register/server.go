package register

import (
	"context"
	"github.com/123shang60/image-load/pkg/common"
	data "github.com/123shang60/image-load/pkg/register/proto"

	"github.com/gin-gonic/gin"
)

type Service struct {
	data.UnimplementedRegisterServer
}

func (*Service)RegistNode(_ context.Context,info *data.NodeInfo) (*data.RegistResp, error)  {
	err := AddCache(info)
	if err != nil {
		common.Logger().Error("注册新服务失败!",err)
		return &data.RegistResp{
			Err:  err.Error(),
			Code: 501,
		},nil
	}

	return &data.RegistResp{
		Err:  "ok!",
		Code: 200,
	}, nil
}

func NodeList(c *gin.Context) {
	c.JSON(200, GetCache())
}
