package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID              string     `gorm:"primaryKey"`
	Name            string     `gorm:"not null"`
	Email           string     `gorm:"unique;not null"`
	HashedPassword  string     `gorm:"not null"`
	EmailVerifiedAt *time.Time `gorm:"null;default:null"`
	PhoneNumber     string     `gorm:"not null"`
	Country         string     `gorm:"not null"`
	Province        string     `gorm:"not null"`
	City            string     `gorm:"not null"`
	ZipCode         string     `gorm:"not null"`
	CompleteAddress string     `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `gorm:"null;default:null"`
}
