package dto

type HotelReviewInput struct {
	HotelOrdersID  string `json:"hotel_order_id" validate:"required"`
	UserID         string
	Review         string `json:"review" validate:"required"`
	Cleanliness    int    `json:"cleanliness" validate:"required,min=1,max=5"`
	Comfort        int    `json:"comfort" validate:"required,min=1,max=5"`
	ServiceQuality int    `json:"service_quality" validate:"required,min=1,max=5"`
	Facilities     int    `json:"facilities" validate:"required,min=1,max=5"`
	ValueForMoney  int    `json:"value_for_money" validate:"required,min=1,max=5"`
	Rating         int    `json:"rating" validate:"required,min=1,max=5"`
}