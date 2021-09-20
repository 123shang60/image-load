package router

import (
	"github.com/123shang60/image-load/pkg/svc"
	"github.com/gin-gonic/gin"
)

func NewAgentRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/load", svc.AgentLoad)
	return r
}
