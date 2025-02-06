package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type HotelReviewOutput struct {
	ID             uint      `json:"id"`
	HotelOrdersID  string    `json:"hotel_order_id"`
	Review         string    `json:"review"`
	Cleanliness    int       `json:"cleanliness"`
	Comfort        int       `json:"comfort"`
	ServiceQuality int       `json:"service_quality"`
	Facilities     int       `json:"facilities"`
	ValueForMoney  int       `json:"value_for_money"`
	Rating         int       `json:"rating"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	User UserOutputDTO `json:"user"`
}

type UserOutputDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}
