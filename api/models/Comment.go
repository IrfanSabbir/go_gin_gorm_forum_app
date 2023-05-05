package models

import (
	"time"
)

type Comment struct {
	ID        uint64
	UserID    uint64 `gorm:"not null" json:"user_id"`
	PostID    uint64 `gorm:"not null" json:"post_id"`
	Body      string `gorm:"text;not null;" json:"body"`
	User      User   `json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
