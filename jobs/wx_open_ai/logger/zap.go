package logger

import (
	"errors"
	"fmt"
	"os"
	"wx_open_ai/core"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zaoLogger = &zap.Logger{}

func GetLog() *zap.Logger {
	return zaoLogger
}

// InitZap Zap 获取 zap.Logger
func InitZap() {
	if ok, _ := PathExists("logger"); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", "logger")
		_ = os.Mkdir("logger", os.ModePerm)
	}

	cores := core.Zap.GetZapCores()
	zaoLogger = zap.New(zapcore.NewTee(cores...))

	zaoLogger = zaoLogger.WithOptions(zap.AddCaller())
	fmt.Println("zap logger: init success")
}

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
