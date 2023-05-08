package log

import "go.uber.org/zap"


func Init() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)  // 把 logger 替换为`全局的`
}