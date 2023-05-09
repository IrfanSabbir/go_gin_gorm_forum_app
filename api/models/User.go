package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/IrfanSabbir/go_gin_gorm_forum_app/api/security"
	"github.com/badoux/checkmail"
)

type User struct {
	ID         uint64
	Username   string         `gorm:"size:255;not null;unique" json:"username"`
	Email      string         `gorm:"size:100;not null;unique" json:"email"`
	Password   string         `gorm:"size:100;not null;" json:"password"`
	AvatarPath string         `gorm:"size:255;null;" json:"avatar_path"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index;->:false;<-:create"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) map[string]string {
	var errList = make(map[string]string)
	var errMessage error
	switch strings.ToLower(action) {
	case "update":
	default:
		if u.Username == "" {
			errMessage = errors.New("Username cant be empty")
			errList["Required_username"] = errMessage.Error()
		}
		if u.Email == "" {
			errMessage = errors.New("Email cant be empty")
			errList["Required_email"] = errMessage.Error()
		}
		if u.Email != "" {
			if errMessage = checkmail.ValidateFormat(u.Email); errMessage != nil {
				errList["Invalid_email"] = errMessage.Error()
			}
		}

		if u.Password == "" {
			errMessage = errors.New("Password cant be empty")
			errList["Required_password"] = errMessage.Error()
		}

		if u.Password != "" && len(u.Password) < 5 {
			errMessage = errors.New("Passowrd length should be at lease 5 character")
			errList["Invalid_password"] = errMessage.Error()
		}
	}
	return errList
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	u.Password = ""
	return u, nil
}
