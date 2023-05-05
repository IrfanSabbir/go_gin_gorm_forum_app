package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64
	Username   string `gorm:"size:255;not null;unique" json:"username"`
	Email      string `gorm:"size:100;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null;" json:"password"`
	AvatarPath string `gorm:"size:255;null;" json:"avatar_path"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
