package routes

import "github.com/gin-gonic/gin"

func InitRoute() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":"pong",
		})
	})

	addApiRoute(r)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
