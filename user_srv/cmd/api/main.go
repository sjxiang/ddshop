package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/cache"
	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/conf"
	"github.com/sjxiang/ddshop/user_srv/cmd/api/pkg/log"
)

func Init() {
	log.Init()
	conf.LoadConfig()
	cache.ConnectToRedis(conf.Conf.Redis.Addr, "", 0)
}

func main() {
	
	Init()

	// 装载路由
	r := NewRouter()
	
	/*
		S() 很好用，提供了一个全局的安全访问 logger 的途径
	 */
	zap.S().Info("启动服务器，端口：8000")

	addr := fmt.Sprintf("%s:%d", conf.Conf.App.Host, conf.Conf.App.Port)
	if err := r.Run(addr); err != nil {
		zap.S().Panicf("启动失败，%s", err.Error())	
	}
}

