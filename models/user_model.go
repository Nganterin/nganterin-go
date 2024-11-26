package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID              string `gorm:"type:uniqueidentifier;default:NEWID()"`
	Name            string `gorm:"not null"`
	Email           string `gorm:"unique;not null"`
	HashedPassword  string `gorm:"not null"`
	EmailVerifiedAt time.Time
	PhoneNumber     string `gorm:"not null"`
	Country         string `gorm:"not null"`
	Province        string `gorm:"not null"`
	City            string `gorm:"not null"`
	ZipCode         string `gorm:"not null"`
	CompleteAddress string `gorm:"not null"`
}
