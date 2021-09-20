package router

import (
	"github.com/123shang60/image-load/pkg/register"
	"github.com/gin-gonic/gin"
)

func NewServerRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/regist", register.RegistNewNode)
	r.GET("/nodelist", register.NodeList)

	return r
}
