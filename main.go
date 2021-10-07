package main

import (
	"server/pkg/config"
	"server/pkg/log"
	"server/server"
)

func main()  {
	srv:=server.NewServer() //创建一个服务器
	srv.Provide(
		log.GetLogger, //依赖注入Logger
		config.NewConfig,//依赖注入配置文件
	)
	srv.Run()//运行服务
}
