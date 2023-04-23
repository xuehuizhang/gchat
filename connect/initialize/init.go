package initialize

import (
	"fmt"
	"go.uber.org/zap"
)

func Init() {
	initLog()
}

func initLog() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("初始化日志组件失败:%v", err))
	}
	zap.ReplaceGlobals(logger)
}
