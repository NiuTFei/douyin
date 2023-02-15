package main

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//数据库初始化
	initError := service.Init()
	if initError != nil {
		os.Exit(-1)
	}

	go service.RunMessageServer()

	//服务器初始化并配置路由
	server := gin.Default()
	initRouter(server)

	//启动服务
	err := server.Run()
	if err != nil {
		return
	}
}
