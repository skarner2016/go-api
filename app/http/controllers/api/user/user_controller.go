package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/app/http/controllers/api"
	"go-api/app/http/services/userService"
	"go-api/library/enum/area"
	"go-api/library/enum/errorCode"
	"go-api/library/helper"
	"go-api/library/log"
)

type UserController struct {
	api.ApiController
}

func (con UserController) AreaCode(c *gin.Context) {
	areaList := area.NewArea().GetAreaList()

	con.Success(c, areaList)
}

func (con UserController) VerificationCode(c *gin.Context) {
	params := &userService.VerificationCodeParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		con.BadRequest(c, helper.GetValidMessage(err, params))
		return
	}

	// TODO: 请求限频，防止刷短信等
	if err := userService.NewLoginService().SendVerificationCode(params.VerificationType, params.AreaCode, params.Mobile, params.Email); err != nil {
		con.Fail(c, errorCode.SendMessageFail)
		return
	}

	con.Success(c, nil)
	return
}

func (con UserController) Register(c *gin.Context) {
	params := &userService.RegisterParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		con.BadRequest(c, helper.GetValidMessage(err, params))
		return
	}

	loginService := userService.NewLoginService()
	vc, err := loginService.GetVerificationCode(params.VerificationType, params.AreaCode, params.Mobile, params.Email)
	if err != nil {
		log.NewLogger().Error(err)
		con.Fail(c, errorCode.UnknownError)
		return
	}

	if vc != params.Code {
		con.Fail(c, errorCode.RegisterCodeError)
		return
	}

	if err := loginService.Register(params); err != nil {
		log.NewLogger().Error(err)
		con.Fail(c, errorCode.RegisterFail)
		return
	}

	if err := loginService.DelVerificationCode(params.VerificationType, params.AreaCode, params.Mobile, params.Email); err != nil {
		log.NewLogger().Error(err)
		con.Fail(c, errorCode.DelRegisterCodeError)
		return
	}

	con.Success(c, nil)
	return
}

func (con UserController) Login(c *gin.Context)  {
	params := &userService.LoginParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		con.BadRequest(c, helper.GetValidMessage(err, params))
		return
	}

	loginService := userService.NewLoginService()
	vc, err := loginService.GetVerificationCode(params.VerificationType, params.AreaCode, params.Mobile, params.Email)
	if err != nil {
		log.NewLogger().Error(err)
		con.Fail(c, errorCode.LoginFail)
		return
	}

	if vc != params.Code {
		con.Fail(c, errorCode.LoginCodeError)
		return
	}

	user, err := loginService.GetUserInfo(params.VerificationType, params.AreaCode, params.Mobile, params.Email)
	if err != nil {
		con.Fail(c, errorCode.LoginFail)
		return
	}


	// TODO: JWT
	//claims := &userService.JWTClaims{
	//	ID:       user.ID,
	//	AreaCode: user.AreaCode,
	//	Mobile:   user.Mobile,
	//	Email:    user.Email,
	//}


	fmt.Println(user)

}

func (con UserController) UserInfo(c *gin.Context) {
	// TODO 用户信息
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "UserController",
		"data": "UserInfo",
	})
}
