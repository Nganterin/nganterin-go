package models

import (
	"time"

	"gorm.io/gorm"
)

type Partners struct {
	gorm.Model
	ID              string     `gorm:"primaryKey"`
	Name            string     `gorm:"not null"`
	Email           string     `gorm:"unique;not null"`
	HashedPassword  string     `gorm:"not null"`
	EmailVerifiedAt *time.Time `gorm:"default:null"`
	CompanyName     string     `gorm:"unique;not null"`
	Owner           string     `gorm:"not null"`
	CompanyField    string     `gorm:"not null"`
	CompanyEmail    string     `gorm:"unique;not null"`
	CompanyAddress  string     `gorm:"not null"`
	LegalityFile    string     `gorm:"not null"`
	MOUFile         string     `gorm:"not null"`
	DataVerifiedAt  *time.Time `gorm:"default:null"`
	CreatedAt       time.Time 
	UpdatedAt       time.Time 
	DeletedAt       *time.Time `gorm:"default:null"`
}
