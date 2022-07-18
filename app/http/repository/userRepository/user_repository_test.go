package userRepository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-api/app/http/models/userModel"
	"go-api/library/config"
	"go-api/library/helper"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	config.InitConfig()

	user := &userModel.UserModel{
		Name:     helper.RandString(15),
		//AreaCode: config.VipConfig.GetInt64("test.areaCode"),
		//Mobile:   config.VipConfig.GetInt64("test.mobile"),
		Email: config.VipConfig.GetString("test.email"),
	}

	err := NewUserRepository().CreateUser(user)

	assert.Nil(t, err)
}

func TestUserRepository_CreateUser2(t *testing.T) {
	config.InitConfig()

	user := &userModel.UserModel{
		Name:     helper.RandString(15),
		AreaCode: config.VipConfig.GetInt64("test.areaCode"),
		Mobile:   config.VipConfig.GetInt64("test.mobile"),
		//Email: config.VipConfig.GetString("test.email"),
	}

	err := NewUserRepository().CreateUser(user)

	assert.Nil(t, err)
}

func TestUserRepository_GetUserByMobile(t *testing.T) {
	config.InitConfig()

	areaCode := config.VipConfig.GetInt64("test.areaCode")
	mobile := config.VipConfig.GetInt64("test.mobile")
	//mobile = 131

	res := NewUserRepository().GetUserByMobile(areaCode, mobile)

	assert.Nil(t, res)



	if res == nil {
		fmt.Println("nil:", res)
	} else {
		fmt.Println("user:", res)
	}
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	config.InitConfig()

	email := config.VipConfig.GetString("test.email")

	email = "abc"

	res := NewUserRepository().GetUserByEmail(email)
	if res == nil {
		fmt.Println("nil:", res)
	} else {
		fmt.Println("user:", res)
	}
}
