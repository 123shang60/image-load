package register

import (
	"github.com/123shang60/image-load/pkg/common"
	"time"

	"github.com/patrickmn/go-cache"
)

// 注册信息
var c *cache.Cache

func init() {
	c = cache.New(30*time.Second, 10*time.Second)
}

func AddCache(info NodeInfo) error {
	common.Logger().Debug("有一个 agent 前来注册。。。。", info)

	c.SetDefault(info.Name, info)

	common.Logger().Debug("服务注册成功！", info)
	return nil
}

func GetCache() map[string]cache.Item {
	return c.Items()
}
