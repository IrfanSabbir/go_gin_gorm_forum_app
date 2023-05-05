package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint64
	Title     string `gorm:"size:255;not null;unique" json:"title"`
	Content   string `gorm:"text;not null;" json:"content"`
	Author    User   `json:"author"`
	AuthorID  uint64 `gorm:"not null" json:"author_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
