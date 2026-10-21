package initialize

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wx_open_ai/logger"

	"go.uber.org/zap"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

func RunServer() {
	address := fmt.Sprintf(":%d", 8088)
	// 初始化路由
	rou := new(R).Routers()
	if rou == nil {
		return
	}
	s := initServer(address, rou)

	logger.GetLog().Info("server run success on ", zap.String("address", address))
	time.Sleep(10 * time.Microsecond)
	logger.GetLog().Error(s.ListenAndServe().Error())

	// 优雅停止
	sigChan := make(chan os.Signal, 1)
	// 监听操作系统发送的中断和终止信号（SIGINT和SIGTERM）
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan // 阻塞在此，当接收到上述两种信号时才会往下执行
	fmt.Printf("Received signal: %v\n", sig)
	// 关闭服务器前等待一段时间，处理完已有的连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP Server
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.Shutdown(ctx); err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	}
	fmt.Println("Server shutdown")
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	// 设置HTTP请求头的读取超时时间为20秒。
	s.ReadHeaderTimeout = 20 * time.Second
	// 设置HTTP响应体的写入超时时间为20秒。
	s.WriteTimeout = 20 * time.Second
	// 设置HTTP请求头的最大字节数为1MB。如果请求头超过1MB，则会返回一个错误。
	s.MaxHeaderBytes = 1 << 20
	return s
}
