package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type HotelOutputDTO struct {
	ID              string                  `json:"id"`
	PartnerID       string                  `json:"partner_id"`
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	PricingStart    int64                   `json:"pricing_start,omitempty"`
	Rating          HotelAverageRating      `json:"rating,omitempty"`
	ReviewStatistic HotelReviewStatistic    `json:"review_statistic,omitempty"`
	HotelRooms      []HotelRoomOutput       `json:"hotel_rooms,omitempty"`
	HotelsLocation  HotelsLocationOutput    `json:"hotels_location,omitempty"`
	HotelPhotos     []HotelPhotoOutput      `json:"hotel_photos,omitempty"`
	HotelFacilities []HotelFacilitiesOutput `json:"hotel_facilities,omitempty"`
	HotelReviews    []HotelReviewOutput     `json:"hotel_reviews,omitempty"`
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

	HotelRoomPhotoInput []HotelRoomPhotoInput `json:"hotel_room_photos" mapstructure:"hotel_room_photos"`
}

type HotelsLocationOutput struct {
	ID              uint   `json:"id"`
	Country         string `json:"country"`
	State           string `json:"state"`
	City            string `json:"city"`
	ZipCode         string `json:"zip_code"`
	CompleteAddress string `json:"complete_address"`
	Gmaps           string `json:"gmaps"`
}

type HotelPhotoOutput struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type HotelRoomPhotoOutput struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type HotelFacilitiesOutput struct {
	ID       uint   `json:"id"`
	Facility string `json:"facility"`
}

type HotelAverageRating struct {
	Rating         float32 `json:"rating"`
	Cleanliness    float32 `json:"cleanliness"`
	Comfort        float32 `json:"comfort"`
	ServiceQuality float32 `json:"service_quality"`
	Facilities     float32 `json:"facilities"`
	ValueForMoney  float32 `json:"value_for_money"`
}

type HotelReviewStatistic struct {
	TotalReviews  int     `json:"total_reviews"`
	AverageRating float32 `json:"average_rating"`
	Percentage5   int     `json:"percentage_5"`
	Percentage4   int     `json:"percentage_4"`
	Percentage3   int     `json:"percentage_3"`
	Percentage2   int     `json:"percentage_2"`
	Percentage1   int     `json:"percentage_1"`
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
