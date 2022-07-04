package user

import (
	"fmt"
	"go-api/app/controllers/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	api.ApiController
}

type Register struct {
	AreaCode int64 `form:"area_code" json:"area_code" binding:"required"`
	Mobile   int64
	Code     int64
}

func (con UserController) UserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":  "UserController",
		"data": "UserInfo",
	})
}

func (con UserController) Register(c *gin.Context) {
	// TODO 表单验证
	var json Register
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	fmt.Println(json)

	c.JSON(200, gin.H{
		"msg":  "UserController",
		"data": "Register",
	})
}
