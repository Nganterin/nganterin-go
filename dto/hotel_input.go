package dto

type HotelInputDTO struct {
	PartnerID      string                `json:"partner_id" validate:"required"`
	Name           string                `json:"name" validate:"required"`
	Description    string                `json:"description" validate:"required"`
	HotelDetails   []HotelDetailInput    `json:"hotel_details"`
	HotelsLocation []HotelsLocationInput `json:"hotels_location"`
	HotelPhotos    []HotelPhotoInput     `json:"hotel_photos"`
}

type HotelDetailInput struct {
	Type           string `json:"type" validate:"required"`
	MaxVisitor     int    `json:"max_visitor" validate:"required"`
	BedType        string `json:"bed_type" validate:"required"`
	RoomSize       int    `json:"room_size" validate:"required"`
	SmokingAllowed bool   `json:"smoking_allowed" validate:"required"`
	OvernightPrice int64  `json:"overnight_price" validate:"required"`
	TotalRoom      int    `json:"total_room" validate:"required"`
	TotalBooked    int    `json:"total_booked" validate:"required"`
}

type HotelsLocationInput struct {
	Country         string `json:"country" validate:"required"`
	State           string `json:"state" validate:"required"`
	City            string `json:"city" validate:"required"`
	ZipCode         string `json:"zip_code" validate:"required"`
	CompleteAddress string `json:"complete_address" validate:"required"`
	Gmaps           string `json:"gmaps" validate:"required"`
}

type HotelPhotoInput struct {
	URL string `json:"url" validate:"required"`
}
