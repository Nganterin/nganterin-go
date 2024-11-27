package models

import (
	"time"

	"gorm.io/gorm"
)

type Tokens struct {
	gorm.Model
	ID        int64     `gorm:"primaryKey"`
	UserID    string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	Token     string    `gorm:"not null"`
	ExpiredAt time.Time `gorm:"not null"`
}
