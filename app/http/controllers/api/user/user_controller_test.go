package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api/app/http/controllers/api/user"
	"go-api/library/config"
	"go-api/routes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
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
	method := http.MethodGet
	url := routes.UserGroup + routes.VerificationCode

	params := user.VerificationCodeParams{
		VerificationType: 1,
		AreaCode:         86,
		Mobile:           13333333333,
		Email:            "123456@qq.com",
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
	assert.Equal(t, 200, w.Code)

	fmt.Println(w.Body.String())
}

func TestUserController_Register(t *testing.T) {
	method := http.MethodGet
	url := routes.UserGroup + routes.Register

	params := user.RegisterParams{
		RegisterType: user.VerificationTypeMobileRegister,
		AreaCode:     86,
		Mobile:       13333333333,
		Email:        "",
		Password:     "",
	}
	paramsJson, err := json.Marshal(&params)
	if err != nil {
		t.Error(err)
	}

	body := bytes.NewReader(paramsJson)
	w := UserControllerInit(url, method, body)
	assert.Equal(t, 200, w.Code)

	fmt.Println(w.Body.String())
}
