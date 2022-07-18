package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api/app/http/services/userService"
	"go-api/library/config"
	"go-api/routes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"

)

func UserControllerInit(url, method string, body io.Reader) *httptest.ResponseRecorder {
	// 初始化配置
	config.InitConfig()
	// 初始化路由
	router := routes.InitRoute()

	w := httptest.NewRecorder()

	request, _ := http.NewRequest(method, url, body)
	request.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, request)

	return w
}

func TestUserController_AreaCode(t *testing.T) {
	method := http.MethodGet
	url := routes.UserGroup + routes.AreaCode

	w := UserControllerInit(url, method, nil)
	assert.Equal(t, 200, w.Code)

	response := w.Body.String()
	fmt.Println(response)
}

func TestUserController_VerificationCode(t *testing.T) {
	config.InitConfig()

	method := http.MethodGet
	url := routes.UserGroup + routes.VerificationCode

	params := userService.VerificationCodeParams{
		VerificationType: userService.VerificationTypeMobileLogin,
		AreaCode:         config.VipConfig.GetInt64("test.areaCode"),
		Mobile:           config.VipConfig.GetInt64("test.mobile"),
		Email:            config.VipConfig.GetString("test.email"),
	}

	paramsJson, err := json.Marshal(&params)
	if err != nil {
		t.Error(err)
	}

	body := bytes.NewReader(paramsJson)
	w := UserControllerInit(url, method, body)

	fmt.Println("response:", w.Body.String())
	assert.Equal(t, 200, w.Code)
}

func TestUserController_UserInfo(t *testing.T) {
	method := http.MethodGet
	url := routes.UserGroup + routes.UserInfo

	w := UserControllerInit(url, method, nil)

	fmt.Println("response:", w.Body.String())
	assert.Equal(t, 200, w.Code)
}

func TestUserController_Register(t *testing.T) {
	config.InitConfig()

	code := int64(355813)
	params := &userService.RegisterParams{
		VerificationType: userService.VerificationTypeMobileRegister,
		AreaCode:         config.VipConfig.GetInt64("test.areaCode"),
		Mobile:           config.VipConfig.GetInt64("test.mobile"),
		Email:            config.VipConfig.GetString("test.email"),
		Password:         config.VipConfig.GetString("test.password"),
		Code:             code,
	}

	paramsJson, err := json.Marshal(&params)
	assert.Nil(t, err)

	body := bytes.NewReader(paramsJson)

	method := http.MethodGet
	url := routes.UserGroup + routes.Register
	w := UserControllerInit(url, method, body)
	fmt.Println("response:", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestUserController_Login(t *testing.T) {
	config.InitConfig()

	verificationType := userService.VerificationTypeMobileLogin
	code := int64(452499)
	params := &userService.LoginParams{
		VerificationType: verificationType,
		AreaCode:         config.VipConfig.GetInt64("test.areaCode"),
		Mobile:           config.VipConfig.GetInt64("test.mobile"),
		Email:            config.VipConfig.GetString("test.email"),
		Code:             code,
	}
	paramsJson, err := json.Marshal(params)
	assert.Nil(t, err)

	body := bytes.NewReader(paramsJson)

	method := http.MethodGet
	url := routes.UserGroup + routes.Login

	w := UserControllerInit(url, method, body)
	fmt.Println(w.Body.String())

	assert.Equal(t, 200, w.Code)
}


