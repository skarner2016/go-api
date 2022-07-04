package main

import (
	"fmt"
	"go-api/library/config"
	"go-api/routes"
)

func main() {
	// 初始化配置
	config.InitConfig()
	fmt.Println("main")

	// 初始化路由
	routes.InitRoute()
}
