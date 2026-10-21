package main

import (
	"wx_open_ai/initialize"
	"wx_open_ai/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release

func main() {
	gin.SetMode(AppMode)

	// TODO：载入日志组件
	logger.InitZap()
	zap.ReplaceGlobals(logger.GetLog())
	logger.GetLog().Info("server run success on ", zap.String("zap_log", "zap_log"))

	// TODO：启动服务
	initialize.RunServer()
}
