package userRepository

import (
	"go-api/app/http/models/userModel"
	"go-api/library/database/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	db, _ := mysql.NewMysql(mysql.InstantDefault)
	return &UserRepository{
		db,
	}
}

func (u *UserRepository) CreateUser(user *userModel.UserModel) error {
	return u.DB.Create(user).Error
}

func (u *UserRepository) GetUserByMobile(areaCode, mobile int64) *userModel.UserModel {
	user := &userModel.UserModel{}
	res := u.DB.Where("area_code", areaCode).
		Where("mobile", mobile).
		Find(&user)
	if res.RowsAffected == 0 {
		return nil
	}

	return user
}

func (u *UserRepository) GetUserByEmail(email string) *userModel.UserModel {
	user := &userModel.UserModel{}
	res := u.DB.Where("email", email).Find(&user)
	if res.RowsAffected == 0 {
		return nil
	}

	return user
}
