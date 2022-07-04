package routes

import (
	"go-api/app/controllers/api/user"

	"github.com/gin-gonic/gin"
)

const (
	UserGroup = "/api/user"

	UserInfo = "info"
	Register = "register"
)

func addApiRoute(apiRoute *gin.Engine) {

	userGroup := apiRoute.Group(UserGroup)
	{
		userGroup.GET(UserInfo, user.UserController{}.UserInfo)
		userGroup.GET(Register, user.UserController{}.Register)
	}
}
