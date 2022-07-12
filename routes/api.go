package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/app/http/controllers/api/user"
	"net/http"
)

const (
	UserGroup = "/api/user"

	AreaCode         = "/area-code"
	VerificationCode = "/verification-code"
	Register         = "/register"
	UserInfo         = "/info"
)

func addApiRoute(apiRoute *gin.Engine) {

	apiRoute.GET(AreaCode, func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	userGroup := apiRoute.Group(UserGroup)
	{
		userGroup.GET(AreaCode, user.UserController{}.AreaCode)
		userGroup.GET(VerificationCode, user.UserController{}.VerificationCode)
		userGroup.GET(UserInfo, user.UserController{}.UserInfo)
		userGroup.GET(Register, user.UserController{}.Register)
	}
}
