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