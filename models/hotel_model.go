package models

import (
	"time"

	"gorm.io/gorm"
)

type Hotels struct {
	gorm.Model

	ID          string `gorm:"primaryKey"`
	PartnerID   string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"null;default:null"`

	HotelRooms      []HotelRooms      `gorm:"foreignKey:HotelID;references:ID"`
	HotelsLocation  HotelsLocation    `gorm:"foreignKey:HotelID;references:ID"`
	HotelPhotos     []HotelPhotos     `gorm:"foreignKey:HotelID;references:ID"`
	HotelFacilities []HotelFacilities `gorm:"foreignKey:HotelID;references:ID"`
	HotelReviews    []HotelReviews    `gorm:"foreignKey:HotelID;references:ID"`
	HotelOrders     []HotelOrders     `gorm:"foreignKey:HotelID;references:ID"`
}

type HotelRooms struct {
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

	HotelOrders     []HotelOrders     `gorm:"foreignKey:HotelRoomID;references:ID"`
	HotelRoomPhotos []HotelRoomPhotos `gorm:"foreignKey:HotelRoomID;references:ID" mapstructure:"hotel_room_photos"`
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

type HotelRoomPhotos struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	HotelRoomID string `gorm:"not null"`
	URL         string `gorm:"not null"`
}

type HotelFacilities struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	HotelID  string `gorm:"not null"`
	Facility string `gorm:"not null"`
}

type HotelOrders struct {
	gorm.Model

	ID               string    `gorm:"primaryKey"`
	UserID           string    `gorm:"not null"`
	HotelID          string    `gorm:"not null"`
	HotelRoomID      uint      `gorm:"not null"`
	CheckInDate      time.Time `gorm:"not null"`
	CheckOutDate     time.Time `gorm:"not null"`
	TotalPrice       int64     `gorm:"not null"`
	PaymentStatus    string    `gorm:"not null;default:pending"` // Enum: "Pending", "Paid", "Cancelled"
	SpecialRequest   string    `gorm:"type:text"`
	IsForSomeoneElse bool      `gorm:"not null;default:false"`
	SomeoneName      string
	SomeoneRegion    string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `gorm:"null;default:null"`

	HotelReservations HotelReservations `gorm:"foreignKey:HotelOrdersID;references:ID"`
}

type HotelReservations struct {
	gorm.Model

	ID                string `gorm:"primaryKey"`
	HotelOrdersID     string `gorm:"not null"`
	UserID            string `gorm:"not null"`
	ReservationKey    string `gorm:"not null"`
	ReservationStatus string `gorm:"not null"` // Enum: "Confirmed", "CheckedIn", "Completed"
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `gorm:"null;default:null"`

	HotelReviews HotelReviews `gorm:"foreignKey:HotelReservationID;references:ID"`
}

type HotelReviews struct {
	gorm.Model

	ID                 uint   `gorm:"primaryKey"`
	HotelReservationID string `gorm:"not null"`
	HotelID            string `gorm:"not null"`
	UserID             string `gorm:"not null"`
	Review             string `gorm:"not null"`
	Rating             int    `gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `gorm:"null;default:null"`
}
