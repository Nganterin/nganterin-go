package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type HotelReservationOutput struct {
	ReservationKey    string    `json:"reservation_key"`
	ReservationStatus string    `json:"reservation_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type HotelMonthlyReservation struct {
	MonthYear        string `json:"month"`
	ReservationCount int    `json:"reservations"`
}

type HotelYearlyReservationAnalytic struct {
	Period             string                    `json:"period"`
	TrendPercentage    float32                   `json:"trend_percentage"`
	MonthlyReservation []HotelMonthlyReservation `json:"monthly_reservation"`
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
}