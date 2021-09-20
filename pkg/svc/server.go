package svc

import (
	"encoding/json"
	"sync"

	"github.com/123shang60/image-load/pkg/common"
	"github.com/123shang60/image-load/pkg/register"
	"github.com/123shang60/image-load/pkg/s3"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

func ServerLoad(c *gin.Context) {
	var s3File s3.S3File
	if err := c.ShouldBind(&s3File); err != nil {
		logrus.Error("请求无法解析！", err)
		c.JSON(200, []LoadResult{{
			Name: "server",
			Code: 500,
			Data: err.Error(),
		}})
		return
	}

	byte, err := json.Marshal(s3File)
	if err != nil {
		logrus.Error("请求无法解析！", err)
		c.JSON(200, []LoadResult{{
			Name: "server",
			Code: 500,
			Data: err.Error(),
		}})
		return
	}

	agentList := register.GetCache()
	result := make([]LoadResult, 0)
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	wg.Add(len(agentList))
	for _, value := range agentList {
		go func(value cache.Item) {
			defer wg.Done()
			node := value.Object.(register.NodeInfo)
			logrus.Debug("开始执行对应节点的导入工作:", node)
			res, err := common.DoJsonHttp("http://"+node.Addr+":"+node.Port+"/load", byte, "POST")
			if err != nil {
				logrus.Error("agent 执行镜像导入失败！", "node 名称：", node.Name, "失败原因：", err)
				lock.Lock()
				defer lock.Unlock()
				result = append(result, LoadResult{
					Name: node.Name,
					Code: 500,
					Data: err.Error(),
				})
				return
			}
			logrus.Info("agent 执行镜像导入结果！", "node 名称：", node.Name, "执行结果：", string(res))
			var rul LoadResult
			err = json.Unmarshal(res, &rul)
			if err != nil {
				logrus.Error("agent 执行镜像导入结果无法解析！", "node 名称：", node.Name, "失败原因：", err)
				lock.Lock()
				defer lock.Unlock()
				result = append(result, LoadResult{
					Name: node.Name,
					Code: 501,
					Data: err.Error(),
				})
				return
			}
			lock.Lock()
			defer lock.Unlock()
			result = append(result, rul)
		}(value)
	}
	wg.Wait()

	c.JSON(200, result)
}
