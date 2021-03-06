package userService

import (
	"github.com/stretchr/testify/assert"
	"go-api/library/config"
	"testing"
)

func TestLoginService_GenerateVerificationCode(t *testing.T) {
	vc := NewLoginService().generateVerificationCode()

	assert.GreaterOrEqual(t, vc, minVerificationCode)
	assert.LessOrEqual(t, vc, maxVerificationCode)
}

func TestLoginService_SendEmailVerificationCode(t *testing.T) {
	config.InitConfig()

	email := config.VipConfig.GetString("test.email")

	err := NewLoginService().sendEmailVerificationCode(VerificationTypeMobileRegister, email)

	assert.Nil(t, err)
}

func TestLoginService_SetEmailVerificationCode(t *testing.T) {
	config.InitConfig()

	areaCode := config.VipConfig.GetInt64("test.areaCode")
	mobile := config.VipConfig.GetInt64("test.mobile")

	err := NewLoginService().sendMobileVerificationCode(VerificationTypeMobileRegister, areaCode, mobile)

	assert.Nil(t, err)
}

func TestLoginService_GetVerificationCode(t *testing.T) {
	config.InitConfig()

	cacheKey := "abc"
	_, err := NewLoginService().getVerificationCodeFromCache(cacheKey)
	assert.Nil(t, err)
}

func TestLoginService_MobileRegister(t *testing.T) {
	config.InitConfig()

	areaCode := config.VipConfig.GetInt64("test.areaCode")
	mobile := config.VipConfig.GetInt64("test.mobile")

	err := NewLoginService().mobileRegister(areaCode, mobile)
	assert.Nil(t, err)
}

func TestLoginService_EmailRegister(t *testing.T) {
	config.InitConfig()

	email := config.VipConfig.GetString("test.email")

	err := NewLoginService().emailRegister(email)
	assert.Nil(t, err)
}

func TestLoginService_DelVerificationCode(t *testing.T) {
	config.InitConfig()

	verificationType := VerificationTypeMobileRegister
	areaCode := config.VipConfig.GetInt64("test.areaCode")
	mobile := config.VipConfig.GetInt64("test.mobile")
	email := config.VipConfig.GetString("test.email")

	err := NewLoginService().DelVerificationCode(verificationType, areaCode, mobile, email)

	assert.Nil(t, err)
}