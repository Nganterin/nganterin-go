package models

import (
	"time"

	"gorm.io/gorm"
)

type Hotels struct {
	gorm.Model

	ID             string           `gorm:"primaryKey"`
	PartnerID      string           `gorm:"not null"`
	Name           string           `gorm:"not null"`
	Description    string           `gorm:"not null"`
	HotelDetails   []HotelDetails   `gorm:"foreignKey:HotelID;references:ID"`
	HotelsLocation []HotelsLocation `gorm:"foreignKey:HotelID;references:ID"`
	HotelPhotos    []HotelPhotos    `gorm:"foreignKey:HotelID;references:ID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"null;default:null"`
}

type HotelDetails struct {
	gorm.Model

	ID             uint   `gorm:"primaryKey"`
	HotelID        string `gorm:"not null"`
	Type           string `gorm:"not null"`
	MaxVisitor     int    `gorm:"not null"`
	BedType        string `gorm:"not null"`
	RoomSize       int    `gorm:"not null"`
	SmokingAllowed bool   `gorm:"not null"`
	OvernightPrice int64  `gorm:"not null"`
	TotalRoom      int    `gorm:"not null"`
	TotalBooked    int    `gorm:"not null"`
}

type HotelsLocation struct {
	gorm.Model

	ID              uint   `gorm:"primaryKey"`
	HotelID         string `gorm:"not null"`
	Country         string `gorm:"not null"`
	State           string `gorm:"not null"`
	City            string `gorm:"not null"`
	ZipCode         string `gorm:"not null"`
	CompleteAddress string `gorm:"not null"`
	Gmaps           string `gorm:"not null"`
}

type HotelPhotos struct {
	gorm.Model

	ID      uint   `gorm:"primaryKey"`
	HotelID string `gorm:"not null"`
	URL     string `gorm:"not null"`
}
