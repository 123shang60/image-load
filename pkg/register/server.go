package register

import "github.com/gin-gonic/gin"

func RegistNewNode(c *gin.Context) {
	var info NodeInfo
	if err := c.ShouldBind(&info);err != nil {
		c.JSON(200, RegistResult{
			Code: 500,
			Data: err.Error(),
		})
		return
	}

	if err := AddCache(info);err != nil {
		c.JSON(200, RegistResult{
			Code: 501,
			Data: err.Error(),
		})
		return
	}

	c.JSON(200, RegistResult{
		Code: 200,
		Data: "ok!",
	})
}

func NodeList(c *gin.Context) {
	c.JSON(200, GetCache())
}
