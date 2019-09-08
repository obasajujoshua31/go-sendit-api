package models

import (
	"sendit-api/api/security"
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
