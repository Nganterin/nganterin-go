package models

import (
	"time"

	"gorm.io/gorm"
)

type UserTokens struct {
	gorm.Model
	ID        int64     `gorm:"primaryKey"`
	UserID    string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	Token     string    `gorm:"not null"`
	ExpiredAt time.Time `gorm:"not null"`
}

type PartnerTokens struct {
	gorm.Model
	ID        int64     `gorm:"primaryKey"`
	PartnerID string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	Token     string    `gorm:"not null"`
	ExpiredAt time.Time `gorm:"not null"`
}
