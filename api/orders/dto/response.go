package dto

import (
	"time"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type HotelOrderDetailsOutput struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	HotelID          string    `json:"hotel_id"`
	HotelRoomID      uint      `json:"hotel_room_id"`
	CheckInDate      time.Time `json:"check_in_date"`
	CheckOutDate     time.Time `json:"check_out_date"`
	TotalDays        int       `json:"total_days"`
	TotalPrice       int64     `json:"total_price"`
	PaymentStatus    string    `json:"payment_status"`
	SnapToken        string    `json:"snap_token"`
	SpecialRequest   string    `json:"special_request"`
	IsForSomeoneElse bool      `json:"is_for_someone_else"`
	SomeoneName      string    `json:"someone_name"`
	SomeoneRegion    string    `json:"someone_region"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	HotelReservations HotelReservationOutput `json:"hotel_reservation"`
	Hotel             HotelOutputDTO         `json:"hotel"`
	User              UserOutputDTO          `json:"user"`
	HotelRoom         HotelRoomOutput        `json:"hotel_room"`
}

type HotelReservationOutput struct {
	ReservationKey    string    `json:"reservation_key"`
	ReservationStatus string    `json:"reservation_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type HotelOutputDTO struct {
	ID              string                  `json:"id"`
	PartnerID       string                  `json:"partner_id"`
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	PricingStart    int64                   `json:"pricing_start,omitempty"`
}

type HotelRoomOutput struct {
	ID             uint   `json:"id"`
	Type           string `json:"type"`
	MaxVisitor     int    `json:"max_visitor"`
	BedType        string `json:"bed_type"`
	RoomSize       int    `json:"room_size"`
	SmokingAllowed bool   `json:"smoking_allowed"`
	OvernightPrice int64  `json:"overnight_price"`
	TotalRoom      int    `json:"total_room"`
	TotalBooked    int    `json:"total_booked"`

	HotelRoomPhotoOutput []HotelRoomPhotoOutput `json:"hotel_room_photos" mapstructure:"hotel_room_photos"`
}

type HotelRoomPhotoOutput struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type UserOutputDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type HotelOrderOutput struct {
	ID          string `json:"id"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}


type HotelMonthlyOrderAnalytic struct {
	TotalOrder       int  `json:"total_order"`
	TotalReservation int  `json:"total_reservation"`
	TotalIncome      int64  `json:"total_income"`
	TotalIncomeAlt   string `json:"total_income_alt"`
}

type HotelYearlyOrderAnalytic struct {
	TotalOrder       int                     `json:"total_order"`
	TotalReservation int                     `json:"total_reservation"`
	TotalIncome      int64                     `json:"total_income"`
	TotalIncomeAlt   string                    `json:"total_income_alt"`
	LastMonthData    HotelMonthlyOrderAnalytic `json:"last_month_data"`
}