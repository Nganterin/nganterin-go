package database

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID                string `gorm:"primaryKey"`
	Name              string `gorm:"not null"`
	Email             string `gorm:"unique;not null"`
	HashedPassword    string 
	HashedGoogleSUB   string
	EmailVerifiedAt   *time.Time `gorm:"null;default:null"`
	PhoneNumber       string     `gorm:"not null"`
	Country           string     `gorm:"not null"`
	Province          string     `gorm:"not null"`
	City              string     `gorm:"not null"`
	ZipCode           string     `gorm:"not null"`
	CompleteAddress   string     `gorm:"not null"`
	IsProfileComplete bool       `gorm:"not null;default:false"`
	Avatar            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `gorm:"null;default:null"`

	UserTokens   []UserTokens   `gorm:"foreignKey:UserID;references:ID"`
	HotelReviews []HotelReviews `gorm:"foreignKey:UserID;references:ID"`
	HotelOrders  []HotelOrders  `gorm:"foreignKey:UserID;references:ID"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) error {
    if u.HashedPassword == "" && u.HashedGoogleSUB == "" {
        return errors.New("either HashedPassword or HashedGoogleSUB must have a value")
    }
    return nil
}

func (u *Users) BeforeUpdate(tx *gorm.DB) error {
    if u.HashedPassword == "" && u.HashedGoogleSUB == "" {
        return errors.New("either HashedPassword or HashedGoogleSUB must have a value")
    }
    return nil
}