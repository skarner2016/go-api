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

func (u *UserRepository) CreateUser(user *userModel.UserModel) {
	u.DB.Create(user)
}
