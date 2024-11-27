package models

import (
	"time"

	"gorm.io/gorm"
)

type Partners struct {
	gorm.Model
	ID              string     `gorm:"type:uniqueidentifier;default:NEWID()"`
	Name            string     `gorm:"not null"`
	Email           string     `gorm:"unique;not null"`
	HashedPassword  string     `gorm:"not null"`
	EmailVerifiedAt *time.Time `gorm:"null;default:null"`
	CompanyName     string     `gorm:"unique;not null"`
	Owner           string     `gorm:"not null"`
	CompanyField    string     `gorm:"not null"`
	CompanyEmail    string     `gorm:"unique;not null"`
	CompanyAddress  string     `gorm:"not null"`
	LegalityFile    string     `gorm:"not null"`
	MOUFile         string     `gorm:"not null"`
	DataVerifiedAt  *time.Time `gorm:"null;default:null"`
	CreatedAt       time.Time  `gorm:"null"`
	UpdatedAt       time.Time  `gorm:"null"`
	DeletedAt       *time.Time `gorm:"null;default:null"`
}
