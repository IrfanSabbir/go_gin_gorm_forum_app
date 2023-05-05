package models

import (
	"time"
)

type Like struct {
	ID        uint64
	UserID    uint64 `gorm:"not null" json:"user_id"`
	PostID    uint64 `gorm:"not null" json:"post_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
