package models

import (
	"errors"
	"strings"
	"time"
)

type Parcel struct {
	ID             uint32    `gorm:"primary_key;auto_increment;" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null;" json:"name"`
	Status         string    `gorm:"type:varchar(100);not null;default:'pending'" json:"status"`
	PickUpLocation string    `gorm:"type:varchar(255);not null;" json:"pick_up_location"`
	Destination    string    `gorm:"type:varchar(255);not null;" json:"destination"`
	CreatedAt      time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	Owner          User      `json:"owner"`
	OwnerID        uint32    `gorm:"not null;" json:"owner_id"`
}

func (p *Parcel) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if p.PickUpLocation == "" {
			return errors.New("Required PickUpLocation")
		}
		if p.Destination == "" {
			return errors.New("Required Destination")
		}
		if p.Name == "" {
			return errors.New("Required Name")
		}
		if p.OwnerID < 1 {
			return errors.New("Required Owner")
		}
		return nil
	default:
		return nil
	}
}
