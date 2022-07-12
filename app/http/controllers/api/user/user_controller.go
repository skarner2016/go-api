package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/app/http/controllers/api"
	"go-api/library/enum/area"
	"go-api/library/enum/errorCode"
	"go-api/library/helper"
)

type UserController struct {
	api.ApiController
}

type RegisterParams struct {
	AreaCode int64  `form:"area_code" json:"area_code" binding:"required" msg:"请选择正确的区域代码"`
	Mobile   int64  `json:"mobile"`
	User     string `json:"user" binding:"required" msg:"用户名不能为空"`
	//Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type VerificationCodeParams struct {
	VerificationType  int64  `json:"type" binding:"required,numeric,gt=0" msg:"错误的验证码类型"`
	VerificationValue string `json:"value" binding:"required"`
}

const (
	VerificationTypeMobile int64 = 1
	VerificationTypeEmail  int64 = 2
)

func (con UserController) AreaCode(c *gin.Context) {
	areaList := area.NewArea().GetAreaList()

	con.Success(c, areaList)
}

func (con UserController) VerificationCode(c *gin.Context) {
	params := &VerificationCodeParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		con.BadRequest(c, helper.GetValidMessage(err, params))
		return
	}

	switch params.VerificationType {
	case VerificationTypeMobile:
		// TODO  发送手机验证码
		fmt.Println(VerificationTypeMobile)
	case VerificationTypeEmail:
		// TODO 发送邮箱验证码
		fmt.Println(VerificationTypeEmail)
	default:
		con.Fail(c, errorCode.NetWorkError)
	}
}

func (con UserController) Register(c *gin.Context) {
	params := &RegisterParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		con.BadRequest(c, helper.GetValidMessage(err, params))
		return
	}

	// TODO 注册逻辑
	c.JSON(200, gin.H{
		"msg":  "UserController",
		"data": "RegisterParams",
	})
}

func (con UserController) UserInfo(c *gin.Context) {
	// TODO 用户信息
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "UserController",
		"data": "UserInfo",
	})
}
