package userModel

import (
	"database/sql"
	"time"
)

type UserModel struct {
	ID        int64        `json:"id" gorm:"primary_key"`
	Name      string       `json:"name" gorm:"not null;unique"`
	Email     string       `json:"email" gorm:"default:null"`
	AreaCode  int64        `json:"area_code" gorm:"default:null"`
	Mobile    int64        `json:"mobile" gorm:"default:null"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"autoDeletedTime"`
}

func (UserModel) TableName() string {
	return "users"
}
