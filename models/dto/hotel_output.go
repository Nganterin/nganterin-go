package dto

import "time"

type HotelOutputDTO struct {
	ID              string                  `json:"id"`
	PartnerID       string                  `json:"partner_id"`
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	PricingStart    int64                   `json:"pricing_start,omitempty"`
	Rating          HotelAverageRating      `json:"rating"`
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

type HotelOrderOutput struct {
	ID          string `json:"id"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
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
	HotelRoom         HotelRoomOutput        `json:"hotel_room"`
}

type HotelReservationOutput struct {
	ReservationKey    string    `json:"reservation_key"`
	ReservationStatus string    `json:"reservation_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
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

	User UserReviewOutputDTO `json:"user"`
}
