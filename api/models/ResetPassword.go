package models

import (
	"time"
)

type ResetPassword struct {
	ID        uint64
	Email     string `gorm:"size:100;not null;" json:"email"`
	Token     string `gorm:"size:255;not null;" json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
