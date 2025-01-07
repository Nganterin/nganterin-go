package dto

type HotelMonthlyReservation struct {
	MonthYear        string `json:"month"`
	ReservationCount int    `json:"reservations"`
}

type HotelYearlyReservationAnalytic struct {
	Period             string                    `json:"period"`
	TrendPercentage    float32                   `json:"trend_percentage"`
	MonthlyReservation []HotelMonthlyReservation `json:"monthly_reservation"`
}
