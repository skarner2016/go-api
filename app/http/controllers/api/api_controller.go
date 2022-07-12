package api

import (
	"github.com/gin-gonic/gin"
	"go-api/app/http/controllers"
	"go-api/library/enum/errorCode"
	"net/http"
)

type ApiController struct {
	controllers.BaseController
}

func (a *ApiController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

func (a *ApiController) Fail(c *gin.Context, code errorCode.ErrorCode)  {
	msg := errorCode.NewError().GetMsg(code)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": nil,
	})
}

func (a *ApiController) BadRequest(c *gin.Context, msg string)  {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  msg,
		"data": nil,
	})
}
