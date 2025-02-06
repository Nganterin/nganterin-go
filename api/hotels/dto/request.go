package dto

type HotelInputDTO struct {
	PartnerID   string
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`

	HotelRooms      []HotelRoomInput       `json:"hotel_rooms" validate:"required,dive"`
	HotelsLocation  HotelsLocationInput    `json:"hotels_location" validate:"required"`
	HotelPhotos     []HotelPhotoInput      `json:"hotel_photos" validate:"required,dive"`
	HotelFacilities []HotelFacilitiesInput `json:"hotel_facilities" validate:"required,dive"`
}

type HotelRoomInput struct {
	Type           string `json:"type" validate:"required"`
	MaxVisitor     int    `json:"max_visitor" validate:"required,number"`
	BedType        string `json:"bed_type" validate:"required"`
	RoomSize       int    `json:"room_size" validate:"required,number"`
	SmokingAllowed bool   `json:"smoking_allowed" validate:"boolean"`
	OvernightPrice int64  `json:"overnight_price" validate:"required,number,min=0"`
	TotalRoom      int    `json:"total_room" validate:"required,number,min=0"`
	TotalBooked    int    `json:"total_booked" validate:"number,min=0"`

	HotelRoomPhotoInput []HotelRoomPhotoInput `json:"hotel_room_photos" validate:"required,dive" mapstructure:"hotel_room_photos"`
}

type HotelsLocationInput struct {
	Country         string `json:"country" validate:"required"`
	State           string `json:"state" validate:"required"`
	City            string `json:"city" validate:"required"`
	ZipCode         string `json:"zip_code" validate:"required"`
	CompleteAddress string `json:"complete_address" validate:"required"`
	Gmaps           string `json:"gmaps" validate:"required,url"`
}

type HotelPhotoInput struct {
	URL string `json:"url" validate:"required,url"`
}

type HotelRoomPhotoInput struct {
	URL string `json:"url" validate:"required,url"`
}

type HotelFacilitiesInput struct {
	Facility string `json:"facility" validate:"required"`
}

type HotelSearch struct {
	Keyword        string
	Name           string
	PriceStart     int64
	PriceEnd       int64
	City           string
	Country        string
	MinimumStars   int
	MinimumVisitor int
}
