package routes

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine {
	r := gin.Default()

	addApiRoute(r)

	return r
}
