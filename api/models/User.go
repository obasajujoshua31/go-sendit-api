package models

import (
	"errors"
	"sendit-api/api/security"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment;" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Email     string    `gorm:"type:varchar(255);unique;not null;" json:"email"`
	Password  string    `gorm:"type:varchar(255);unique;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

func (u *User) BeforeSave() error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	default:
		if u.Name == "" {
			return errors.New("Required Email")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	}
}

