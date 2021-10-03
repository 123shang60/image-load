package router

import (
	"github.com/123shang60/image-load/pkg/register"
	"github.com/123shang60/image-load/pkg/svc"
	"github.com/gin-gonic/gin"
)

func NewServerRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/nodelist", register.NodeList)

	r.POST("/load", svc.ServerLoad)

	return r
}
