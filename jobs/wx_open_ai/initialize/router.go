package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type R struct {
}

func (*R) Routers() *gin.Engine {
	Router := gin.New()
	// 路由处理函数中的panic可能会导致整个服务器崩溃，为了避免这种情况，Gin提供的Recovery中间件来处理panic。
	Router.Use(gin.Recovery())

	// 静态文件服务
	//公共路由
	PublicGroup := Router.Group("webDao")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	return Router
}
