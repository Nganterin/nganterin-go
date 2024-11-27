package models

import "time"

type Hotels struct {
	ID              string     `gorm:"primaryKey"`
	Type            string     `gorm:"not null"`
	MaxVisitor      int        `gorm:"not null"`
	BedType         string     `gorm:"not null"`
	RoomSize        int        `gorm:"not null"`
	SmokeingAllowed bool       `gorm:"not null"`
	OvernightPrice  int64      `gorm:"not null"`
	TotalRoom       int        `gorm:"not null"`
	TotalBooked     int        `gorm:"not null"`
	CreatedAt       time.Time  
	UpdatedAt       time.Time  
	DeletedAt       *time.Time `gorm:"null;default:null"`
}
