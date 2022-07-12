package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/library/config"
	"go-api/routes"
)

func ApiControllerInit() *gin.Engine {
	// 初始化配置
	config.InitConfig()
	fmt.Println("main")

	// 初始化路由
	routes.InitRoute()

	return routes.InitRoute()
}
