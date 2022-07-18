package models

import "time"

type BaseModel struct {
	ID uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
