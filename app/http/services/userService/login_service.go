package userService

import (
	"context"
	"errors"
	redis2 "github.com/go-redis/redis/v8"
	"go-api/app/http/models/userModel"
	"go-api/app/http/repository/userRepository"
	"go-api/library/cache"
	"go-api/library/cache/redis"
	"go-api/library/helper"
	"go-api/library/log"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
)

type LoginService struct {
}

const (
	VerificationTypeMobileRegister int64 = 1
	VerificationTypeEmailRegister  int64 = 2

	// 验证码的范围
	maxVerificationCode int64 = 999999
	minVerificationCode int64 = 100000

	// 验证码有效期，十分钟
	verificationCodeTTL = 600
)

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (l *LoginService) SendVerificationCode(vcp *VerificationCodeParams) error {

	switch vcp.VerificationType {
	case VerificationTypeMobileRegister:
		return l.sendMobileVerificationCode(vcp.AreaCode, vcp.Mobile)
	case VerificationTypeEmailRegister:
		return l.sendEmailVerificationCode(vcp.Email)
	default:
		// lang
		return errors.New("错误的类型")
	}
}

func (l *LoginService) sendMobileVerificationCode(areaCode, mobile int64) error {
	vc := l.generateVerificationCode()
	// TODO  发送手机验证码

	return l.SetMobileVerificationCode(areaCode, mobile, vc)
}

func (l *LoginService) sendEmailVerificationCode(email string) error {
	vc := l.generateVerificationCode()
	// TODO 发送邮箱验证码

	return l.SetEmailVerificationCode(email, vc)
}

func (l *LoginService) generateVerificationCode() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	var vc int64
	for {
		vc = r.Int63n(maxVerificationCode)
		if vc >= minVerificationCode {
			break
		}
	}

	return vc
}

func (l *LoginService) SetMobileVerificationCode(areaCode, mobile, verificationCode int64) error {
	cacheKey := cache.NewKey().GetMobileVerificationCode(areaCode, mobile)

	return l.setVerificationCode(cacheKey, verificationCode)
}

func (l *LoginService) SetEmailVerificationCode(email string, verificationCode int64) error {
	cacheKey := cache.NewKey().GetEmailVerificationCode(email)

	return l.setVerificationCode(cacheKey, verificationCode)
}

func (l *LoginService) setVerificationCode(cacheKey string, verificationCode int64) error {

	redisInstance, err := redis.NewRedis(redis.InstantDefault)
	if err != nil {
		log.NewLogger().Error("初始化redis失败", zap.String("error:", err.Error()))
		return err
	}

	ttl := time.Second * verificationCodeTTL
	redisInstance.SetEX(context.Background(), cacheKey, verificationCode, ttl)

	return nil
}

func (l *LoginService) getVerificationCode(cacheKey string) (int64, error) {
	redisInstance, err := redis.NewRedis(redis.InstantDefault)
	if err != nil {
		log.NewLogger().Error("初始化redis失败", zap.String("error:", err.Error()))
		return 0, err
	}

	vcs, err := redisInstance.Get(context.Background(), cacheKey).Result()
	if err == redis2.Nil {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(vcs, 10, 64)
}

func (l *LoginService) GetVerificationCode(r *RegisterParams) (int64, error) {
	switch r.RegisterType {
	case VerificationTypeMobileRegister:
		return l.getMobileVerificationCode(r.AreaCode, r.Mobile)
	case VerificationTypeEmailRegister:
		return l.getEmailVerificationCode(r.Email)
	default:
		// lang
		return 0, errors.New("不支持的注册类型")
	}
}

func (l *LoginService) getMobileVerificationCode(areaCode, mobile int64) (int64, error) {
	cacheKey := cache.NewKey().GetMobileVerificationCode(areaCode, mobile)

	return l.getVerificationCode(cacheKey)
}

func (l *LoginService) getEmailVerificationCode(email string) (int64, error) {
	cacheKey := cache.NewKey().GetEmailVerificationCode(email)

	return l.getVerificationCode(cacheKey)
}

func (l *LoginService) Register(r *RegisterParams) error {
	switch r.RegisterType {
	case VerificationTypeMobileRegister:
		return l.mobileRegister(r.AreaCode, r.Mobile)
	case VerificationTypeEmailRegister:
		return l.emailRegister(r.Email)
	default:
		// lang
		return errors.New("不支持的注册类型")
	}
}

func (l *LoginService) generateNickName() string {
	// TODO  随机字符串
	return helper.RandString(15)
}

func (l *LoginService) mobileRegister(areaCode, mobile int64) error {
	name := l.generateNickName()
	user := &userModel.UserModel{
		Name:     name,
		AreaCode: areaCode,
		Mobile:   mobile,
	}

	userRepository.NewUserRepository().CreateUser(user)

	return nil
}

func (l *LoginService) emailRegister(email string) error {
	name := l.generateNickName()
	user := &userModel.UserModel{
		Name:  name,
		Email: email,
	}

	userRepository.NewUserRepository().CreateUser(user)

	return nil
}
