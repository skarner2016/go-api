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
	VerificationTypeMobileRegister int64 = 10
	VerificationTypeEmailRegister  int64 = 11
	VerificationTypeMobileLogin    int64 = 20
	VerificationTypeEmailLogin     int64 = 21

	// 验证码的范围
	maxVerificationCode int64 = 999999
	minVerificationCode int64 = 100000

	// 验证码有效期，十分钟
	verificationCodeTTL = 600
)

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (l *LoginService) SendVerificationCode(verificationType, areaCode, mobile int64, email string) error {
	switch verificationType {
	case VerificationTypeMobileRegister:
		return l.sendMobileVerificationCode(verificationType, areaCode, mobile)
	case VerificationTypeEmailRegister:
		return l.sendEmailVerificationCode(verificationType, email)
	case VerificationTypeMobileLogin:
		return l.sendMobileVerificationCode(verificationType, areaCode, mobile)
	case VerificationTypeEmailLogin:
		return l.sendEmailVerificationCode(verificationType, email)
	default:
		// lang
		return errors.New("错误的类型")
	}
}

func (l *LoginService) sendMobileVerificationCode(verificationType, areaCode, mobile int64) error {
	vc := l.generateVerificationCode()
	// TODO  发送手机验证码

	return l.SetMobileVerificationCode(verificationType, areaCode, mobile, vc)
}

func (l *LoginService) sendEmailVerificationCode(verificationType int64, email string) error {
	vc := l.generateVerificationCode()
	// TODO 发送邮箱验证码

	return l.SetEmailVerificationCode(email, verificationType, vc)
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

func (l *LoginService) SetMobileVerificationCode(verificationType, areaCode, mobile, verificationCode int64) error {
	cacheKey := cache.NewKey().GetMobileVerificationCode(verificationType, areaCode, mobile)

	return l.setVerificationCodeCache(cacheKey, verificationCode)
}

func (l *LoginService) SetEmailVerificationCode(email string, verificationType, verificationCode int64) error {
	cacheKey := cache.NewKey().GetEmailVerificationCode(verificationType, email)

	return l.setVerificationCodeCache(cacheKey, verificationCode)
}

func (l *LoginService) setVerificationCodeCache(cacheKey string, verificationCode int64) error {
	redisInstance, err := redis.NewRedis(redis.InstantDefault)
	if err != nil {
		log.NewLogger().Error("初始化redis失败", zap.String("error:", err.Error()))
		return err
	}

	ttl := time.Second * verificationCodeTTL
	redisInstance.SetEX(context.Background(), cacheKey, verificationCode, ttl)

	return nil
}

func (l *LoginService) getVerificationCodeFromCache(cacheKey string) (int64, error) {
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

func (l *LoginService) DelVerificationCode(verificationType, areaCode, mobile int64, email string) error {
	var cacheKey string

	switch verificationType {
	case VerificationTypeMobileRegister, VerificationTypeMobileLogin:
		cacheKey = cache.NewKey().GetMobileVerificationCode(verificationType, areaCode, mobile)
	case VerificationTypeEmailRegister, VerificationTypeEmailLogin:
		cacheKey = cache.NewKey().GetEmailVerificationCode(verificationType, email)
	default:
		return errors.New("不支持的类型")
	}

	err := l.delVerificationCodeCache(cacheKey)
	if err != redis2.Nil {
		return nil
	}

	return err
}

func (l *LoginService) delVerificationCodeCache(cacheKey string) error {
	redisInstance, err := redis.NewRedis(redis.InstantDefault)
	if err != nil {
		log.NewLogger().Error("初始化redis失败", zap.String("error:", err.Error()))
		return err
	}

	err = redisInstance.Del(context.Background(), cacheKey).Err()
	if err == redis2.Nil {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func (l *LoginService) GetVerificationCode(verificationType, areaCode, mobile int64, email string) (int64, error) {
	switch verificationType {
	case VerificationTypeMobileRegister,VerificationTypeMobileLogin:
		return l.getMobileVerificationCode(verificationType, areaCode, mobile)
	case VerificationTypeEmailRegister, VerificationTypeEmailLogin:
		return l.getEmailVerificationCode(verificationType, email)
	default:
		// lang
		return 0, errors.New("不支持的注册类型")
	}
}

func (l *LoginService) getMobileVerificationCode(verificationType, areaCode, mobile int64) (int64, error) {
	cacheKey := cache.NewKey().GetMobileVerificationCode(verificationType, areaCode, mobile)

	return l.getVerificationCodeFromCache(cacheKey)
}

func (l *LoginService) getEmailVerificationCode(verificationType int64, email string) (int64, error) {
	cacheKey := cache.NewKey().GetEmailVerificationCode(verificationType, email)

	return l.getVerificationCodeFromCache(cacheKey)
}

func (l *LoginService) Register(r *RegisterParams) error {
	switch r.VerificationType {
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

	return userRepository.NewUserRepository().CreateUser(user)
}

func (l *LoginService) emailRegister(email string) error {
	name := l.generateNickName()
	user := &userModel.UserModel{
		Name:  name,
		Email: email,
	}

	return userRepository.NewUserRepository().CreateUser(user)
}

func (l *LoginService) GetUserInfo(verificationType, areaCode, mobile int64, email string) (*userModel.UserModel, error) {
	switch verificationType {
	case VerificationTypeMobileLogin:
		return userRepository.
			NewUserRepository().
			GetUserByMobile(areaCode, mobile), nil
	case VerificationTypeEmailLogin:
		return userRepository.
			NewUserRepository().
			GetUserByEmail(email), nil
	default:
		return nil, errors.New("不支持的注册类型")
	}
}
