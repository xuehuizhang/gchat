package initialize

import (
	"com.gchat.business/user/dao"
	"fmt"
	"go.uber.org/zap"
)

func Init() {
	initLog()
	dao.Setup("root", "admin123", "127.0.0.1", 3306, "gchat")
}

func initLog() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("初始化日志组件失败:%v", err))
	}
	zap.ReplaceGlobals(logger)
}
