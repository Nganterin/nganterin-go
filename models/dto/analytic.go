package dto

type HotelMonthlyReservation struct {
	MonthYear        string `json:"month"`
	ReservationCount int    `json:"reservations"`
}
