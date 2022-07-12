package main

import (
	"go-api/library/config"
	"go-api/routes"
)

func main() {
	// 初始化配置
	config.InitConfig()
	// 初始化路由
	router := routes.InitRoute()

	if err := router.Run(); err != nil {
		panic(err)
	}
}
